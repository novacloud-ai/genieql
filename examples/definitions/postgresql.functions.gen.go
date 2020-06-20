package definitions

import (
	"bitbucket.org/jatone/genieql/internal/sqlx"
	"github.com/jackc/pgtype"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental functions types -o postgresql.functions.gen.go
// invoked by go generate @ definitions/example.go line 7

// customQueryFunction generated by genieql
func customQueryFunction(q sqlx.Queryer, query string, x1, x2, x3 int) ProfileScanner {
	var (
		c0 pgtype.Int8
		c1 pgtype.Int8
		c2 pgtype.Int8
	)

	if err := c0.Set(x1); err != nil {
		return NewProfileScannerDynamic(nil, err)
	}

	if err := c1.Set(x2); err != nil {
		return NewProfileScannerDynamic(nil, err)
	}

	if err := c2.Set(x3); err != nil {
		return NewProfileScannerDynamic(nil, err)
	}

	return NewProfileScannerDynamic(q.Query(query, c0, c1, c2))
}

// customQueryFunction2 generated by genieql
func customQueryFunction2(q sqlx.Queryer, x1, x2, x3 int) ProfileScanner {
	var query = query1
	var (
		c0 pgtype.Int8
		c1 pgtype.Int8
		c2 pgtype.Int8
	)

	if err := c0.Set(x1); err != nil {
		return NewProfileScannerDynamic(nil, err)
	}

	if err := c1.Set(x2); err != nil {
		return NewProfileScannerDynamic(nil, err)
	}

	if err := c2.Set(x3); err != nil {
		return NewProfileScannerDynamic(nil, err)
	}

	return NewProfileScannerDynamic(q.Query(query, c0, c1, c2))
}
