package scanners

import (
	"database/sql"
	"math"
	"time"

	"github.com/jackc/pgtype"
)

// DO NOT EDIT: This File was auto generated by the following command:
// genieql scanner default --config=generators-test.config --output=type1_scanner.gen.go Type1 type1
// invoked by go generate @ scanners/type1.go line 7

// Type1Scanner scanner interface.
type Type1Scanner interface {
	Scan(arg0 *Type1) error
	Next() bool
	Close() error
	Err() error
}

type errType1Scanner struct {
	e error
}

func (t errType1Scanner) Scan(arg0 *Type1) error {
	return t.e
}

func (t errType1Scanner) Next() bool {
	return false
}

func (t errType1Scanner) Err() error {
	return t.e
}

func (t errType1Scanner) Close() error {
	return nil
}

// Type1ScannerStaticColumns generated by genieql
const Type1ScannerStaticColumns = `"field1","field2","field3","field4","field5","field6","field7","field8"`

// NewType1ScannerStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewType1ScannerStatic(rows *sql.Rows, err error) Type1Scanner {
	if err != nil {
		return errType1Scanner{e: err}
	}

	return type1ScannerStatic{
		Rows: rows,
	}
}

// type1ScannerStatic generated by genieql
type type1ScannerStatic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t type1ScannerStatic) Scan(arg0 *Type1) error {
	var (
		c0 pgtype.Text
		c1 pgtype.Text
		c2 pgtype.Bool
		c3 pgtype.Bool
		c4 pgtype.Int4
		c5 pgtype.Int4
		c6 pgtype.Timestamptz
		c7 pgtype.Timestamptz
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7); err != nil {
		return err
	}

	if err := c0.AssignTo(&arg0.Field1); err != nil {
		return err
	}

	if err := c1.AssignTo(&arg0.Field2); err != nil {
		return err
	}

	if err := c2.AssignTo(&arg0.Field3); err != nil {
		return err
	}

	if err := c3.AssignTo(&arg0.Field4); err != nil {
		return err
	}

	if err := c4.AssignTo(&arg0.Field5); err != nil {
		return err
	}

	if err := c5.AssignTo(&arg0.Field6); err != nil {
		return err
	}

	switch c6.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		arg0.Field7 = tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		arg0.Field7 = tmp
	default:
		if err := c6.AssignTo(&arg0.Field7); err != nil {
			return err
		}
	}

	switch c7.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		arg0.Field8 = &tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		arg0.Field8 = &tmp
	default:
		if err := c7.AssignTo(&arg0.Field8); err != nil {
			return err
		}
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t type1ScannerStatic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t type1ScannerStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t type1ScannerStatic) Next() bool {
	return t.Rows.Next()
}

// NewType1ScannerStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewType1ScannerStaticRow(row *sql.Row) Type1ScannerStaticRow {
	return Type1ScannerStaticRow{
		row: row,
	}
}

// Type1ScannerStaticRow generated by genieql
type Type1ScannerStaticRow struct {
	err error
	row *sql.Row
}

// Scan generated by genieql
func (t Type1ScannerStaticRow) Scan(arg0 *Type1) error {
	var (
		c0 pgtype.Text
		c1 pgtype.Text
		c2 pgtype.Bool
		c3 pgtype.Bool
		c4 pgtype.Int4
		c5 pgtype.Int4
		c6 pgtype.Timestamptz
		c7 pgtype.Timestamptz
	)

	if t.err != nil {
		return t.err
	}

	if err := t.row.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7); err != nil {
		return err
	}

	if err := c0.AssignTo(&arg0.Field1); err != nil {
		return err
	}

	if err := c1.AssignTo(&arg0.Field2); err != nil {
		return err
	}

	if err := c2.AssignTo(&arg0.Field3); err != nil {
		return err
	}

	if err := c3.AssignTo(&arg0.Field4); err != nil {
		return err
	}

	if err := c4.AssignTo(&arg0.Field5); err != nil {
		return err
	}

	if err := c5.AssignTo(&arg0.Field6); err != nil {
		return err
	}

	switch c6.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		arg0.Field7 = tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		arg0.Field7 = tmp
	default:
		if err := c6.AssignTo(&arg0.Field7); err != nil {
			return err
		}
	}

	switch c7.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		arg0.Field8 = &tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		arg0.Field8 = &tmp
	default:
		if err := c7.AssignTo(&arg0.Field8); err != nil {
			return err
		}
	}

	return nil
}

// Err set an error to return by scan
func (t Type1ScannerStaticRow) Err(err error) Type1ScannerStaticRow {
	t.err = err
	return t
}
