package generators

import (
	"fmt"
	"go/ast"
	"go/types"
	"io"
	"strings"
	"text/template"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/astutil"
	"bitbucket.org/jatone/genieql/internal/drivers"
)

type exploderFunction struct {
	fields []*ast.Field
	queryFunction
}

func (t exploderFunction) Generate(dst io.Writer) error {
	type context struct {
		Name       string
		Fields     []*ast.Field
		Parameters []*ast.Field
	}

	return t.queryFunction.Template.Execute(dst, context{
		Fields:     t.fields,
		Name:       t.Name,
		Parameters: t.Parameters,
	})
}

// NewExploderFunction ...
func NewExploderFunction(ctx Context, param *ast.Field, fields []*ast.Field, options ...QueryFunctionOption) genieql.Generator {
	const defaultQueryFunc = `// {{.Name}} generated by genieql
		func {{.Name}}({{ .Parameters | arguments }}) ([]interface{}, error) {
			var (
				{{- range $index, $column := .Fields }}
				c{{ $index }} {{ $column | type | typedef | sqltype -}} // {{ $column | name -}}
				{{ end }}
			)

			{{ range $index, $field := .Fields }}
			{{- $d := $field | type | typedef -}}
			{{- $cinfo := $field | cinfo -}}
			{{- range $_, $stmt := encode $index $cinfo error -}}
			{{ $stmt | ast }}
			{{ end }}
			{{ end }}
			return []interface{}{{"{"}}{{ .Fields | localvars }}{{"}"}}, nil
		}
		`

	var (
		typedef             = composeTypeDefinitionsExpr(ctx.Driver.LookupType, drivers.DefaultTypeDefinitions)
		defaultQueryFuncMap = template.FuncMap{
			"typedef": typedef,
			"type": func(field *ast.Field) ast.Expr {
				return field.Type
			},
			"sqltype": func(d genieql.ColumnDefinition) string {
				return d.ColumnType
			},
			"arguments": argumentsAsPointers,
			"ast":       astPrint,
			"encode":    ColumnMapEncoder(ctx),
			"error": func() func(string) ast.Node {
				return func(local string) ast.Node {
					return astutil.Return(
						ast.NewIdent("[]interface{}(nil)"),
						ast.NewIdent(local),
					)
				}
			},
			"localvars": func(fields []*ast.Field) (s string) {
				locals := make([]string, 0, len(fields))
				for idx := range fields {
					locals = append(locals, fmt.Sprintf("c%d", idx))
				}
				return strings.Join(locals, ",")
			},
			"cinfo": func(field *ast.Field) genieql.ColumnMap {
				td, err := typedef(field.Type)
				if err != nil {
					td = fallbackDefinition(types.ExprString(field.Type))
				}
				return genieql.ColumnMap{
					ColumnInfo: genieql.ColumnInfo{
						Definition: td,
					},
					Dst: &ast.SelectorExpr{
						X:   param.Names[0],
						Sel: astutil.MapFieldsToNameIdent(field)[0],
					},
				}
			},
			"name": func(field *ast.Field) string {
				for _, n := range field.Names {
					return n.Name
				}
				return ""
			},
		}
		tmpl = template.Must(template.New("explode-function").Funcs(defaultQueryFuncMap).Parse(defaultQueryFunc))
	)

	qf := queryFunction{
		Context: ctx,
	}
	qf.Apply(append(options, QFOExplodeStructParam(param, fields...), QFOTemplate(tmpl))...)

	return exploderFunction{
		fields:        fields,
		queryFunction: qf,
	}
}

// QFOExplodeStructParam explodes a structure parameter's fields in the query parameters.
func QFOExplodeStructParam(param *ast.Field, fields ...*ast.Field) QueryFunctionOption {
	selectors := structureQueryParameters(normalizeFieldNames(param)[0], fields...)
	return func(qf *queryFunction) {
		qf.Parameters = append(qf.Parameters, param)
		qf.QueryParameters = append(qf.QueryParameters, selectors...)
	}
}

// StructureQueryParameters - generates QueryParameters for the given struct and its component
// fields.
func StructureQueryParameters(param *ast.Field, fields ...*ast.Field) []ast.Expr {
	return structureQueryParameters(param, fields...)
}

func structureQueryParameters(param *ast.Field, fields ...*ast.Field) []ast.Expr {
	selectors := make([]ast.Expr, 0, len(fields)*len(param.Names))
	for _, name := range param.Names {
		for _, field := range fields {
			selectors = append(selectors, &ast.SelectorExpr{
				X:   name,
				Sel: astutil.MapFieldsToNameIdent(field)[0],
			})
		}
	}

	return selectors
}
