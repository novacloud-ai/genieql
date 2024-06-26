package example

import (
	"database/sql"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// queryFunction2 generated by genieql
func queryFunction2(q sqlx.Queryer, query string, arg1 int) ExampleScanner {
	var (
		c0 sql.NullInt64
	)

	c0.Valid = true
	c0.Int64 = int64(arg1)

	return StaticExampleScanner(q.Query(query, c0))
}
