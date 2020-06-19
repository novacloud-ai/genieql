package functions

import (
	"database/sql"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate insert experimental batch-function -o postgresql.insert.batch.gen.go
// invoked by go generate @ functions/functions.go line 8

// NewExample4BatchInsertFunction creates a scanner that inserts a batch of
// records into the database.
func NewExample4BatchInsertFunction(q sqlx.Queryer, p ...Example4) Example4Scanner {
	return &example4BatchInsertFunction{
		q:         q,
		remaining: p,
	}
}

type example4BatchInsertFunction struct {
	q         sqlx.Queryer
	remaining []Example4
	scanner   Example4Scanner
}

func (t *example4BatchInsertFunction) Scan(dst *Example4) error {
	return t.scanner.Scan(dst)
}

func (t *example4BatchInsertFunction) Err() error {
	if t.scanner == nil {
		return nil
	}

	return t.scanner.Err()
}

func (t *example4BatchInsertFunction) Close() error {
	if t.scanner == nil {
		return nil
	}
	return t.scanner.Close()
}

func (t *example4BatchInsertFunction) Next() bool {
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

func (t *example4BatchInsertFunction) advance(q sqlx.Queryer, p ...Example4) (Example4Scanner, []Example4, bool) {
	switch len(p) {
	case 0:
		return nil, []Example4(nil), false
	case 1:
		const query = `INSERT INTO example4 ("created","email","id","updated") VALUES ($1,$2,$3,$4) RETURNING "created","email","id","updated"`
		exploder := func(p ...Example4) (r [4]interface{}, err error) {
			for idx, v := range p[:1] {
				var (
					c0 sql.NullTime
					c1 sql.NullString
					c2 sql.NullString
					c3 sql.NullTime
				)
				c0.Valid = true
				c0.Time = v.Created
				c1.Valid = true
				c1.String = v.Email
				c2.Valid = true
				c2.String = v.ID
				c3.Valid = true
				c3.Time = v.Updated
				r[idx*4+0], r[idx*4+1], r[idx*4+2], r[idx*4+3] = c0, c1, c2, c3
			}
			return r, nil
		}

		tmp, err := exploder(p...)

		if err != nil {
			return NewExample4ScannerStatic(nil, err), []Example4(nil), false
		}

		return NewExample4ScannerStatic(q.Query(query, tmp[:]...)), []Example4(nil), true
	case 2:
		const query = `INSERT INTO example4 ("created","email","id","updated") VALUES ($1,$2,$3,$4),($5,$6,$7,$8) RETURNING "created","email","id","updated"`
		exploder := func(p ...Example4) (r [8]interface{}, err error) {
			for idx, v := range p[:2] {
				var (
					c0 sql.NullTime
					c1 sql.NullString
					c2 sql.NullString
					c3 sql.NullTime
				)
				c0.Valid = true
				c0.Time = v.Created
				c1.Valid = true
				c1.String = v.Email
				c2.Valid = true
				c2.String = v.ID
				c3.Valid = true
				c3.Time = v.Updated
				r[idx*4+0], r[idx*4+1], r[idx*4+2], r[idx*4+3] = c0, c1, c2, c3
			}
			return r, nil
		}

		tmp, err := exploder(p...)

		if err != nil {
			return NewExample4ScannerStatic(nil, err), []Example4(nil), false
		}

		return NewExample4ScannerStatic(q.Query(query, tmp[:]...)), []Example4(nil), true
	case 3:
		const query = `INSERT INTO example4 ("created","email","id","updated") VALUES ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12) RETURNING "created","email","id","updated"`
		exploder := func(p ...Example4) (r [12]interface{}, err error) {
			for idx, v := range p[:3] {
				var (
					c0 sql.NullTime
					c1 sql.NullString
					c2 sql.NullString
					c3 sql.NullTime
				)
				c0.Valid = true
				c0.Time = v.Created
				c1.Valid = true
				c1.String = v.Email
				c2.Valid = true
				c2.String = v.ID
				c3.Valid = true
				c3.Time = v.Updated
				r[idx*4+0], r[idx*4+1], r[idx*4+2], r[idx*4+3] = c0, c1, c2, c3
			}
			return r, nil
		}

		tmp, err := exploder(p...)

		if err != nil {
			return NewExample4ScannerStatic(nil, err), []Example4(nil), false
		}

		return NewExample4ScannerStatic(q.Query(query, tmp[:]...)), []Example4(nil), true
	case 4:
		const query = `INSERT INTO example4 ("created","email","id","updated") VALUES ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12),($13,$14,$15,$16) RETURNING "created","email","id","updated"`
		exploder := func(p ...Example4) (r [16]interface{}, err error) {
			for idx, v := range p[:4] {
				var (
					c0 sql.NullTime
					c1 sql.NullString
					c2 sql.NullString
					c3 sql.NullTime
				)
				c0.Valid = true
				c0.Time = v.Created
				c1.Valid = true
				c1.String = v.Email
				c2.Valid = true
				c2.String = v.ID
				c3.Valid = true
				c3.Time = v.Updated
				r[idx*4+0], r[idx*4+1], r[idx*4+2], r[idx*4+3] = c0, c1, c2, c3
			}
			return r, nil
		}

		tmp, err := exploder(p...)

		if err != nil {
			return NewExample4ScannerStatic(nil, err), []Example4(nil), false
		}

		return NewExample4ScannerStatic(q.Query(query, tmp[:]...)), []Example4(nil), true
	default:
		const query = `INSERT INTO example4 ("created","email","id","updated") VALUES ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12),($13,$14,$15,$16),($17,$18,$19,$20) RETURNING "created","email","id","updated"`
		exploder := func(p ...Example4) (r [20]interface{}, err error) {
			for idx, v := range p[:5] {
				var (
					c0 sql.NullTime
					c1 sql.NullString
					c2 sql.NullString
					c3 sql.NullTime
				)
				c0.Valid = true
				c0.Time = v.Created
				c1.Valid = true
				c1.String = v.Email
				c2.Valid = true
				c2.String = v.ID
				c3.Valid = true
				c3.Time = v.Updated
				r[idx*4+0], r[idx*4+1], r[idx*4+2], r[idx*4+3] = c0, c1, c2, c3
			}
			return r, nil
		}

		tmp, err := exploder(p[:5]...)

		if err != nil {
			return NewExample4ScannerStatic(nil, err), []Example4(nil), false
		}

		return NewExample4ScannerStatic(q.Query(query, tmp[:]...)), p[5:], true
	}
}
