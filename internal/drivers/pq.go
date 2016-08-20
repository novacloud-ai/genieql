package drivers

import (
	"fmt"
	"go/ast"
	"go/types"

	"bitbucket.org/jatone/genieql"
)

// implements the lib/pq driver https://github.com/lib/pq

func init() {
	genieql.RegisterDriver("github.com/lib/pq", genieql.NewDriver(pqNullableTypes, pqLookupNullableType))
}

func pqNullableTypes(typ, from ast.Expr) (ast.Expr, bool) {
	var expr ast.Expr
	ok := true

	// if its not a starexpr its not nullable
	x, ok := typ.(*ast.StarExpr)
	if !ok {
		return typ, false
	}

	typeToExpr := func(selector string) ast.Expr {
		return mustParseExpr(fmt.Sprintf("%s.%s", types.ExprString(from), selector))
	}

	switch types.ExprString(x.X) {
	case timeExprString:
		expr = typeToExpr("Time")
	default:
		expr, ok = typ, false
	}

	return expr, ok
}

func pqLookupNullableType(typ ast.Expr) ast.Expr {
	// if its not a starexpr its not nullable
	x, ok := typ.(*ast.StarExpr)
	if !ok {
		return typ
	}

	switch types.ExprString(x.X) {
	case timeExprString:
		return mustParseExpr("pq.NullTime").(*ast.SelectorExpr)
	default:
		return typ
	}
}
