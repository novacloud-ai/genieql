package generators

import (
	"go/build"
	"go/token"

	"bitbucket.org/jatone/genieql"
)

// Generators generate schema and configuration for testing.
//go:generate dropdb --if-exists -U postgres genieql_test_template
//go:generate createdb -U postgres genieql_test_template
//go:generate psql -1 -f structure.sql genieql_test_template
//go:generate genieql bootstrap --queryer=sqlx.Queryer --driver=github.com/lib/pq --output-file=generators-test.config postgres://$USER@localhost:5432/genieql_test_template?sslmode=disable

// Context - context for generators
type Context struct {
	CurrentPackage *build.Package
	FileSet        *token.FileSet
	Configuration  genieql.Configuration
	Dialect        genieql.Dialect
}

func reserved(s string) bool {
	switch s {
	case "type":
		return true
	case "func":
		return true
	case "default":
		return true
	default:
		return false
	}
}
