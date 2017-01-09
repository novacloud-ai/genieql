package definitions

import (
	"errors"

	"bitbucket.org/jatone/genieql/sqlx"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate insert experimental batch-function -o postgresql.insert.batch.gen.go
// invoked by go generate @ definitions/example.go line 8

func example1BatchInsertFunction(q sqlx.Queryer, p ...Example1) (Example1Scanner, []Example1) {
	switch len(p) {
	case 0:
		return NewExample1ScannerStatic(nil, errors.New("need at least 1 value to execute a batch query")), p
	case 1:
		const query = `QUERY 1`
		exploder := func(p ...Example1) (r [5]interface{}) {
			for idx, v := range p[:1] {
				r[idx*5+0], r[idx*5+1], r[idx*5+2], r[idx*5+3], r[idx*5+4] = v.CreatedAt, v.ID, v.TextField, v.UpdatedAt, v.UUIDField
			}
			return
		}
		tmp := exploder(p...)
		return NewExample1ScannerStatic(q.Query(query, tmp[:]...)), p[len(p)-1:]
	case 2:
		const query = `QUERY 2`
		exploder := func(p ...Example1) (r [10]interface{}) {
			for idx, v := range p[:2] {
				r[idx*5+0], r[idx*5+1], r[idx*5+2], r[idx*5+3], r[idx*5+4] = v.CreatedAt, v.ID, v.TextField, v.UpdatedAt, v.UUIDField
			}
			return
		}
		tmp := exploder(p...)
		return NewExample1ScannerStatic(q.Query(query, tmp[:]...)), p[len(p)-1:]
	case 3:
		const query = `QUERY 3`
		exploder := func(p ...Example1) (r [15]interface{}) {
			for idx, v := range p[:3] {
				r[idx*5+0], r[idx*5+1], r[idx*5+2], r[idx*5+3], r[idx*5+4] = v.CreatedAt, v.ID, v.TextField, v.UpdatedAt, v.UUIDField
			}
			return
		}
		tmp := exploder(p...)
		return NewExample1ScannerStatic(q.Query(query, tmp[:]...)), p[len(p)-1:]
	case 4:
		const query = `QUERY 4`
		exploder := func(p ...Example1) (r [20]interface{}) {
			for idx, v := range p[:4] {
				r[idx*5+0], r[idx*5+1], r[idx*5+2], r[idx*5+3], r[idx*5+4] = v.CreatedAt, v.ID, v.TextField, v.UpdatedAt, v.UUIDField
			}
			return
		}
		tmp := exploder(p...)
		return NewExample1ScannerStatic(q.Query(query, tmp[:]...)), p[len(p)-1:]
	default:
		const query = `QUERY 5`
		exploder := func(p ...Example1) (r [25]interface{}) {
			for idx, v := range p[:5] {
				r[idx*5+0], r[idx*5+1], r[idx*5+2], r[idx*5+3], r[idx*5+4] = v.CreatedAt, v.ID, v.TextField, v.UpdatedAt, v.UUIDField
			}
			return
		}
		tmp := exploder(p[:5]...)
		return NewExample1ScannerStatic(q.Query(query, tmp[:]...)), p[5:]
	}
}
