package generators

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"io"
	"text/template"

	"github.com/pkg/errors"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/astutil"
)

// QueryFunctionOption options for building query functions.
type QueryFunctionOption func(*queryFunction) error

// QFOName specify the name of the query function.
func QFOName(n string) QueryFunctionOption {
	return func(qf *queryFunction) error {
		qf.Name = n
		return nil
	}
}

// QFOScanner specify the scanner of the query function
func QFOScanner(n *ast.FuncDecl) QueryFunctionOption {
	return func(qf *queryFunction) error {
		qf.ScannerDecl = n
		return nil
	}
}

// QFOBuiltinQuery force the query function to only execute the specified
// query.
func QFOBuiltinQuery(q string) QueryFunctionOption {
	return func(qf *queryFunction) error {
		qf.BuiltinQuery = q
		return nil
	}
}

// QFOQueryer the name/type used to execute queries.
func QFOQueryer(name string, x ast.Expr) QueryFunctionOption {
	return func(qf *queryFunction) error {
		qf.Queryer = x
		qf.QueryerName = name
		return nil
	}
}

// QFOQueryerFunction the function to invoke on the Queryer.
func QFOQueryerFunction(x *ast.Ident) QueryFunctionOption {
	return func(qf *queryFunction) error {
		qf.QueryerFunction = x
		return nil
	}
}

// QFOParameters specify the parameters for the query function.
func QFOParameters(params ...*ast.Field) QueryFunctionOption {
	return func(qf *queryFunction) error {
		qf.Parameters = params
		return nil
	}
}

// NewQueryFunction build a query function generator from the provided options.
func NewQueryFunction(options ...QueryFunctionOption) genieql.Generator {
	qf := queryFunction{
		Parameters:  []*ast.Field{},
		QueryerName: "q",
		Queryer:     &ast.StarExpr{X: &ast.SelectorExpr{X: ast.NewIdent("sql"), Sel: ast.NewIdent("DB")}},
	}

	if err := qf.Apply(options...); err != nil {
		return genieql.NewErrGenerator(err)
	}

	pattern := astutil.MapFieldsToTypExpr(qf.ScannerDecl.Type.Params.List...)
	// attempt to infer the type from the pattern of the scanner function.
	if qf.QueryerFunction != nil {
		// do nothing, the function was specified.
	} else if queryPattern(pattern...) {
		qf.QueryerFunction = ast.NewIdent("Query")
	} else if queryRowPattern(pattern...) {
		qf.QueryerFunction = ast.NewIdent("QueryRow")
	} else {
		return genieql.NewErrGenerator(fmt.Errorf("a query function was not provided and failed to infer from the scanner function"))
	}

	return qf
}

type queryFunction struct {
	Name            string
	ScannerDecl     *ast.FuncDecl
	BuiltinQuery    string
	Queryer         ast.Expr
	QueryerName     string
	QueryerFunction *ast.Ident
	Parameters      []*ast.Field
}

func (t *queryFunction) Apply(options ...QueryFunctionOption) error {
	for _, opt := range options {
		if err := opt(t); err != nil {
			return err
		}
	}
	return nil
}

