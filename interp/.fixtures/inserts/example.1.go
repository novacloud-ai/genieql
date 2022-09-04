package example

import (
	"context"
	"database/sql"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// InsertExample1StaticColumns generated by genieql
const InsertExample1StaticColumns = `a,b,c,d,e,f,g,h`

// InsertExample1Explode generated by genieql
func InsertExample1Explode(arg1 *StructA) ([]interface{}, error) {
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

	c0.Valid = true
	c0.Int64 = int64(arg1.A)

	c1.Valid = true
	c1.Int64 = int64(arg1.B)

	c2.Valid = true
	c2.Int64 = int64(arg1.C)

	c3.Valid = true
	c3.Bool = arg1.D

	c4.Valid = true
	c4.Bool = arg1.E

	c5.Valid = true
	c5.Bool = arg1.F

	c6.Valid = true
	c6.Int64 = int64(*arg1.G)

	c7.Valid = true
	c7.Bool = *arg1.H

	return []interface{}{c0, c1, c2, c3, c4, c5, c6, c7}, nil
}

// InsertExample1
func InsertExample1(ctx context.Context, q sqlx.Queryer, a StructA) ExampleScanner {
	const query = `INSERT INTO foo (a,b,c,d,e,f,g,h) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
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
	c0.Valid = true
	c0.Int64 = int64(a.A)
	c1.Valid = true
	c1.Int64 = int64(a.B)
	c2.Valid = true
	c2.Int64 = int64(a.C)
	c3.Valid = true
	c3.Bool = a.D
	c4.Valid = true
	c4.Bool = a.E
	c5.Valid = true
	c5.Bool = a.F
	c6.Valid = true
	c6.Int64 = int64(*a.G)
	c7.Valid = true
	c7.Bool = *a.H
	return NewExampleScannerStatic(q.QueryContext(ctx, query, c0, c1, c2, c3, c4, c5, c6, c7))
}
