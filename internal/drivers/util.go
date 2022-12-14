package drivers

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"bitbucket.org/jatone/genieql/astutil"
)

const (
	boolExprString    = "bool"
	intExprString     = "int"
	intArrExpr        = "[]int"
	int16ExprString   = "int16"
	int32ExprString   = "int32"
	int64ExprString   = "int64"
	stringExprString  = "string"
	stringArrExpr     = "[]string"
	float32ExprString = "float32"
	float64ExprString = "float64"
	timeExprString    = "time.Time"
	durationExpr      = "time.Duration"
	ipExpr            = "net.IP"
	macExpr           = "net.HardwareAddr"
	cidrExpr          = "net.IPNet"
	cidrArrayExpr     = "[]net.IPNet"
	bytesExpr         = "[]byte"
)

func typeToExpr(from ast.Expr, selector string) ast.Expr {
	return astutil.MustParseExpr(token.NewFileSet(), fmt.Sprintf("%s.%s", types.ExprString(from), selector))
}

func castedTypeToExpr(castType, expr ast.Expr) ast.Expr {
	return astutil.Expr(fmt.Sprintf("%s(%s)", types.ExprString(castType), types.ExprString(expr)))
}
