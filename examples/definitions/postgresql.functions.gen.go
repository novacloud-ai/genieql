package definitions

import (
	"database/sql"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental functions types -o postgresql.functions.gen.go
// invoked by go generate @ definitions/example.go line 7

// customQueryFunction generated by genieql
func customQueryFunction(q sqlx.Queryer, query string, x1, x2, x3 int) ProfileScanner {
	var (
		c0 sql.NullInt64
		c1 sql.NullInt64
		c2 sql.NullInt64
	)

	c0.Int64 = int64(x1)
	c1.Int64 = int64(x2)
	c2.Int64 = int64(x3)

	return NewProfileScannerDynamic(q.Query(query, c0, c1, c2))
}

// customQueryFunction2 generated by genieql
func customQueryFunction2(q sqlx.Queryer, x1, x2, x3 int) ProfileScanner {
	var query = query1
	var (
		c0 sql.NullInt64
		c1 sql.NullInt64
		c2 sql.NullInt64
	)

	c0.Int64 = int64(x1)
	c1.Int64 = int64(x2)
	c2.Int64 = int64(x3)

	return NewProfileScannerDynamic(q.Query(query, c0, c1, c2))
}
