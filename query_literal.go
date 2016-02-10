package genieql

import (
	"fmt"
	"go/ast"
	"go/token"
)

// QueryLiteral creates a const with the given name for the provided
// query.
func QueryLiteral(name, query string) *ast.GenDecl {
	return &ast.GenDecl{
		Tok: token.CONST,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{
					&ast.Ident{
						Name: name,
						Obj: &ast.Object{
							Kind: ast.Con,
							Name: name,
						},
					},
				},
				Values: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: fmt.Sprintf("`%s`", query),
					},
				},
			},
		},
	}
}
