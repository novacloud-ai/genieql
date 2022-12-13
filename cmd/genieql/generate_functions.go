package main

import (
	"go/ast"
	"go/build"
	"log"
	"os"
	"path/filepath"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/cmd"
	"bitbucket.org/jatone/genieql/compiler"
	"bitbucket.org/jatone/genieql/generators"
	"bitbucket.org/jatone/genieql/internal/buildx"
	"github.com/alecthomas/kingpin"
)

type generateFunctionTypes struct {
	buildInfo
	configName string
	output     string
	pkg        string
}

func (t *generateFunctionTypes) configure(cmd *kingpin.CmdClause) *kingpin.CmdClause {
	functions := cmd.Command("functions", "commands for generating functions")
	c := functions.Command("types", "generates functions defined by function types within a package").Action(t.execute)
	c.Flag("config", "name of the genieql configuration to use").Default(defaultConfigurationName).StringVar(&t.configName)
	c.Flag("output", "output filename").Short('o').StringVar(&t.output)
	c.Flag("package", "package to search for definitions").Default(t.buildInfo.CurrentPackageImport()).StringVar(&t.pkg)

	return c
}

func (t *generateFunctionTypes) execute(*kingpin.ParseContext) (err error) {
	var (
		ctx         generators.Context
		taggedFiles compiler.TaggedFiles
		tags        = []string{
			"genieql", "generate", "functions",
		}
		pkg *build.Package
	)

	bctx := buildx.Clone(build.Default, buildx.Tags(tags...))

	if pkg, err = genieql.LocatePackage(t.pkg, ".", bctx, nil); err != nil {
		return err
	}

	if ctx, err = generators.NewContext(bctx, t.configName, pkg); err != nil {
		return err
	}

	if taggedFiles, err = compiler.FindTaggedFiles(ctx.Build, t.pkg, tags...); err != nil {
		return err
	}

	if taggedFiles.Empty() {
		// nothing to do.
		log.Println("no files tagged, ignoring")
		return nil
	}

	g := []genieql.Generator{}
	genieql.NewUtils(ctx.FileSet).WalkFiles(func(path string, file *ast.File) {
		if !taggedFiles.IsTagged(filepath.Base(path)) {
			return
		}

		functionsTypes := mapDeclsToGenerator(func(d *ast.GenDecl) []genieql.Generator {
			return generators.NewQueryFunctionFromGenDecl(
				ctx,
				d,
			)
		}, genieql.SelectFuncType(genieql.FindTypes(file)...)...)

		g = append(g, functionsTypes...)

		functions := mapFuncDeclsToGenerator(func(d *ast.FuncDecl) genieql.Generator {
			return generators.NewQueryFunctionFromFuncDecl(ctx, d)
		}, genieql.SelectFuncDecl(func(*ast.FuncDecl) bool { return true }, genieql.FindFunc(file)...)...)

		g = append(g, functions...)
	}, ctx.CurrentPackage)

	mg := genieql.MultiGenerate(g...)

	hg := headerGenerator{
		fset: ctx.FileSet,
		pkg:  ctx.CurrentPackage,
		args: os.Args[1:],
	}

	pg := printGenerator{
		delegate: genieql.MultiGenerate(hg, mg),
	}

	if err = cmd.WriteStdoutOrFile(pg, t.output, cmd.DefaultWriteFlags); err != nil {
		log.Fatalln(err)
	}

	return nil
}
