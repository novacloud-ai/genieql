package generators

import (
	"bufio"
	"bytes"
	"go/ast"
	"go/build"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/serenize/snaker"
	"github.com/zieckey/goini"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/astutil"
)

func exprToArray(x ast.Expr) ast.Expr {
	return &ast.ArrayType{
		Elt: x,
	}
}

func fieldToType(f *ast.Field) ast.Expr {
	return f.Type
}

func fieldToNames(f *ast.Field) []*ast.Ident {
	return f.Names
}

// utility function that converts a set of ast.Field into
// a string representation of a function's arguments.
func arguments(fields []*ast.Field) string {
	xtransformer := func(x ast.Expr) ast.Expr {
		return x
	}
	return _arguments(xtransformer, fields)
}

func argumentsAsPointers(fields []*ast.Field) string {
	xtransformer := func(x ast.Expr) ast.Expr {
		return &ast.StarExpr{X: x}
	}
	return _arguments(xtransformer, fields)
}

func _arguments(xtransformer func(ast.Expr) ast.Expr, fields []*ast.Field) string {
	result := []string{}
	for _, field := range fields {
		result = append(result,
			strings.Join(
				astutil.MapExprToString(astutil.MapIdentToExpr(field.Names...)...),
				", ",
			)+" "+types.ExprString(xtransformer(field.Type)))
	}
	return strings.Join(result, ", ")
}

// normalizes the names of the field.
func normalizeFieldNames(fields ...*ast.Field) []*ast.Field {
	result := make([]*ast.Field, 0, len(fields))
	for _, field := range fields {
		result = append(result, astutil.Field(field.Type, normalizeIdent(field.Names)...))
	}
	return result
}

// normalize's the idents.
func normalizeIdent(idents []*ast.Ident) []*ast.Ident {
	result := make([]*ast.Ident, 0, len(idents))

	for _, ident := range idents {
		n := ident.Name
		if !strings.Contains(n, "_") {
			n = snaker.CamelToSnake(ident.Name)
		}
		result = append(result, ast.NewIdent(toPrivate(n)))
	}

	return result
}

func toPrivate(s string) string {
	parts := strings.SplitN(s, "_", 2)
	switch len(parts) {
	case 2:
		return strings.ToLower(parts[0]) + snaker.SnakeToCamel(strings.ToLower(parts[1]))
	default:
		return strings.ToLower(s)
	}
}

func astPrint(n ast.Node) (string, error) {
	if n == nil {
		return "", nil
	}

	dst := bytes.NewBuffer([]byte{})
	fset := token.NewFileSet()
	err := printer.Fprint(dst, fset, n)

	return dst.String(), errors.Wrap(err, "failure to print ast")
}

// OptionsFromCommentGroup parses a configuration and converts it into an array of options.
func OptionsFromCommentGroup(comments *ast.CommentGroup) (*goini.INI, error) {
	const magicPrefix = `genieql.options:`

	ini := goini.New()
	ini.SetParseSection(true)
	scanner := bufio.NewScanner(strings.NewReader(comments.Text()))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if !strings.HasPrefix(text, magicPrefix) {
			continue
		}

		text = strings.TrimSpace(strings.TrimPrefix(text, magicPrefix))

		if err := ini.Parse([]byte(text), "||", "="); err != nil {
			return nil, errors.Wrap(err, "failed to parse comment configuration")
		}
	}

	return ini, nil
}

func areArrayType(xs ...ast.Expr) bool {
	for _, x := range xs {
		if _, ok := x.(*ast.ArrayType); !ok {
			return false
		}
	}
	return true
}

func extractArrayInfo(x *ast.ArrayType) (int, ast.Expr, error) {
	var (
		err error
		max int
		ok  bool
		lit *ast.BasicLit
	)
	if lit, ok = x.Len.(*ast.BasicLit); !ok {
		return max, x.Elt, errors.New("expected a basic literal for the array")
	}

	if lit.Kind != token.INT {
		return max, x.Elt, errors.New("expected the basic literal of the array to be of type integer")
	}

	if max, err = strconv.Atoi(lit.Value); err != nil {
		return max, x.Elt, errors.Wrap(err, "failed to convert the array size to an integer")
	}

	return max, x.Elt, nil
}

func selectType(x ast.Expr) bool {
	_, ok := x.(*ast.SelectorExpr)
	return ok
}

func allBuiltinTypes(xs ...ast.Expr) bool {
	for _, x := range xs {
		if !builtinType(x) {
			return false
		}
	}

	return true
}

func builtinType(x ast.Expr) bool {
	name := types.ExprString(x)
	for _, t := range types.Typ {
		if name == t.Name() {
			return true
		}
	}

	switch name {
	case "interface{}":
		return true
	case "time.Time":
		return true
	default:
		return false
	}
}

// builtinParam converts a *ast.Field that represents a builtin type
// (time.Time,int,float,bool, etc) into an array of ColumnMap.
func builtinParam(param *ast.Field) ([]genieql.ColumnMap, error) {
	columns := make([]genieql.ColumnMap, 0, len(param.Names))
	for _, name := range param.Names {
		columns = append(columns, genieql.ColumnMap{
			Name:   name.Name,
			Type:   &ast.StarExpr{X: param.Type},
			Dst:    &ast.StarExpr{X: name},
			PtrDst: false,
		})
	}
	return columns, nil
}

func unwrapExpr(x ast.Expr) ast.Expr {
	switch real := x.(type) {
	case *ast.Ellipsis:
		return real.Elt
	case *ast.StarExpr:
		return real.X
	default:
		return x
	}
}
func packageName(pkg *build.Package, x ast.Expr) string {
	switch x := x.(type) {
	case *ast.SelectorExpr:
		// TODO
		log.Println("imports", types.ExprString(x.X), x.Sel.Name, pkg.Imports)
		panic("unimplemented code path: currently structures from other packages are not supported")
	default:
		return pkg.ImportPath
	}
}
