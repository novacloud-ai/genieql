package example

import (
	"database/sql"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// NewBatchFunction3 creates a scanner that inserts a batch of
// records into the database.
func NewBatchFunction3(q sqlx.Queryer, v ...custom) ExampleScanner {
	return &batchFunction3{
		q:         q,
		remaining: v,
	}
}

type batchFunction3 struct {
	q         sqlx.Queryer
	remaining []custom
	scanner   ExampleScanner
}

func (t *batchFunction3) Scan(dst *custom) error {
	return t.scanner.Scan(dst)
}

func (t *batchFunction3) Err() error {
	if t.scanner == nil {
		return nil
	}

	return t.scanner.Err()
}

func (t *batchFunction3) Close() error {
	if t.scanner == nil {
		return nil
	}
	return t.scanner.Close()
}

func (t *batchFunction3) Next() bool {
	var (
		advanced bool
	)

	if t.scanner != nil && t.scanner.Next() {
		return true
	}

	// advance to the next check
	if len(t.remaining) > 0 && t.Close() == nil {
		t.scanner, t.remaining, advanced = t.advance(t.q, t.remaining...)
		return advanced && t.scanner.Next()
	}

	return false
}

func (t *batchFunction3) advance(q sqlx.Queryer, v ...custom) (ExampleScanner, []custom, bool) {
	switch len(v) {
	case 0:
		return nil, []custom(nil), false
	case 1:
		const query = `QUERY 1`
		exploder := func(v ...custom) (r [3]interface{}) {
			for idx, v := range v[:1] {
				var (
					c0 sql.NullInt64
					c1 sql.NullInt64
					c2 sql.NullInt64
				)
				c0.Int64 = int64(v.A)
				c1.Int64 = int64(v.B)
				c2.Int64 = int64(v.C)
				r[idx*3+0], r[idx*3+1], r[idx*3+2] = c0, c1, c2
			}
			return r
		}
		tmp := exploder(v...)
		return StaticExampleScanner(q.Query(query, tmp[:]...)), []custom(nil), true
	case 2:
		const query = `QUERY 2`
		exploder := func(v ...custom) (r [6]interface{}) {
			for idx, v := range v[:2] {
				var (
					c0 sql.NullInt64
					c1 sql.NullInt64
					c2 sql.NullInt64
				)
				c0.Int64 = int64(v.A)
				c1.Int64 = int64(v.B)
				c2.Int64 = int64(v.C)
				r[idx*3+0], r[idx*3+1], r[idx*3+2] = c0, c1, c2
			}
			return r
		}
		tmp := exploder(v...)
		return StaticExampleScanner(q.Query(query, tmp[:]...)), []custom(nil), true
	default:
		const query = `QUERY 3`
		exploder := func(v ...custom) (r [9]interface{}) {
			for idx, v := range v[:3] {
				var (
					c0 sql.NullInt64
					c1 sql.NullInt64
					c2 sql.NullInt64
				)
				c0.Int64 = int64(v.A)
				c1.Int64 = int64(v.B)
				c2.Int64 = int64(v.C)
				r[idx*3+0], r[idx*3+1], r[idx*3+2] = c0, c1, c2
			}
			return r
		}
		tmp := exploder(v[:3]...)
		return StaticExampleScanner(q.Query(query, tmp[:]...)), v[3:], true
	}
}
