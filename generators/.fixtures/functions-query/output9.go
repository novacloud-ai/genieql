package example

import (
	"database/sql"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// queryFunction9 generated by genieql
func queryFunction9(q sqlx.Queryer, query string, arg1 *StructA) ExampleScanner {
	var (
		c0 sql.NullInt64
		c1 sql.NullInt64
		c2 sql.NullInt64
		c3 sql.NullBool
		c4 sql.NullBool
		c5 sql.NullBool
		c6 sql.NullInt64
		c7 sql.NullBool
	)

	c0.Int64 = int64(arg1.A)
	c1.Int64 = int64(arg1.B)
	c2.Int64 = int64(arg1.C)
	c3.Bool = arg1.D
	c4.Bool = arg1.E
	c5.Bool = arg1.F
	c6.Int64 = int64(*arg1.G)
	c7.Bool = *arg1.H

	return StaticExampleScanner(q.Query(query, c0, c1, c2, c3, c4, c5, c6, c7))
}
