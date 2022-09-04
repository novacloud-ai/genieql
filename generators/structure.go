package generators

import (
	"go/ast"
	"go/types"
	"html/template"
	"io"
	"strings"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/internal/drivers"
	"bitbucket.org/jatone/genieql/internal/transformx"
)

// StructOption option to provide the structure function.
type StructOption func(*structure)

// StructOptionName provide the name of the struct to the structure.
func StructOptionName(n string) StructOption {
	return func(s *structure) {
		s.Name = n
	}
}

// StructOptionComment specify the comment for the structure
func StructOptionComment(comment *ast.CommentGroup) StructOption {
	return func(s *structure) {
		s.Comment = comment
	}
}

// StructOptionAliasStrategy provides the default aliasing strategy for
// generating the a struct's field names.
func StructOptionAliasStrategy(mcp genieql.MappingConfigOption) StructOption {
	return func(s *structure) {
		s.aliaser = mcp
	}
}

// StructOptionColumnsStrategy strategy for resolving column info for the structure.
func StructOptionColumnsStrategy(strategy columnsStrategy) StructOption {
	return func(s *structure) {
		s.columns = strategy
	}
}

// StructOptionTableStrategy convience function for creating a table based structure.
func StructOptionTableStrategy(table string) StructOption {
	return StructOptionColumnsStrategy(func(ctx Context) ([]genieql.ColumnInfo, error) {
		return ctx.Dialect.ColumnInformationForTable(ctx.Driver, table)
	})
}

// StructOptionQueryStrategy convience function for creating a query based structure.
func StructOptionQueryStrategy(query string) StructOption {
	return StructOptionColumnsStrategy(func(ctx Context) ([]genieql.ColumnInfo, error) {
		return ctx.Dialect.ColumnInformationForQuery(ctx.Driver, query)
	})
}

// StructOptionRenameMap provides explicit rename mappings when
// generating the struct's field names.
func StructOptionRenameMap(m map[string]string) StructOption {
	return func(s *structure) {
		s.renameMap = genieql.MCORenameMap(m)
	}
}

// StructOptionContext sets the Context for the structure generator.
func StructOptionContext(c Context) StructOption {
	return func(s *structure) {
		s.Context = c
	}
}

// StructOptionMappingConfigOptions sets the base configuration to be used for
// the MappingConfig.
func StructOptionMappingConfigOptions(options ...genieql.MappingConfigOption) StructOption {
	return func(s *structure) {
		s.mappingOptions = options
	}
}

// StructOptionFromCommentGroup parses a configuration and converts it into an array of options.
func StructOptionFromCommentGroup(comment *ast.CommentGroup) ([]StructOption, error) {
	const aliasOption = `alias`
	const generalSection = `general`
	const renameSection = `rename.columns`

	options := []StructOption{}
	ini, err := ParseCommentOptions(comment)
	if err != nil {
		return options, err
	}
	if kvmap, ok := ini.GetKvmap(renameSection); ok {
		options = append(options, StructOptionRenameMap(kvmap))
	}

	if kvmap, ok := ini.GetKvmap(generalSection); ok {
		if alias := genieql.AliaserSelect(kvmap[aliasOption]); alias != nil {
			// this could cause multiple aliasers to be applied to the Generator
			// but it doesn't matter as last one will win.
			options = append(options, StructOptionAliasStrategy(genieql.MCOTransformations(kvmap[aliasOption])))
		}
	}

	return options, nil
}

// NewStructure creates a Generator that builds structures from column information.
func NewStructure(opts ...StructOption) genieql.Generator {
	s := structure{
		renameMap:      genieql.MCORenameMap(map[string]string{}),
		aliaser:        genieql.MCOTransformations("camelcase"),
		mappingOptions: []genieql.MappingConfigOption{},
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

// StructureFromGenDecl creates a structure generator from  from the provided *ast.GenDecl
func StructureFromGenDecl(decl *ast.GenDecl, columnStrategyBuilder func(string) StructOption, options ...StructOption) []genieql.Generator {
	var (
		err        error
		configOpts []StructOption
	)

	if decl.Doc == nil {
		configOpts = options
	} else {
		if configOpts, err = StructOptionFromCommentGroup(decl.Doc); err != nil {
			return []genieql.Generator{genieql.NewErrGenerator(err)}
		}
		configOpts = append(options, configOpts...)
	}

	specs := genieql.FindValueSpecs(decl)
	g := make([]genieql.Generator, 0, len(specs))

	for _, vs := range specs {
		m := mapStructureToGenerator{
			options:               configOpts,
			columnStrategyBuilder: columnStrategyBuilder,
		}
		g = append(g, m.Map(vs)...)
	}

	return g
}

type columnsStrategy func(Context) ([]genieql.ColumnInfo, error)
type structure struct {
	Context
	Name           string
	Comment        *ast.CommentGroup
	columns        columnsStrategy
	aliaser        genieql.MappingConfigOption
	renameMap      genieql.MappingConfigOption
	mappingOptions []genieql.MappingConfigOption
}

func (t structure) Generate(dst io.Writer) error {
	const tmpl = `type {{.Name}} struct {
	{{- range $column := .Columns }}
	{{ $column.Name | transformation }} {{ if $column.Definition.Nullable }}*{{ end }}{{ $column.Definition.Native | type -}}
	{{ end }}
}`
	type context struct {
		Name    string
		Columns []genieql.ColumnInfo
	}
	var (
		err     error
		columns []genieql.ColumnInfo
	)

	if columns, err = t.columns(t.Context); err != nil {
		return err
	}

	mapping := genieql.NewMappingConfig(
		append(
			t.mappingOptions,
			t.renameMap,
			t.aliaser,
			genieql.MCOColumns(columns...),
			genieql.MCOType(t.Name),
			genieql.MCOPackage(t.Context.CurrentPackage),
		)...,
	)

	if err = t.Context.Configuration.WriteMap(mapping); err != nil {
		return err
	}

	if err = GenerateComment(t.Comment, DefaultFunctionComment(t.Name)).Generate(dst); err != nil {
		return err
	}

	typeDefinitions := composeTypeDefinitions(t.Driver.LookupType, drivers.DefaultTypeDefinitions)
	ctx := context{
		Name:    t.Name,
		Columns: mapping.Columns,
	}

	a := mapping.Aliaser()

	return template.Must(template.New("scanner template").Funcs(template.FuncMap{
		"transformation": func(s string) string { return transformx.String(s, a) },
		"type": func(s string) string {
			if d, err := typeDefinitions(s); err == nil {
				return d.Native
			}

			return s
		},
	}).Parse(tmpl)).Execute(dst, ctx)
}

type mapStructureToGenerator struct {
	columnStrategyBuilder func(string) StructOption
	options               []StructOption
}

func (t mapStructureToGenerator) Map(vs *ast.ValueSpec) []genieql.Generator {
	dst := make([]genieql.Generator, 0, len(vs.Names))

	for idx := range vs.Names {
		tableOrQuery := strings.Trim(types.ExprString(vs.Values[idx]), "\"")
		s := NewStructure(
			append(t.options,
				StructOptionName(
					vs.Names[idx].Name,
				),
				t.columnStrategyBuilder(tableOrQuery),
			)...,
		)
		dst = append(dst, s)
	}

	return dst
}
