package generators_test

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/astutil"
	"bitbucket.org/jatone/genieql/dialects"
	"bitbucket.org/jatone/genieql/internal/drivers"
	"bitbucket.org/jatone/genieql/internal/errorsx"
	_ "bitbucket.org/jatone/genieql/internal/postgresql"

	. "bitbucket.org/jatone/genieql/generators"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Batch Functions", func() {
	bctx := build.Default
	bctx.Dir = "."

	pkg := &build.Package{
		Name:       "example",
		Dir:        ".fixtures",
		ImportPath: "./.fixtures",
		GoFiles: []string{
			"example.go",
		},
	}

	configuration := genieql.MustConfiguration(
		genieql.ConfigurationOptionLocation(
			filepath.Join(".", ".fixtures", ".genieql", "generators-test.config"),
		),
	)

	driver, err := genieql.LookupDriver(drivers.StandardLib)
	errorsx.PanicOnError(err)

	exampleScanner := &ast.FuncDecl{
		Name: ast.NewIdent("StaticExampleScanner"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					astutil.Field(astutil.Expr("*sql.Rows"), ast.NewIdent("rows")),
					astutil.Field(astutil.Expr("error"), ast.NewIdent("err")),
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{astutil.Field(ast.NewIdent("ExampleScanner"))},
			},
		},
	}
	builder := func(n int) ast.Decl {
		return genieql.QueryLiteral("query", fmt.Sprintf("QUERY %d", n))
	}
	ginkgo.DescribeTable("batch insert scanner generator",
		func(fixture string, maximum int, field *ast.Field, options ...BatchFunctionOption) {
			var (
				buffer    bytes.Buffer
				formatted bytes.Buffer
			)
			ctx := Context{
				Configuration:  configuration,
				CurrentPackage: pkg,
				Build:          bctx,
				FileSet:        token.NewFileSet(),
				Dialect:        dialects.Test{},
				Driver:         driver,
			}
			buffer.WriteString("package example\n\n")
			Expect(NewBatchFunction(ctx, maximum, field, options...).Generate(&buffer)).ToNot(HaveOccurred())
			buffer.WriteString("\n")
			// log.Println("GENERATED", buffer.String())
			Expect(genieql.FormatOutput(&formatted, buffer.Bytes())).ToNot(HaveOccurred())
			expected, err := os.ReadFile(fixture)
			Expect(err).ToNot(HaveOccurred())
			Expect(formatted.String()).To(Equal(string(expected)))
		},
		ginkgo.Entry(
			"batch function (1) integers",
			".fixtures/functions-batch/output1.go",
			1,
			astutil.Field(ast.NewIdent("int"), ast.NewIdent("i")),
			BatchFunctionQueryBuilder(builder),
			BatchFunctionQFOptions(
				QFOName("batchFunction1"),
				QFOScanner(exampleScanner),
				QFOQueryer("q", astutil.MustParseExpr(token.NewFileSet(), "sqlx.Queryer")),
				QFOQueryerFunction(ast.NewIdent("Query")),
			),
		),
		ginkgo.Entry(
			"batch function (2) integers",
			".fixtures/functions-batch/output2.go",
			2,
			astutil.Field(ast.NewIdent("int"), ast.NewIdent("i")),
			BatchFunctionQueryBuilder(builder),
			BatchFunctionQFOptions(
				QFOName("batchFunction2"),
				QFOScanner(exampleScanner),
				QFOQueryer("q", astutil.MustParseExpr(token.NewFileSet(), "sqlx.Queryer")),
				QFOQueryerFunction(ast.NewIdent("Query")),
			),
		),
		ginkgo.Entry(
			"batch function (3) integers",
			".fixtures/functions-batch/output3.go",
			3,
			astutil.Field(ast.NewIdent("custom"), ast.NewIdent("v")),
			BatchFunctionQueryBuilder(builder),
			BatchFunctionExploder(astutil.Field(ast.NewIdent("int"), ast.NewIdent("A")), astutil.Field(ast.NewIdent("int"), ast.NewIdent("B")), astutil.Field(ast.NewIdent("int"), ast.NewIdent("C"))),
			BatchFunctionQFOptions(
				QFOName("batchFunction3"),
				QFOScanner(exampleScanner),
				QFOQueryer("q", astutil.MustParseExpr(token.NewFileSet(), "sqlx.Queryer")),
				QFOQueryerFunction(ast.NewIdent("Query")),
			),
		),
	)

	ginkgo.DescribeTable("build a query function from a function prototype",
		func(prototype, fixture string, options ...BatchFunctionOption) {
			buffer := bytes.NewBuffer([]byte{})
			formatted := bytes.NewBuffer([]byte{})
			builder := func(local string, n int, columns ...string) ast.Decl {
				return genieql.QueryLiteral("query", fmt.Sprintf("QUERY %d", n))
			}
			ctx := Context{
				Configuration:  configuration,
				CurrentPackage: pkg,
				Build:          bctx,
				FileSet:        token.NewFileSet(),
				Dialect:        dialects.Test{},
				Driver:         driver,
			}

			file, err := parser.ParseFile(ctx.FileSet, "prototypes.go", prototype, parser.ParseComments)
			Expect(err).ToNot(HaveOccurred())

			buffer.WriteString("package example\n\n")
			for _, decl := range genieql.FindTypes(file) {
				gen := genieql.MultiGenerate(NewBatchFunctionFromGenDecl(ctx, decl, builder, []string{}, options...)...)
				Expect(gen.Generate(buffer)).ToNot(HaveOccurred())
			}
			buffer.WriteString("\n")
			Expect(genieql.FormatOutput(formatted, buffer.Bytes())).ToNot(HaveOccurred())

			expected, err := os.ReadFile(fixture)
			Expect(err).ToNot(HaveOccurred())
			Expect(formatted.String()).To(Equal(string(expected)))
		},
		ginkgo.Entry(
			"example 1 - structure insert",
			"package example; type batchFunction4 func(q sqlx.Queryer, p [5]StructA) StaticExampleScanner",
			".fixtures/functions-batch/output4.go",
		),
	)
})
