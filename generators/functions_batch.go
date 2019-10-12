package generators

import (
	"go/ast"
	"go/token"
	"go/types"
	"io"
	"strconv"
	"text/template"

	"github.com/pkg/errors"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/astutil"
	"bitbucket.org/jatone/genieql/internal/x/stringsx"
)

// BatchFunctionOption ...
type BatchFunctionOption func(*batchFunction)

// BatchFunctionQueryBuilder ...
func BatchFunctionQueryBuilder(query func(n int) ast.Decl) BatchFunctionOption {
	return func(b *batchFunction) {
		b.Builder = query
	}
}

// BatchFunctionQFOptions ...
func BatchFunctionQFOptions(options ...QueryFunctionOption) BatchFunctionOption {
	return func(b *batchFunction) {
		b.queryFunction.Apply(options...)
	}
}

// BatchFunctionExploder ...
func BatchFunctionExploder(sel ...*ast.Field) BatchFunctionOption {
	return func(b *batchFunction) {
		b.Selectors = sel
	}
}

type builder func(local string, n int, columns ...string) ast.Decl

// NewBatchFunctionFromGenDecl creates a function generator from the provided *ast.GenDecl
func NewBatchFunctionFromGenDecl(ctx Context, decl *ast.GenDecl, b builder, defaults []string, options ...BatchFunctionOption) []genieql.Generator {
	g := make([]genieql.Generator, 0, len(decl.Specs))
	for _, spec := range decl.Specs {
		if ts, ok := spec.(*ast.TypeSpec); ok {
			if ft, ok := ts.Type.(*ast.FuncType); ok {
				g = append(g, batchGeneratorFromFuncType(ctx, ts.Name, decl.Doc, ft, b, defaults, options...))
			}
		}
	}

	return g
}

func batchGeneratorFromFuncType(ctx Context, name *ast.Ident, comment *ast.CommentGroup, ft *ast.FuncType, b builder, ignoreSet []string, poptions ...BatchFunctionOption) genieql.Generator {
	var (
		err        error
		qf         queryFunction
		fields     []*ast.Field
		columns    []genieql.ColumnInfo
		qfoOptions []QueryFunctionOption
	)

	// validition...
	if len(ft.Params.List[1:]) > 1 && areArrayType(astutil.MapFieldsToTypExpr(ft.Params.List[1:]...)...) {
		return genieql.NewErrGenerator(errors.New("batch only supports a single array type parameter"))
	}

	max, elt, err := extractArrayInfo(ft.Params.List[1].Type.(*ast.ArrayType))
	if err != nil {
		return genieql.NewErrGenerator(err)
	}
	ft.Params.List[1] = astutil.Field(elt, ft.Params.List[1].Names...)
	field := ft.Params.List[1]
	if !builtinType(elt) && !selectType(elt) {
		if fields, err = mappedFields(ctx, field, ignoreSet...); err != nil {
			return genieql.NewErrGenerator(errors.Wrap(err, "failed to map params"))
		}

		poptions = append(poptions, BatchFunctionExploder(fields...))
	}

	if qfoOptions, err = generatorFromFuncType(ctx, name, comment, ft); err != nil {
		return genieql.NewErrGenerator(err)
	}
	qf.Apply(qfoOptions...)

	if _, columns, err = mappedParam(ctx, field); err != nil {
		return genieql.NewErrGenerator(err)
	}

	builder := func(n int) ast.Decl {
		return b("query", n, genieql.ColumnInfoSet(columns).ColumnNames()...)
	}

	poptions = append(
		poptions,
		BatchFunctionQueryBuilder(builder),
		BatchFunctionQFOptions(
			QFOName(qf.Name),
			QFOScanner(qf.ScannerDecl),
			QFOQueryer(qf.QueryerName, qf.Queryer),
			QFOQueryerFunction(ast.NewIdent("Query")),
		),
	)

	return NewBatchFunction(max, field, poptions...)
}

