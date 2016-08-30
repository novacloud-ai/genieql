package scanner

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"bitbucket.org/jatone/genieql"
	"bitbucket.org/jatone/genieql/astutil"
)

// BuildScannerInterface takes in a name and a set of parameters
// for the scan method, outputs a ast.Decl representing the scanner interface.
func BuildScannerInterface(name string, scannerParams ...*ast.Field) ast.Decl {
	return interfaceDeclaration(
		&ast.Ident{Name: name},
		funcDeclarationField(
			&ast.Ident{Name: "Scan"},
			&ast.FieldList{List: scannerParams},          // parameters
			&ast.FieldList{List: unnamedFields("error")}, // returns
		),
	)
}

// BuildRowsScannerInterface takes in a name and a set of parameters
// for the scan method, output a ast.Decl.
func BuildRowsScannerInterface(name string, scannerParams ...*ast.Field) ast.Decl {
	return interfaceDeclaration(
		&ast.Ident{Name: name},
		funcDeclarationField(
			&ast.Ident{Name: "Scan"},
			&ast.FieldList{List: scannerParams},          // parameters
			&ast.FieldList{List: unnamedFields("error")}, // returns
		),
		funcDeclarationField(
			&ast.Ident{Name: "Next"},
			nil, // no parameters
			&ast.FieldList{List: unnamedFields("bool")}, // returns
		),
		funcDeclarationField(
			&ast.Ident{Name: "Close"},
			nil, // no parameters
			&ast.FieldList{List: unnamedFields("error")}, // returns
		),
		funcDeclarationField(
			&ast.Ident{Name: "Err"},
			nil, // no parameters
			&ast.FieldList{List: unnamedFields("error")}, // returns
		),
	)
}

// NewScannerFunc structure that builds the function to get a scanner
// after executing a query.
type NewScannerFunc struct {
	ScannerName    string
	InterfaceName  string
	ErrScannerName string
}

// Build - generates a function declaration for building the scanner.
func (t NewScannerFunc) Build() *ast.FuncDecl {
	name := ast.NewIdent(fmt.Sprintf("New%s", t.ScannerName))
	rowsParam := typeDeclarationField(astutil.Expr("*sql.Rows"), ast.NewIdent("rows"))
	errParam := typeDeclarationField(&ast.Ident{Name: "error"}, ast.NewIdent("err"))
	result := unnamedFields(t.InterfaceName)
	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X: &ast.Ident{
						Name: "err",
					},
					Op: token.NEQ,
					Y: &ast.Ident{
						Name: "nil",
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CompositeLit{
									Type: &ast.Ident{Name: t.ErrScannerName},
									Elts: []ast.Expr{
										&ast.KeyValueExpr{
											Key: &ast.Ident{
												Name: "err",
											},
											Value: &ast.Ident{
												Name: "err",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.Ident{Name: t.ScannerName},
						Elts: []ast.Expr{
							&ast.KeyValueExpr{
								Key:   ast.NewIdent("Rows"),
								Value: ast.NewIdent("rows"),
							},
						},
					},
				},
			},
		},
	}
	return funcDecl(nil, name, []*ast.Field{rowsParam, errParam}, result, body)
}

// NewRowScannerFunc structure that builds the function to get a scanner after
// executing a query row.
type NewRowScannerFunc struct {
	ScannerName    string
	InterfaceName  string
	ErrScannerName string
}

// Build - generates a function declaration for building the scanner.
func (t NewRowScannerFunc) Build() *ast.FuncDecl {
	name := ast.NewIdent(fmt.Sprintf("New%s", t.ScannerName))
	rowsParam := typeDeclarationField(astutil.Expr("*sql.Row"), ast.NewIdent("row"))
	result := unnamedFields(t.InterfaceName)
	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.Ident{Name: t.ScannerName},
						Elts: []ast.Expr{
							&ast.KeyValueExpr{
								Key:   ast.NewIdent("row"),
								Value: ast.NewIdent("row"),
							},
						},
					},
				},
			},
		},
	}
	return funcDecl(nil, name, []*ast.Field{rowsParam}, result, body)
}

// Functions responsible for generating the functions
// associated with the scanner.
type Functions struct {
	Parameters []*ast.Field
}

// Generate return a list of ast Declarations representing the functions of the scanner.
// parameters:
// name - represents the type of the scanner that acts as the receiver for the function.
func (t Functions) Generate(name string, scan, err, close *ast.BlockStmt) []ast.Decl {
	scanFunc := scanFunctionBuilder(name, t.Parameters, scan)

	errFunc := errFuncBuilder(name, t.Parameters, err)

	closeFunc := closeFuncBuilder(name, t.Parameters, close)

	return []ast.Decl{scanFunc, errFunc, closeFunc}
}

func scanFunctionBuilder(name string, params []*ast.Field, body *ast.BlockStmt) ast.Decl {
	return funcDecl(
		&ast.Ident{Name: name},
		&ast.Ident{Name: "Scan"},
		params,
		unnamedFields("error"),
		body,
	)
}

func errFuncBuilder(name string, params []*ast.Field, body *ast.BlockStmt) ast.Decl {
	return funcDecl(
		&ast.Ident{Name: name},
		&ast.Ident{Name: "Err"},
		nil, // no parameters
		unnamedFields("error"),
		body,
	)
}

func closeFuncBuilder(name string, params []*ast.Field, body *ast.BlockStmt) ast.Decl {
	return funcDecl(
		&ast.Ident{Name: name},
		&ast.Ident{Name: "Close"},
		nil, // no parameters
		unnamedFields("error"),
		body,
	)
}

func nextFuncBuilder(name string, body *ast.BlockStmt) ast.Decl {
	return funcDecl(
		&ast.Ident{Name: name},
		&ast.Ident{Name: "Next"},
		nil, // no parameters
		unnamedFields("bool"),
		body,
	)
}

func columnMapToQuery(columnMaps ...genieql.ColumnMap) string {
	result := make([]string, 0, len(columnMaps))
	for _, columnMap := range columnMaps {
		result = append(result, columnMap.ColumnName)
	}

	return strings.Join(result, ",")
}