func (t queryFunction) Generate(dst io.Writer) error {
	type context struct {
		Name         string
		ScannerFunc  ast.Expr
		ScannerType  ast.Expr
		BuiltinQuery ast.Node
		Queryer      ast.Expr
		Parameters   []*ast.Field
	}

	var (
		tmpl            *template.Template
		parameters      []*ast.Field
		queryParameters []ast.Expr
		builtinQuery    ast.Decl
		query           *ast.CallExpr
	)

	t.Parameters = normalizeFieldNames(t.Parameters)

	queryFieldParam := astutil.Field(ast.NewIdent("string"), ast.NewIdent("query"))

	funcMap := template.FuncMap{
		"expr":      types.ExprString,
		"arguments": arguments,
		"printAST":  astPrint,
	}

	parameters = append(parameters, astutil.Field(t.Queryer, ast.NewIdent(t.QueryerName)))
	if t.BuiltinQuery == "" {
		parameters = append(parameters, queryFieldParam)
	} else {
		builtinQuery = genieql.QueryLiteral("query", t.BuiltinQuery)
	}

	parameters = append(parameters, t.Parameters...)

	queryParameters = append(queryParameters, astutil.MapFieldsToNameExpr(queryFieldParam)...)
	queryParameters = append(queryParameters, astutil.MapFieldsToNameExpr(t.Parameters...)...)

	// if we're dealing with an ellipsis parameter function
	// mark the CallExpr Ellipsis
	// this should only be the case when t.Parameters ends with
	// an ast.Ellipsis expression.
	// this allows for the creation of a generic function:
	// func F(q sql.DB, query, params ...interface{}) StaticExampleScanner
	query = &ast.CallExpr{
		Fun:      &ast.SelectorExpr{X: ast.NewIdent(t.QueryerName), Sel: t.QueryerFunction},
		Args:     queryParameters,
		Ellipsis: isEllipsis(t.Parameters),
	}

	ctx := context{
		Name:         t.Name,
		ScannerType:  t.ScannerDecl.Type.Results.List[0].Type,
		ScannerFunc:  t.ScannerDecl.Name,
		BuiltinQuery: builtinQuery,
		Parameters:   parameters,
		Queryer:      query,
	}

	tmpl = template.Must(template.New("query-function").Funcs(funcMap).Parse(queryFunc))
	return errors.Wrap(tmpl.Execute(dst, ctx), "failed to generate static scanner")
}

var queryPattern = astutil.TypePattern(astutil.ExprList("*sql.Rows", "error")...)
var queryRowPattern = astutil.TypePattern(astutil.Expr("*sql.Row"))

const queryFunc = `func {{.Name}}({{ .Parameters | arguments }}) {{ .ScannerType | expr }} {
	{{ if .BuiltinQuery -}}
	{{ .BuiltinQuery | printAST }}
	{{ end -}}
	return {{ .ScannerFunc | expr }}({{ .Queryer | expr }})
}
`

func isEllipsis(fields []*ast.Field) token.Pos {
	var (
		x ast.Expr
	)

	if len(fields) == 0 {
		return token.Pos(0)
	}

	x = fields[len(fields)-1].Type

	if _, isEllipsis := x.(*ast.Ellipsis); !isEllipsis {
		return token.Pos(0)
	}

	return token.Pos(1)
}

// notes: should be able to remove the queryer-function for general use cases by inspecting the return function.
// if a custom queryer-function was provided, use that.
// if it has the pattern (*sql.Row) then we use QueryRow.
// if it matches the pattern (*sql.Rows, error) then we use Query.
// if it doesn't match any of the above: error.
//
// genieql.options: [general] inlined-query="SELECT * FROM foo WHERE bar = $1 || bar = $2"
// type MyQueryFunction func(q sqlx.Queryer, param1, param2 int) StaticExampleScanner
// creates:
// func MyQueryFunction(q sqlx.Queryer, param1, param2 int) ExampleScanner {
// 	const query = `SELECT * FROM foo WHERE bar = $1 || bar = $2`
// 	return StaticExampleScanner(q.Query(query, param1, param2))
// }
//
// type MyQueryFunction func(q sqlx.Queryer, param1, param2 int) StaticExampleScanner
// creates:
// func MyQueryFunction func(q sqlx.Queryer, query string, param1, param2 int) ExampleScanner {
// 	return StaticExampleScanner(q.Query(query, param1, param2))
// }
//
// type MyQueryFunction func(q sqlx.Queryer) DynamicExampleScanner
// creates:
// func MyQueryFunction(q sqlx.Queryer, query string, params ...interface{}) ExampleScanner {
// 	return DynamicExampleScanner(q.Query(query, params...))
// }
//
// genieql.options: [general] queryer-function=QueryRow
// type MyQueryFunction func(q sqlx.Queryer) NewStaticRowExample
// creates:
// func MyQueryFunction(q sqlx.Queryer, query string, params ...interface{}) NewStaticRowExample {
// 	return NewStaticRowExample(q.QueryRow(query, params...))
// }