// NewBatchFunction builds functions that execute on batches of values, such as update and insert.
func NewBatchFunction(maximum int, typ *ast.Field, options ...BatchFunctionOption) genieql.Generator {
	b := batchFunction{
		Maximum:  maximum,
		Type:     typ,
		Template: batchInsertScannerTemplate,
	}

	for _, opt := range options {
		opt(&b)
	}

	b.queryFunction.Apply(QFOSharedParameters(&ast.Field{
		Names: typ.Names,
		Type:  &ast.Ellipsis{Elt: typ.Type},
	}))

	return b
}

type batchFunction struct {
	Context
	Type          *ast.Field
	Maximum       int
	queryFunction queryFunction
	Template      *template.Template
	Builder       func(n int) ast.Decl
	Selectors     []*ast.Field
}

func (t batchFunction) Generate(dst io.Writer) error {
	type queryFunctionContext struct {
		Number       int
		BuiltinQuery ast.Node
		Queryer      ast.Expr
		Exploder     ast.Node
		Explode      ast.Node
	}
	type context struct {
		Type             *ast.Field
		QueryFunction    queryFunction
		ScannerType      ast.Expr
		ScannerFunc      ast.Expr
		DefaultStatement queryFunctionContext
		Statements       []queryFunctionContext
		Parameters       []*ast.Field
	}

	var (
		parameters         []*ast.Field
		queryParameters    []ast.Expr
		defaultQueryParams []ast.Expr
		statements         []queryFunctionContext
		exploderName       = ast.NewIdent("exploder")
		tmpName            = ast.NewIdent("tmp")
		queryField         = astutil.Field(ast.NewIdent("string"), ast.NewIdent("query"))
	)

	parameters = buildParameters(
		t.queryFunction.BuiltinQuery == nil,
		astutil.Field(t.queryFunction.Queryer, ast.NewIdent(t.queryFunction.QueryerName)),
		astutil.Field(&ast.Ellipsis{Elt: t.Type.Type}, t.Type.Names...),
	)

	queryParameters = buildQueryParameters(queryField)
	if len(t.Selectors) == 0 {
		defaultQueryParams = append(queryParameters, &ast.SliceExpr{
			X:    astutil.MapFieldsToNameExpr(t.Type)[0],
			High: &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(t.Maximum)},
		})
		queryParameters = append(queryParameters, astutil.MapFieldsToNameExpr(t.Type)...)
	} else {
		defaultQueryParams = append(queryParameters, &ast.SliceExpr{
			X: tmpName,
		})
		queryParameters = append(queryParameters, &ast.SliceExpr{
			X: tmpName,
		})
	}

	statements = make([]queryFunctionContext, 0, t.Maximum)
	for i := 1; i < t.Maximum; i++ {
		tmp := queryFunctionContext{
			Number:       i,
			BuiltinQuery: t.Builder(i),
			Queryer: &ast.CallExpr{
				Fun:      &ast.SelectorExpr{X: ast.NewIdent(t.queryFunction.QueryerName), Sel: t.queryFunction.QueryerFunction},
				Args:     queryParameters,
				Ellipsis: token.Pos(1),
			},
			Exploder: buildExploder(i, exploderName, t.Type, t.Selectors...),
			Explode:  buildExploderAssign(tmpName, exploderName, astutil.MapFieldsToNameExpr(t.Type), t.Selectors...),
		}

		statements = append(statements, tmp)
	}

	defaultStatement := queryFunctionContext{
		Number:       t.Maximum,
		BuiltinQuery: t.Builder(t.Maximum),
		Exploder:     buildExploder(t.Maximum, exploderName, t.Type, t.Selectors...),
		Explode:      buildExploderAssign(tmpName, exploderName, astutil.ExprList(&ast.SliceExpr{X: astutil.MapFieldsToNameExpr(t.Type)[0], High: &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(t.Maximum)}}), t.Selectors...),
		Queryer: &ast.CallExpr{
			Fun:      &ast.SelectorExpr{X: ast.NewIdent(t.queryFunction.QueryerName), Sel: t.queryFunction.QueryerFunction},
			Args:     defaultQueryParams,
			Ellipsis: token.Pos(1),
		},
	}

	ctx := context{
		QueryFunction:    t.queryFunction,
		Statements:       statements,
		DefaultStatement: defaultStatement,
		ScannerFunc:      t.queryFunction.ScannerDecl.Name,
		ScannerType:      t.queryFunction.ScannerDecl.Type.Results.List[0].Type,
		Parameters:       parameters,
		Type:             t.Type,
	}

	return errors.Wrap(t.Template.Execute(dst, ctx), "failed to generate batch insert")
}

