package generators

import (
	"go/ast"
	"go/token"
	"go/types"
	"io"
	"strings"
	"text/template"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/astutil"
)

type exploderFunction struct {
	queryFunction
}

func (t exploderFunction) Generate(dst io.Writer) error {
	type context struct {
		Name               string
		Parameters         []*ast.Field
		ExplodedParameters []ast.Expr
	}

	return t.queryFunction.Template.Execute(dst, context{
		Name:               t.Name,
		Parameters:         t.Parameters,
		ExplodedParameters: t.QueryParameters,
	})
}

// NewExploderFunction ...
func NewExploderFunction(param *ast.Field, fields []*ast.Field, options ...QueryFunctionOption) genieql.Generator {
	const defaultQueryFunc = `// {{.Name}} generated by genieql
		func {{.Name}}({{ .Parameters | arguments }}) []interface{} {
			return []interface{{"{}"}}{{"{"}}{{ .ExplodedParameters | expr_list }}{{"}"}}
		}
		`
	var (
		defaultQueryFuncMap = template.FuncMap{
			"expr":      types.ExprString,
			"arguments": argumentsAsPointers,
			"ast":       astPrint,
			"expr_list": func(args []ast.Expr) string {
				return strings.Join(astutil.MapExprToString(args...), ",")
			},
		}
		tmpl = template.Must(template.New("explode-function").Funcs(defaultQueryFuncMap).Parse(defaultQueryFunc))
	)
	qf := queryFunction{}
	qf.Apply(append(options, QFOExplodeStructParam(param, fields...), QFOTemplate(tmpl))...)
	return exploderFunction{
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
	key := ast.NewIdent("idx")
	value := ast.NewIdent("v")
	assignlhs := make([]ast.Expr, 0, len(selectors))
	assignrhs := make([]ast.Expr, 0, len(selectors))
	encodings := make([]ast.Stmt, 0, len(selectors))
	localspec := make([]ast.Spec, 0, len(selectors))

	for idx, sel := range selectors {
		var (
			encoded ast.Stmt
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

		if encoded, err = encoder(info.Dst, idx, info); err != nil {
			return nil, err
		}
		encodings = append(encodings, encoded)

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
					Params:  &ast.FieldList{List: []*ast.Field{astutil.Field(input, typ.Names...)}},
					Results: &ast.FieldList{List: []*ast.Field{astutil.Field(output, ast.NewIdent("r"))}},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						body,
						astutil.Return(returnc),
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

func buildExploderAssign(tmpName, exploderName ast.Expr, exploderArg []ast.Expr, selectors ...*ast.Field) ast.Stmt {
	if len(selectors) == 0 {
		return nil
	}

	return astutil.Assign(
		astutil.ExprList(tmpName),
		token.DEFINE,
		astutil.ExprList(
			&ast.CallExpr{
				Fun:      exploderName,
				Args:     exploderArg,
				Ellipsis: token.Pos(1),
			},
		),
	)
}
