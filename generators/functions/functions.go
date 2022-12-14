package functions

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"io"
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
		log.Println("missing function results")
		return nil
	}

	test := func(s string) bool {
		return s == types.ExprString(fnt.Results.List[0].Type)
	}

	util := genieql.NewSearcher(ctx.FileSet, ctx.CurrentPackage)

	if r, err = util.FindFunction(test); err != nil {
		log.Println("failed to find scanner", types.ExprString(fnt.Results.List[0].Type))
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

	pattern := astutil.MapFieldsToTypeExpr(fnt.Params.List[0])

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
	Transforms      []ast.Stmt
	QueryInputs     []ast.Expr
}

func SanitizeQueryIdents(i *ast.Ident) *ast.Ident {
	switch i.Name {
	case defaultQueryParamName, defaultQuery:
		return ast.NewIdent("_genieql_" + i.Name)
	}

	return i
}

func QueryInputsFromFields(inputs ...*ast.Field) (output []ast.Expr) {
	for _, i := range inputs {
		output = append(output, astutil.MapFieldsToNameExpr(i)...)
	}
	return output
}

func ScannerErrorHandling(sid *ast.Ident) func(local string) ast.Node {
	return func(local string) ast.Node {
		return astutil.Return(
			astutil.CallExpr(
				&ast.SelectorExpr{
					X:   astutil.CallExpr(sid, ast.NewIdent("nil")),
					Sel: ast.NewIdent("Err"),
				},
				ast.NewIdent(local),
			),
		)
	}
}
func QueryLiteralColumnMapReplacer(ctx generators.Context, columns ...genieql.ColumnMap) *strings.Replacer {
	replacements := []string{}
	cidx := ctx.Dialect.ColumnValueTransformer()
	for _, c := range columns {
		dst := types.ExprString(astutil.DereferencedIdent(c.Dst))
		// log.Println(c.Name, "->", dst, spew.Sdump(c.Dst), spew.Sdump(astutil.DereferencedIdent(c.Dst)))
		replacements = append(
			replacements,
			fmt.Sprintf("{%s.query.input}", dst), cidx.Transform(c.ColumnInfo),
		)
	}
	// log.Println("REPLACEMENTS")
	// for i := 0; i < len(replacements); i = i + 2 {
	// 	log.Println(replacements[i], "->", replacements[i+1])
	// }
	return strings.NewReplacer(replacements...)
}

func ColumnMapFromFields(ctx generators.Context, inputs ...*ast.Field) (rcmaps []genieql.ColumnMap, err error) {
	for _, input := range inputs {
		for _, name := range input.Names {
			var (
				cmaps []genieql.ColumnMap
			)

			if cmaps, err = generators.MapField(ctx, astutil.Field(input.Type, name)); err != nil {
				return rcmaps, errors.Wrapf(
					err,
					"failed to map columns for: %s:%s",
					ctx.CurrentPackage.Name, types.ExprString(input.Type),
				)
			}

			rcmaps = append(rcmaps, cmaps...)
		}
	}

	return rcmaps, nil
}

// Compile using the provided definition.
func (t Query) Compile(d Definition) (_ *ast.FuncDecl, err error) {
	var (
		query        = astutil.Expr(defaultQuery)
		queryerIdent = ast.NewIdent(defaultQueryParamName)
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

	pattern := astutil.MapFieldsToTypeExpr(t.Scanner.Type.Params.List...)

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
	d.Signature.Params.List = generators.SanitizeFieldIdents(SanitizeQueryIdents, d.Signature.Params.List...)

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
	if len(t.QueryInputs) == 0 {
		qinputs = append(qinputs, QueryInputsFromFields(d.Signature.Params.List...)...)
	} else {
		qinputs = append(qinputs, t.QueryInputs...)
	}

	// rewrite function parameters with the queryer and context
	d.Signature.Params.List = append(
		finputs,
		d.Signature.Params.List...,
	)
	d.Signature.Results = t.Scanner.Type.Results

	stmts := []ast.Stmt{
		astutil.ConstDecl(types.ExprString(query), t.Query),
	}

	if len(t.Transforms) > 0 {
		stmts = append(stmts, t.Transforms...)
	}

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

func NewFn(body ...ast.Stmt) Body {
	return Body(body)
}

// Body function compile combinues a sequence of statements with the provided definition.
type Body []ast.Stmt

// Compile using the provided definition.
func (t Body) Compile(d Definition) (_ *ast.FuncDecl, err error) {
	return combine(d, astutil.Block(t...)), nil
}

// Compile a definition using the provided compiler
func Compile(d Definition, c compiler) (*ast.FuncDecl, error) {
	return c.Compile(d)
}

// CompileInto the provided io.Writer
func CompileInto(dst io.Writer, d Definition, c compiler) (err error) {
	var (
		n ast.Node
	)

	if n, err = Compile(d, c); err != nil {
		return err
	}

	return format.Node(dst, token.NewFileSet(), n)
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
	Name      string // name of the generated function
	Recv      *ast.FieldList
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

func OptionRecv(r *ast.FieldList) Option {
	return func(d *Definition) {
		d.Recv = r
	}
}

func combine(d Definition, b *ast.BlockStmt) (res *ast.FuncDecl) {
	return &ast.FuncDecl{
		Recv: d.Recv,
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
