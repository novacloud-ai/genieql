package functions

import (
	"fmt"
	"go/ast"
	"go/types"
	"log"
	"strings"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/astutil"
	"bitbucket.org/jatone/genieql/generators"
	"github.com/pkg/errors"
)

const (
	defaultQueryParamName = "q"
	defaultQuery          = "query"
)

var queryRecordsPattern = astutil.TypePattern(astutil.ExprTemplateList("*sql.Rows", "error")...)
var queryUniquePattern = astutil.TypePattern(astutil.Expr("*sql.Row"))
var contextPattern = astutil.TypePattern(astutil.Expr("context.Context"))

// DetectScanner - extracts the scanner from the function definition.
// by convention the scanner is the first field in the result type.
func DetectScanner(ctx generators.Context, fnt *ast.FuncType) (r *ast.FuncDecl) {
	var (
		err error
	)

	if fnt.Results == nil || len(fnt.Results.List) == 0 {
		return nil
	}

	test := func(s string) bool {
		return s == types.ExprString(fnt.Results.List[0].Type)
	}

	util := genieql.NewSearcher(ctx.FileSet, ctx.CurrentPackage)

	if r, err = util.FindFunction(test); err != nil {
		log.Println("failed to find scanner")
		return nil
	}

	return r
}

// DetectContext - detects if a context.Context is being used.
// by convention the scanner is the first field in the inputs type.
func DetectContext(fnt *ast.FuncType) (r *ast.Field) {
	if fnt.Params == nil || len(fnt.Params.List) <= 1 {
		return nil
	}

	pattern := astutil.MapFieldsToTypExpr(fnt.Params.List[0])

	if contextPattern(pattern...) {
		return fnt.Params.List[0]
	}

	return nil
}

// compiler consumes a definition and returns a function declaration node.
type compiler interface {
	Compile(Definition) (*ast.FuncDecl, error)
}

// Query function compiler
type Query struct {
	generators.Context
	Query           ast.Expr
	ContextField    *ast.Field // is there a context field
	Queryer         ast.Expr   // the type of the queryer
	QueryerFunction *ast.Ident
	Scanner         *ast.FuncDecl
}

func (t Query) sanitizeFields(i *ast.Ident) *ast.Ident {
	switch i.Name {
	case defaultQueryParamName, defaultQuery:
		return ast.NewIdent("_genieql" + strings.Title(i.Name))
	}

	return i
}

// transform placeholder...
func (t Query) transform(inputs ...*ast.Field) (transforms []ast.Stmt, output []*ast.Field) {
	output = astutil.TransformFields(func(field *ast.Field) *ast.Field {
		log.Println("transforms for field", field.Names, field.Type)
		return field
	}, inputs...)
	return transforms, output
}

// Compile using the provided definition.
func (t Query) Compile(d Definition) (*ast.FuncDecl, error) {
	var (
		query        = astutil.Expr(defaultQuery)
		queryerIdent = ast.NewIdent(defaultQueryParamName)
		stmts        []ast.Stmt
	)
	defer func() {
		recovered := recover()
		if recovered == nil {
			return
		}

		if cause, ok := recovered.(error); ok {
			log.Println("panic", cause)
		}
	}()

	// basic validations
	if t.Scanner == nil {
		return nil, errors.Errorf("a scanner was not provided")
	}

	pattern := astutil.MapFieldsToTypExpr(t.Scanner.Type.Params.List...)

	// attempt to infer the type from the pattern of the scanner function.
	if t.QueryerFunction != nil {
		// do nothing, the function was specified.
	} else if queryRecordsPattern(pattern...) && t.ContextField != nil {
		t.QueryerFunction = ast.NewIdent("QueryContext")
	} else if queryUniquePattern(pattern...) && t.ContextField != nil {
		t.QueryerFunction = ast.NewIdent("QueryRowContext")
	} else if queryRecordsPattern(pattern...) {
		t.QueryerFunction = ast.NewIdent("Query")
	} else if queryUniquePattern(pattern...) {
		t.QueryerFunction = ast.NewIdent("QueryRow")
	} else {
		return nil, errors.Errorf("a query function was not provided and failed to infer from the scanner function parameter list")
	}

	// prevent name collisions.
	d.Signature.Params.List = generators.SanitizeFieldIdents(t.sanitizeFields, d.Signature.Params.List...)
	// TODO: generate input transformation statements if necessary.
	stmts, d.Signature.Params.List = t.transform(d.Signature.Params.List...)
	// TODO: generate input transformation statements if necessary.

	// setup function arguments.
	finputs := []*ast.Field{astutil.Field(t.Queryer, queryerIdent)}
	if t.ContextField != nil {
		finputs = []*ast.Field{t.ContextField, astutil.Field(t.Queryer, queryerIdent)}
	}

	// setup query inputs
	qinputs := []ast.Expr{}
	if t.ContextField != nil {
		qinputs = append(qinputs, astutil.MapFieldsToNameExpr(t.ContextField)...)
	}
	qinputs = append(qinputs, query)
	qinputs = append(qinputs, astutil.MapFieldsToNameExpr(d.Signature.Params.List...)...)

	// rewrite function parameters with the queryer and context
	d.Signature.Params.List = append(
		finputs,
		d.Signature.Params.List...,
	)
	d.Signature.Results = t.Scanner.Type.Results

	stmts = append([]ast.Stmt{&ast.DeclStmt{Decl: astutil.Const(types.ExprString(query), t.Query)}}, stmts...)
	stmts = append(stmts, astutil.Return(
		astutil.CallExpr(
			t.Scanner.Name,
			astutil.CallExpr(
				astutil.SelExpr(queryerIdent.Name, t.QueryerFunction.Name),
				qinputs...,
			),
		),
	))

	return combine(d, astutil.Block(stmts...)), nil
}

// Compile a definition using the provided compiler
func Compile(d Definition, c compiler) (*ast.FuncDecl, error) {
	return c.Compile(d)
}

// New function definition
func New(name string, signature *ast.FuncType, options ...Option) Definition {
	var (
		defaultComment = &ast.CommentGroup{
			List: []*ast.Comment{
				{Text: fmt.Sprintf("// %s generated by genieql", name)},
			},
		}
	)

	signature.Params.List = generators.NormalizeFieldNames(signature.Params.List...)

	return Definition{
		Name: name,
		Signature: &ast.FuncType{
			Params:  signature.Params,
			Results: signature.Results,
		},
		Comment: defaultComment,
	}.apply(options...)
}

// Definition of a function.
type Definition struct {
	Name      string            // name of the generated function
	Comment   *ast.CommentGroup // comment of the generated function.
	Signature *ast.FuncType     // signature of the generated function defining expected inputs and output.
}

func (t Definition) apply(options ...Option) Definition {
	for _, opt := range options {
		opt(&t)
	}

	return t
}

// Option options for building query functions.
type Option func(*Definition)

// OptionNoop do nothing
func OptionNoop(*Definition) {}

func combine(d Definition, b *ast.BlockStmt) (res *ast.FuncDecl) {
	return &ast.FuncDecl{
		Name: &ast.Ident{
			Name: d.Name,
		},
		Type: &ast.FuncType{
			Params:  d.Signature.Params,
			Results: d.Signature.Results,
		},
		Body: b,
	}
}
