package main

import (
	"fmt"
	"go/build"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/commands"
	"bitbucket.org/jatone/genieql/crud"
	"bitbucket.org/jatone/genieql/generators"
)

type generateInsert struct {
	configName  string
	constSuffix string
	packageType string
	table       string
	output      string
	mapName     string
	batch       int
	defaults    []string
}

func (t *generateInsert) configure(cmd *kingpin.CmdClause) *kingpin.CmdClause {
	insert := cmd.Command("insert", "generate more complicated insert queries")
	insert.Flag(
		"config",
		"name of configuration file to use",
	).Default("default.config").StringVar(&t.configName)

	insert.Flag(
		"mapping",
		"name of the map to use",
	).Default("default").StringVar(&t.mapName)

	insert.Flag("default", "specifies a name of a column to default to database value").
		StringsVar(&t.defaults)

	insert.Flag(
		"output",
		"path of output file",
	).Default("").StringVar(&t.output)

	insert.Flag("batch", "number of records to insert").Default("1").IntVar(&t.batch)

	cmd = insert.Command("constant", "output the query as a constant").Action(t.constant).Default()
	cmd.Flag(
		"suffix",
		"suffix for the name of the generated constant",
	).Required().StringVar(&t.constSuffix)

	cmd.Arg(
		"package.Type",
		"package prefixed structure we want to build the scanner/query for",
	).Required().StringVar(&t.packageType)

	cmd.Arg(
		"table",
		"table you want to build the queries for",
	).Required().StringVar(&t.table)

	x := insert.Command("experimental", "experimental insert commands")
	cmd = x.Command("batch-function", "generate a batch insert function")

	return insert
}

func (t *generateInsert) constant(*kingpin.ParseContext) error {
	var (
		err           error
		configuration genieql.Configuration
		mapping       genieql.MappingConfig
		fset          = token.NewFileSet()
	)

	configuration = genieql.MustReadConfiguration(
		genieql.ConfigurationOptionLocation(
			filepath.Join(genieql.ConfigurationDirectory(), t.configName),
		),
	)

	pkgName, typName := extractPackageType(t.packageType)

	if err = configuration.ReadMap(pkgName, typName, t.mapName, &mapping); err != nil {
		return err
	}

	details, err := genieql.LoadInformation(configuration, t.table)
	if err != nil {
		log.Fatalln(err)
	}

	fields, err := mapping.TypeFields(fset, build.Default, genieql.StrictPackageName(filepath.Base(pkgName)))
	if err != nil {
		log.Println("type fields error")
		log.Fatalln(err)
	}

	constName := fmt.Sprintf("%sInsert%s", typName, t.constSuffix)

	details = details.OnlyMappedColumns(fields, mapping.Mapper().Aliasers...)

	pkg, err := genieql.LocatePackage(pkgName, build.Default, genieql.StrictPackageName(filepath.Base(pkgName)))
	if err != nil {
		return err
	}

	hg := headerGenerator{
		fset: fset,
		pkg:  pkg,
		args: os.Args[1:],
	}

	cc := generators.NewColumnConstants(
		fmt.Sprintf("%sStaticColumns", constName),
		genieql.ColumnValueTransformer{
			Defaults:           t.defaults,
			DialectTransformer: details.Dialect.ColumnValueTransformer(),
		},
		details.Columns,
	)
	cg := crud.Insert(details).Build(t.batch, constName, t.defaults)

	pg := printGenerator{
		delegate: genieql.MultiGenerate(hg, cc, cg),
	}

	if err = commands.WriteStdoutOrFile(pg, t.output, commands.DefaultWriteFlags); err != nil {
		log.Fatalln(err)
	}

	return nil
}