var batchInsertScannerTemplate = template.Must(template.New("batch-function").Funcs(batchQueryFuncMap).Parse(batchScannerTemplate))
var batchQueryFuncMap = template.FuncMap{
	"expr":      types.ExprString,
	"arguments": arguments,
	"ast":       astPrint,
	"array":     exprToArray,
	"name": func(f *ast.Field) ast.Expr {
		return astutil.MapFieldsToNameExpr(f)[0]
	},
	"title":   stringsx.ToPublic,
	"private": stringsx.ToPrivate,
}

const batchScannerTemplate = `// New{{.QueryFunction.Name | title}} creates a scanner that inserts a batch of
// records into the database.
func New{{.QueryFunction.Name | title}}({{ .Parameters | arguments }}) {{ .ScannerType | expr }} {
	return &{{.QueryFunction.Name | private}}{
		q: {{.QueryFunction.QueryerName}},
		remaining: {{.Type | name }},
	}
}

type {{.QueryFunction.Name | private}} struct {
	q         {{.QueryFunction.Queryer | expr}}
	remaining {{ .Type.Type | array | expr }}
	scanner   {{ .ScannerType | expr }}
}

func (t *{{.QueryFunction.Name | private}}) Scan(dst *{{.Type.Type | expr}}) error {
	return t.scanner.Scan(dst)
}

func (t *{{.QueryFunction.Name | private}}) Err() error {
	if t.scanner == nil {
		return nil
	}

	return t.scanner.Err()
}

func (t *{{.QueryFunction.Name | private}}) Close() error {
	if t.scanner == nil {
		return nil
	}
	return t.scanner.Close()
}

func (t *{{.QueryFunction.Name | private}}) Next() bool {
	var (
		advanced bool
	)

	if t.scanner != nil && t.scanner.Next() {
		return true
	}

	// advance to the next check
	if len(t.remaining) > 0 && t.Close() == nil {
		t.scanner, t.remaining, advanced = t.advance(t.q, t.remaining...)
		return advanced && t.scanner.Next()
	}

	return false
}

func (t *{{.QueryFunction.Name | private}}) advance(q sqlx.Queryer, {{.Type | name}} ...{{.Type.Type | expr}}) ({{ .ScannerType | expr }}, {{ .Type.Type | array | expr }}, bool) {
	switch len({{.Type | name }}) {
	case 0:
		return nil, []{{.Type.Type | expr}}(nil), false
	{{- range $ctx := .Statements }}
	case {{ $ctx.Number }}:
		{{ $ctx.BuiltinQuery | ast }}
		{{ $ctx.Exploder | ast }}
		{{ $ctx.Explode | ast }}
		return {{ $.ScannerFunc | expr }}({{ $ctx.Queryer | expr }}), {{$.Type.Type | array | expr}}(nil), true
	{{- end }}
	default:
		{{ .DefaultStatement.BuiltinQuery | ast }}
		{{ .DefaultStatement.Exploder | ast }}
		{{ .DefaultStatement.Explode | ast }}
		return {{ .ScannerFunc | expr }}({{ .DefaultStatement.Queryer | expr }}), {{.Type | name}}[{{.DefaultStatement.Number}}:], true
	}
}
`
