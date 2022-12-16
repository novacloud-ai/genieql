package example

import (
	"context"
	"database/sql"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// BatchInsertExample1 generated by genieql
func NewBatchInsertExample1(ctx context.Context, q sqlx.Queryer, a ...StructA) ExampleScanner {
	return &batchInsertExample1{ctx: ctx, q: q, remaining: a}
}

type batchInsertExample1 struct {
	ctx       context.Context
	q         sqlx.Queryer
	remaining []StructA
	scanner   ExampleScanner
}

func (t *batchInsertExample1) Scan(a *StructA) error {
	return t.scanner.Scan(a)
}

func (t *batchInsertExample1) Err() error {
	if t.scanner == nil {
		return nil
	}
	return t.scanner.Err()
}

func (t *batchInsertExample1) Close() error {
	if t.scanner == nil {
		return nil
	}
	return t.scanner.Close()
}

func (t *batchInsertExample1) Next() bool {
	var advanced bool
	if t.scanner != nil && t.scanner.Next() {
		return true
	}
	if len(t.remaining) > 0 && t.Close() == nil {
		t.scanner, t.remaining, advanced = t.advance(t.remaining...)
		return advanced && t.scanner.Next()
	}
	return false
}

func (t *batchInsertExample1) advance(a ...StructA) (ExampleScanner, []StructA, bool) {
	transform := func(a StructA) (c0 sql.NullInt64, c1 sql.NullInt64, c2 sql.NullInt64, c3 sql.NullBool, c4 sql.NullBool, c5 sql.NullBool, c6 sql.NullInt64, c7 sql.NullBool, err error) {
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
		return c0, c1, c2, c3, c4, c5, c6, c7, nil
	}
	switch len(a) {
	case 0:
		return nil, []StructA(nil), false
	case 1:
		const query = `INSERT INTO foo (a,b,c,d,e,f,g,h) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING a,b,c,d,e,f,g,h`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullInt64
			r0c2 sql.NullInt64
			r0c3 sql.NullBool
			r0c4 sql.NullBool
			r0c5 sql.NullBool
			r0c6 sql.NullInt64
			r0c7 sql.NullBool
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, r0c5, r0c6, r0c7, err = transform(a[0]); err != nil {
			return NewExampleScannerStatic(nil, err), []StructA(nil), false
		}
		return NewExampleScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r0c5, r0c6, r0c7), a[1:], true)
	default:
		const query = `INSERT INTO foo (a,b,c,d,e,f,g,h) VALUES ($1,$2,$3,$4,$5,$6,$7,$8),($8,$9,$10,$11,$12,$13,$14,$15) RETURNING a,b,c,d,e,f,g,h`
		var (
			r0c0 sql.NullInt64
			r0c1 sql.NullInt64
			r0c2 sql.NullInt64
			r0c3 sql.NullBool
			r0c4 sql.NullBool
			r0c5 sql.NullBool
			r0c6 sql.NullInt64
			r0c7 sql.NullBool
			r1c0 sql.NullInt64
			r1c1 sql.NullInt64
			r1c2 sql.NullInt64
			r1c3 sql.NullBool
			r1c4 sql.NullBool
			r1c5 sql.NullBool
			r1c6 sql.NullInt64
			r1c7 sql.NullBool
			err  error
		)
		if r0c0, r0c1, r0c2, r0c3, r0c4, r0c5, r0c6, r0c7, err = transform(a[0]); err != nil {
			return NewExampleScannerStatic(nil, err), []StructA(nil), false
		}
		if r1c0, r1c1, r1c2, r1c3, r1c4, r1c5, r1c6, r1c7, err = transform(a[1]); err != nil {
			return NewExampleScannerStatic(nil, err), []StructA(nil), false
		}
		return NewExampleScannerStatic(t.q.QueryContext(t.ctx, query, r0c0, r0c1, r0c2, r0c3, r0c4, r0c5, r0c6, r0c7, r1c0, r1c1, r1c2, r1c3, r1c4, r1c5, r1c6, r1c7)), []StructA(nil), false
	}
}
