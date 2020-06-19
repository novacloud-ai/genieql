package generators

import (
	"fmt"
	"go/ast"
	"go/token"
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
	// return t.queryFunction.Generate(dst)
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
				c{{ $index }} {{ $column | type | typedef | sqltype -}}
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
		defaultQueryFuncMap = template.FuncMap{
			"typedef": composeTypeDefinitionsExpr(ctx.Driver.LookupType, drivers.DefaultTypeDefinitions),
			"type": func(field *ast.Field) ast.Expr {
				return field.Type
			},
			"sqltype": func(d genieql.NullableTypeDefinition) string {
				return d.NullType
			},
			"arguments": argumentsAsPointers,
			"ast":       astPrint,
			"encode":    encode(ctx),
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
				return genieql.ColumnMap{
					Type: field.Type,
					Dst: &ast.SelectorExpr{
						X:   param.Names[0],
						Sel: astutil.MapFieldsToNameIdent(field)[0],
					},
				}
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

func buildExploder(ctx Context, n int, name ast.Expr, typ *ast.Field, selectors ...*ast.Field) (_ ast.Stmt, err error) {
	// nothing to do.
	if len(selectors) == 0 {
		return nil, nil
	}

	nulltype := nulltypes(ctx)
	encoder := encode(ctx)
	input := &ast.Ellipsis{Elt: typ.Type}
	output := &ast.ArrayType{Elt: ast.NewIdent("interface{}"), Len: astutil.IntegerLiteral(n * len(selectors))}
	returnc := ast.NewIdent("r")
	encoderErr := func(local string) ast.Node {
		return astutil.Return(
			returnc,
			ast.NewIdent(local),
		)
	}
	key := ast.NewIdent("idx")
	value := ast.NewIdent("v")
	assignlhs := make([]ast.Expr, 0, len(selectors))
	assignrhs := make([]ast.Expr, 0, len(selectors))
	encodings := make([]ast.Stmt, 0, len(selectors))
	localspec := make([]ast.Spec, 0, len(selectors))

	for idx, sel := range selectors {
		var (
			encoded []ast.Stmt
		)
		info := genieql.ColumnMap{
			Type: sel.Type,
			Dst: &ast.SelectorExpr{
				X:   value,
				Sel: astutil.MapFieldsToNameIdent(sel)[0],
			},
		}

		assignlhs = append(assignlhs, &ast.IndexExpr{
			X: returnc,
			Index: &ast.BinaryExpr{
				X: &ast.BinaryExpr{
					X:  key,
					Op: token.MUL,
					Y:  astutil.IntegerLiteral(len(selectors)),
				},
				Op: token.ADD,
				Y:  astutil.IntegerLiteral(idx),
			},
		})

		if encoded, err = encoder(idx, info, encoderErr); err != nil {
			return nil, err
		}
		encodings = append(encodings, encoded...)

		localspec = append(localspec, astutil.ValueSpec(nulltype(info.Type), info.Local(idx)))
		assignrhs = append(assignrhs, info.Local(idx))
	}

	stmts := []ast.Stmt{
		&ast.DeclStmt{
			Decl: astutil.VarList(localspec...),
		},
	}
	stmts = append(stmts, encodings...)
	stmts = append(stmts, astutil.Assign(assignlhs, token.ASSIGN, assignrhs))

	body := &ast.RangeStmt{
		Key:   key,
		Value: value,
		Tok:   token.DEFINE,
		X:     &ast.SliceExpr{X: typ.Names[0], High: astutil.IntegerLiteral(n)},
		Body: &ast.BlockStmt{
			List: stmts,
		},
	}

	return &ast.AssignStmt{
		Tok: token.DEFINE,
		Lhs: []ast.Expr{name},
		Rhs: []ast.Expr{
			&ast.FuncLit{
				Type: &ast.FuncType{
					Params: &ast.FieldList{List: []*ast.Field{astutil.Field(input, typ.Names...)}},
					Results: &ast.FieldList{List: []*ast.Field{
						astutil.Field(output, ast.NewIdent("r")),
						astutil.Field(ast.NewIdent("error"), ast.NewIdent("err")),
					}},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						body,
						astutil.Return(returnc, ast.NewIdent("nil")),
					},
				},
			},
		},
	}, nil
}

func buildExploderInvocations(n int, fun ast.Expr, arg ast.Expr) []ast.Expr {
	r := make([]ast.Expr, 0, n)
	for i := 0; i < n; i++ {
		r = append(r, astutil.CallExpr(fun, arg))
	}
	return r
}

func buildExploderAssign(tmpName, exploderName ast.Expr, errReturn ast.Stmt, exploderArg []ast.Expr, selectors ...*ast.Field) []ast.Stmt {
	if len(selectors) == 0 {
		return nil
	}
	errExpr := ast.NewIdent("err")
	return []ast.Stmt{
		astutil.Assign(
			astutil.ExprList(tmpName, errExpr),
			token.DEFINE,
			astutil.ExprList(
				&ast.CallExpr{
					Fun:      exploderName,
					Args:     exploderArg,
					Ellipsis: token.Pos(1),
				},
			),
		),
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{X: errExpr, Op: token.NEQ, Y: ast.NewIdent("nil")},
			Body: astutil.Block(
				errReturn,
			),
		},
	}
}
