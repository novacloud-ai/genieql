package example

import (
	"context"
	"database/sql"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// InsertExample3StaticColumns generated by genieql
const InsertExample3StaticColumns = `DEFAULT,b,c,d,e,f,g,h`

// InsertExample3Explode generated by genieql
func InsertExample3Explode(a *StructA) ([]interface{}, error) {
	var (
		c0 sql.NullInt64 // b
		c1 sql.NullInt64 // c
		c2 sql.NullBool  // d
		c3 sql.NullBool  // e
		c4 sql.NullBool  // f
		c5 sql.NullInt64 // g
		c6 sql.NullBool  // h
	)

	c0.Valid = true
	c0.Int64 = int64(a.B)

	c1.Valid = true
	c1.Int64 = int64(a.C)

	c2.Valid = true
	c2.Bool = a.D

	c3.Valid = true
	c3.Bool = a.E

	c4.Valid = true
	c4.Bool = a.F

	c5.Valid = true
	c5.Int64 = int64(*a.G)

	c6.Valid = true
	c6.Bool = *a.H

	return []interface{}{c0, c1, c2, c3, c4, c5, c6}, nil
}

// InsertExample3 generated by genieql
func InsertExample3(ctx context.Context, q sqlx.Queryer, a StructA) ExampleScanner {
	const query = `INSERT INTO foo (a,b,c,d,e,f,g,h) VALUES (DEFAULT,$1,$2,$3,$4,$5,$6,$7) RETURNING a,b,c,d,e,f,g,h`
	var (
		c0 sql.NullInt64 // b
		c1 sql.NullInt64 // c
		c2 sql.NullBool  // d
		c3 sql.NullBool  // e
		c4 sql.NullBool  // f
		c5 sql.NullInt64 // g
		c6 sql.NullBool
	)
	c0.Valid = true
	c0.Int64 = int64(a.B)
	c1.Valid = true
	c1.Int64 = int64(a.C)
	c2.Valid = true
	c2.Bool = a.D
	c3.Valid = true
	c3.Bool = a.E
	c4.Valid = true
	c4.Bool = a.F
	c5.Valid = true
	c5.Int64 = int64(*a.G)
	c6.Valid = true
	c6.Bool = *a.H // h
	return NewExampleScannerStatic(q.QueryContext(ctx, query, c0, c1, c2, c3, c4, c5, c6))
}