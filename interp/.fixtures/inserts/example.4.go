package example

import (
	"context"
	"database/sql"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// InsertExample4StaticColumns generated by genieql
const InsertExample4StaticColumns = `DEFAULT,c,d,e,f,g,h`

// InsertExample4Explode generated by genieql
func InsertExample4Explode(arg1 *StructA) ([]interface{}, error) {
	var (
		c0 sql.NullInt64
		c1 sql.NullBool
		c2 sql.NullBool
		c3 sql.NullBool
		c4 sql.NullInt64
		c5 sql.NullBool
	)

	c0.Valid = true
	c0.Int64 = int64(arg1.C)

	c1.Valid = true
	c1.Bool = arg1.D

	c2.Valid = true
	c2.Bool = arg1.E

	c3.Valid = true
	c3.Bool = arg1.F

	c4.Valid = true
	c4.Int64 = int64(*arg1.G)

	c5.Valid = true
	c5.Bool = *arg1.H

	return []interface{}{c0, c1, c2, c3, c4, c5}, nil
}

// InsertExample4 generated by genieql
func InsertExample4(ctx context.Context, q sqlx.Queryer, a StructA) ExampleScanner {
	const query = `INSERT INTO foo (b,c,d,e,f,g,h) VALUES (DEFAULT,$1,$2,$3,$4,$5,$6)`
	var (
		c0 sql.NullInt64
		c1 sql.NullBool
		c2 sql.NullBool
		c3 sql.NullBool
		c4 sql.NullInt64
		c5 sql.NullBool
	)
	c0.Valid = true
	c0.Int64 = int64(a.C)
	c1.Valid = true
	c1.Bool = a.D
	c2.Valid = true
	c2.Bool = a.E
	c3.Valid = true
	c3.Bool = a.F
	c4.Valid = true
	c4.Int64 = int64(*a.G)
	c5.Valid = true
	c5.Bool = *a.H
	return NewExampleScannerStatic(q.QueryContext(ctx, query, c0, c1, c2, c3, c4, c5))
}