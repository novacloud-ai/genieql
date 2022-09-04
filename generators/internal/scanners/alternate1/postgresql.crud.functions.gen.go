package alternate1

import (
	"math"
	"time"

	"bitbucket.org/jatone/genieql/internal/sqlx"
	"github.com/jackc/pgtype"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental crud --config=generators-test.config -o postgresql.crud.functions.gen.go --table=type1 --scanner=NewType1ScannerStatic --unique-scanner=NewType1ScannerStaticRow Type1
// invoked by go generate @ alternate1/10_genieql.go line 5

// Type1Insert generated by genieql
func Type1Insert(q sqlx.Queryer, arg1 Type1) Type1ScannerStaticRow {
	const query = `INSERT INTO type1 ("field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield"`
	var (
		c0 pgtype.Text
		c1 pgtype.Text
		c2 pgtype.Bool
		c3 pgtype.Bool
		c4 pgtype.Int4
		c5 pgtype.Int4
		c6 pgtype.Timestamptz
		c7 pgtype.Timestamptz
		c8 pgtype.Int4
	)

	if err := c0.Set(arg1.Field1); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c1.Set(arg1.Field2); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c2.Set(arg1.Field3); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c3.Set(arg1.Field4); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c4.Set(arg1.Field5); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c5.Set(arg1.Field6); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	switch arg1.Field7 {
	case time.Unix(math.MaxInt64-62135596800, 999999999):
		if err := c6.Set(pgtype.Infinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	case time.Unix(math.MinInt64, math.MinInt64):
		if err := c6.Set(pgtype.NegativeInfinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	default:
		if err := c6.Set(arg1.Field7); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	}

	switch *arg1.Field8 {
	case time.Unix(math.MaxInt64-62135596800, 999999999):
		if err := c7.Set(pgtype.Infinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	case time.Unix(math.MinInt64, math.MinInt64):
		if err := c7.Set(pgtype.NegativeInfinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	default:
		if err := c7.Set(arg1.Field8); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	}

	if err := c8.Set(arg1.Unmappedfield); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0, c1, c2, c3, c4, c5, c6, c7, c8))
}

// Type1FindByField1 generated by genieql
func Type1FindByField1(q sqlx.Queryer, field1 string) Type1ScannerStaticRow {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field1" = $1`
	var (
		c0 pgtype.Text
	)

	if err := c0.Set(field1); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}

// Type1LookupByField1 generated by genieql
func Type1LookupByField1(q sqlx.Queryer, field1 string) Type1Scanner {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field1" = $1`
	var (
		c0 pgtype.Text
	)

	if err := c0.Set(field1); err != nil {
		return NewType1ScannerStatic(nil, err)
	}

	return NewType1ScannerStatic(q.Query(query, c0))
}

// Type1FindByField2 generated by genieql
func Type1FindByField2(q sqlx.Queryer, field2 string) Type1ScannerStaticRow {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field2" = $1`
	var (
		c0 pgtype.Text
	)

	if err := c0.Set(field2); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}

// Type1LookupByField2 generated by genieql
func Type1LookupByField2(q sqlx.Queryer, field2 string) Type1Scanner {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field2" = $1`
	var (
		c0 pgtype.Text
	)

	if err := c0.Set(field2); err != nil {
		return NewType1ScannerStatic(nil, err)
	}

	return NewType1ScannerStatic(q.Query(query, c0))
}

// Type1FindByField3 generated by genieql
func Type1FindByField3(q sqlx.Queryer, field3 bool) Type1ScannerStaticRow {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field3" = $1`
	var (
		c0 pgtype.Bool
	)

	if err := c0.Set(field3); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}

// Type1LookupByField3 generated by genieql
func Type1LookupByField3(q sqlx.Queryer, field3 bool) Type1Scanner {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field3" = $1`
	var (
		c0 pgtype.Bool
	)

	if err := c0.Set(field3); err != nil {
		return NewType1ScannerStatic(nil, err)
	}

	return NewType1ScannerStatic(q.Query(query, c0))
}

// Type1FindByField4 generated by genieql
func Type1FindByField4(q sqlx.Queryer, field4 bool) Type1ScannerStaticRow {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field4" = $1`
	var (
		c0 pgtype.Bool
	)

	if err := c0.Set(field4); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}

// Type1LookupByField4 generated by genieql
func Type1LookupByField4(q sqlx.Queryer, field4 bool) Type1Scanner {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field4" = $1`
	var (
		c0 pgtype.Bool
	)

	if err := c0.Set(field4); err != nil {
		return NewType1ScannerStatic(nil, err)
	}

	return NewType1ScannerStatic(q.Query(query, c0))
}

// Type1FindByField5 generated by genieql
func Type1FindByField5(q sqlx.Queryer, field5 int) Type1ScannerStaticRow {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field5" = $1`
	var (
		c0 pgtype.Int8
	)

	if err := c0.Set(field5); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}

// Type1LookupByField5 generated by genieql
func Type1LookupByField5(q sqlx.Queryer, field5 int) Type1Scanner {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field5" = $1`
	var (
		c0 pgtype.Int8
	)

	if err := c0.Set(field5); err != nil {
		return NewType1ScannerStatic(nil, err)
	}

	return NewType1ScannerStatic(q.Query(query, c0))
}

// Type1FindByField6 generated by genieql
func Type1FindByField6(q sqlx.Queryer, field6 int) Type1ScannerStaticRow {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field6" = $1`
	var (
		c0 pgtype.Int8
	)

	if err := c0.Set(field6); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}

// Type1LookupByField6 generated by genieql
func Type1LookupByField6(q sqlx.Queryer, field6 int) Type1Scanner {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field6" = $1`
	var (
		c0 pgtype.Int8
	)

	if err := c0.Set(field6); err != nil {
		return NewType1ScannerStatic(nil, err)
	}

	return NewType1ScannerStatic(q.Query(query, c0))
}

// Type1FindByField7 generated by genieql
func Type1FindByField7(q sqlx.Queryer, field7 time.Time) Type1ScannerStaticRow {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field7" = $1`
	var (
		c0 pgtype.Timestamptz
	)

	switch field7 {
	case time.Unix(math.MaxInt64-62135596800, 999999999):
		if err := c0.Set(pgtype.Infinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	case time.Unix(math.MinInt64, math.MinInt64):
		if err := c0.Set(pgtype.NegativeInfinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	default:
		if err := c0.Set(field7); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}

// Type1LookupByField7 generated by genieql
func Type1LookupByField7(q sqlx.Queryer, field7 time.Time) Type1Scanner {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field7" = $1`
	var (
		c0 pgtype.Timestamptz
	)

	switch field7 {
	case time.Unix(math.MaxInt64-62135596800, 999999999):
		if err := c0.Set(pgtype.Infinity); err != nil {
			return NewType1ScannerStatic(nil, err)
		}
	case time.Unix(math.MinInt64, math.MinInt64):
		if err := c0.Set(pgtype.NegativeInfinity); err != nil {
			return NewType1ScannerStatic(nil, err)
		}
	default:
		if err := c0.Set(field7); err != nil {
			return NewType1ScannerStatic(nil, err)
		}
	}

	return NewType1ScannerStatic(q.Query(query, c0))
}

// Type1FindByField8 generated by genieql
func Type1FindByField8(q sqlx.Queryer, field8 time.Time) Type1ScannerStaticRow {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field8" = $1`
	var (
		c0 pgtype.Timestamptz
	)

	switch field8 {
	case time.Unix(math.MaxInt64-62135596800, 999999999):
		if err := c0.Set(pgtype.Infinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	case time.Unix(math.MinInt64, math.MinInt64):
		if err := c0.Set(pgtype.NegativeInfinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	default:
		if err := c0.Set(field8); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}

// Type1LookupByField8 generated by genieql
func Type1LookupByField8(q sqlx.Queryer, field8 time.Time) Type1Scanner {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field8" = $1`
	var (
		c0 pgtype.Timestamptz
	)

	switch field8 {
	case time.Unix(math.MaxInt64-62135596800, 999999999):
		if err := c0.Set(pgtype.Infinity); err != nil {
			return NewType1ScannerStatic(nil, err)
		}
	case time.Unix(math.MinInt64, math.MinInt64):
		if err := c0.Set(pgtype.NegativeInfinity); err != nil {
			return NewType1ScannerStatic(nil, err)
		}
	default:
		if err := c0.Set(field8); err != nil {
			return NewType1ScannerStatic(nil, err)
		}
	}

	return NewType1ScannerStatic(q.Query(query, c0))
}

// Type1FindByUnmappedfield generated by genieql
func Type1FindByUnmappedfield(q sqlx.Queryer, unmappedfield int) Type1ScannerStaticRow {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "unmappedfield" = $1`
	var (
		c0 pgtype.Int8
	)

	if err := c0.Set(unmappedfield); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}

// Type1LookupByUnmappedfield generated by genieql
func Type1LookupByUnmappedfield(q sqlx.Queryer, unmappedfield int) Type1Scanner {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "unmappedfield" = $1`
	var (
		c0 pgtype.Int8
	)

	if err := c0.Set(unmappedfield); err != nil {
		return NewType1ScannerStatic(nil, err)
	}

	return NewType1ScannerStatic(q.Query(query, c0))
}

// Type1FindByKey generated by genieql
func Type1FindByKey(q sqlx.Queryer, field1 string) Type1ScannerStaticRow {
	const query = `SELECT "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield" FROM type1 WHERE "field1" = $1`
	var (
		c0 pgtype.Text
	)

	if err := c0.Set(field1); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}

// Type1UpdateByID generated by genieql
func Type1UpdateByID(q sqlx.Queryer, field1 string, update Type1) Type1ScannerStaticRow {
	const query = `UPDATE type1 SET "field2" = $1, "field3" = $2, "field4" = $3, "field5" = $4, "field6" = $5, "field7" = $6, "field8" = $7, "unmappedfield" = $8 WHERE "field1" = $9 RETURNING "field2","field3","field4","field5","field6","field7","field8","unmappedfield"`
	var (
		c0 pgtype.Text
		c1 pgtype.Text
		c2 pgtype.Text
		c3 pgtype.Bool
		c4 pgtype.Bool
		c5 pgtype.Int4
		c6 pgtype.Int4
		c7 pgtype.Timestamptz
		c8 pgtype.Timestamptz
		c9 pgtype.Int4
	)

	if err := c0.Set(field1); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c1.Set(update.Field1); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c2.Set(update.Field2); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c3.Set(update.Field3); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c4.Set(update.Field4); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c5.Set(update.Field5); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	if err := c6.Set(update.Field6); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	switch update.Field7 {
	case time.Unix(math.MaxInt64-62135596800, 999999999):
		if err := c7.Set(pgtype.Infinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	case time.Unix(math.MinInt64, math.MinInt64):
		if err := c7.Set(pgtype.NegativeInfinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	default:
		if err := c7.Set(update.Field7); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	}

	switch *update.Field8 {
	case time.Unix(math.MaxInt64-62135596800, 999999999):
		if err := c8.Set(pgtype.Infinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	case time.Unix(math.MinInt64, math.MinInt64):
		if err := c8.Set(pgtype.NegativeInfinity); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	default:
		if err := c8.Set(update.Field8); err != nil {
			return NewType1ScannerStaticRow(nil).Err(err)
		}
	}

	if err := c9.Set(update.Unmappedfield); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0, c1, c2, c3, c4, c5, c6, c7, c8, c9))
}

// Type1DeleteByID generated by genieql
func Type1DeleteByID(q sqlx.Queryer, field1 string) Type1ScannerStaticRow {
	const query = `DELETE FROM type1 WHERE "field1" = $1 RETURNING "field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield"`
	var (
		c0 pgtype.Text
	)

	if err := c0.Set(field1); err != nil {
		return NewType1ScannerStaticRow(nil).Err(err)
	}

	return NewType1ScannerStaticRow(q.QueryRow(query, c0))
}
