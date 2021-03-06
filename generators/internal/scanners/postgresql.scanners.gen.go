package scanners

import (
	"database/sql"
	"math"
	"time"

	"bitbucket.org/jatone/genieql/generators/internal/scanners/alternate1"
	"bitbucket.org/jatone/genieql/generators/internal/scanners/alternate2"
	"github.com/jackc/pgtype"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental scanners types --config=generators-test.config -o postgresql.scanners.gen.go
// invoked by go generate @ scanners/type1.go line 6

// ComboScanner scanner interface.
type ComboScanner interface {
	Scan(t1 *alternate1.Type1, t2 *alternate2.Type1, t3 *Type1) error
	Next() bool
	Close() error
	Err() error
}

type errComboScanner struct {
	e error
}

func (t errComboScanner) Scan(t1 *alternate1.Type1, t2 *alternate2.Type1, t3 *Type1) error {
	return t.e
}

func (t errComboScanner) Next() bool {
	return false
}

func (t errComboScanner) Err() error {
	return t.e
}

func (t errComboScanner) Close() error {
	return nil
}

// NewComboScannerStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewComboScannerStatic(rows *sql.Rows, err error) ComboScanner {
	if err != nil {
		return errComboScanner{e: err}
	}

	return comboScannerStatic{
		Rows: rows,
	}
}

// comboScannerStatic generated by genieql
type comboScannerStatic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t comboScannerStatic) Scan(t1 *alternate1.Type1, t2 *alternate2.Type1, t3 *Type1) error {
	var (
		c0  pgtype.Text
		c1  pgtype.Text
		c2  pgtype.Bool
		c3  pgtype.Bool
		c4  pgtype.Int4
		c5  pgtype.Int4
		c6  pgtype.Timestamptz
		c7  pgtype.Timestamptz
		c8  pgtype.Int4
		c9  pgtype.Text
		c10 pgtype.Text
		c11 pgtype.Bool
		c12 pgtype.Bool
		c13 pgtype.Int4
		c14 pgtype.Int4
		c15 pgtype.Timestamptz
		c16 pgtype.Timestamptz
		c17 pgtype.Int4
		c18 pgtype.Text
		c19 pgtype.Text
		c20 pgtype.Bool
		c21 pgtype.Bool
		c22 pgtype.Int8
		c23 pgtype.Int8
		c24 pgtype.Timestamptz
		c25 pgtype.Timestamptz
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15, &c16, &c17, &c18, &c19, &c20, &c21, &c22, &c23, &c24, &c25); err != nil {
		return err
	}

	if err := c0.AssignTo(&t1.Field1); err != nil {
		return err
	}

	if err := c1.AssignTo(&t1.Field2); err != nil {
		return err
	}

	if err := c2.AssignTo(&t1.Field3); err != nil {
		return err
	}

	if err := c3.AssignTo(&t1.Field4); err != nil {
		return err
	}

	if err := c4.AssignTo(&t1.Field5); err != nil {
		return err
	}

	if err := c5.AssignTo(&t1.Field6); err != nil {
		return err
	}

	switch c6.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t1.Field7 = tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t1.Field7 = tmp
	default:
		if err := c6.AssignTo(&t1.Field7); err != nil {
			return err
		}
	}

	switch c7.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t1.Field8 = &tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t1.Field8 = &tmp
	default:
		if err := c7.AssignTo(&t1.Field8); err != nil {
			return err
		}
	}

	if err := c8.AssignTo(&t1.Unmappedfield); err != nil {
		return err
	}

	if err := c9.AssignTo(&t2.Field1); err != nil {
		return err
	}

	if err := c10.AssignTo(&t2.Field2); err != nil {
		return err
	}

	if err := c11.AssignTo(&t2.Field3); err != nil {
		return err
	}

	if err := c12.AssignTo(&t2.Field4); err != nil {
		return err
	}

	if err := c13.AssignTo(&t2.Field5); err != nil {
		return err
	}

	if err := c14.AssignTo(&t2.Field6); err != nil {
		return err
	}

	switch c15.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t2.Field7 = tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t2.Field7 = tmp
	default:
		if err := c15.AssignTo(&t2.Field7); err != nil {
			return err
		}
	}

	switch c16.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t2.Field8 = &tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t2.Field8 = &tmp
	default:
		if err := c16.AssignTo(&t2.Field8); err != nil {
			return err
		}
	}

	if err := c17.AssignTo(&t2.Unmappedfield); err != nil {
		return err
	}

	if err := c18.AssignTo(&t3.Field1); err != nil {
		return err
	}

	if err := c19.AssignTo(&t3.Field2); err != nil {
		return err
	}

	if err := c20.AssignTo(&t3.Field3); err != nil {
		return err
	}

	if err := c21.AssignTo(&t3.Field4); err != nil {
		return err
	}

	if err := c22.AssignTo(&t3.Field5); err != nil {
		return err
	}

	if err := c23.AssignTo(&t3.Field6); err != nil {
		return err
	}

	switch c24.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t3.Field7 = tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t3.Field7 = tmp
	default:
		if err := c24.AssignTo(&t3.Field7); err != nil {
			return err
		}
	}

	switch c25.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t3.Field8 = &tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t3.Field8 = &tmp
	default:
		if err := c25.AssignTo(&t3.Field8); err != nil {
			return err
		}
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t comboScannerStatic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t comboScannerStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t comboScannerStatic) Next() bool {
	return t.Rows.Next()
}

// NewComboScannerStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewComboScannerStaticRow(row *sql.Row) ComboScannerStaticRow {
	return ComboScannerStaticRow{
		row: row,
	}
}

// ComboScannerStaticRow generated by genieql
type ComboScannerStaticRow struct {
	err error
	row *sql.Row
}

// Scan generated by genieql
func (t ComboScannerStaticRow) Scan(t1 *alternate1.Type1, t2 *alternate2.Type1, t3 *Type1) error {
	var (
		c0  pgtype.Text
		c1  pgtype.Text
		c2  pgtype.Bool
		c3  pgtype.Bool
		c4  pgtype.Int4
		c5  pgtype.Int4
		c6  pgtype.Timestamptz
		c7  pgtype.Timestamptz
		c8  pgtype.Int4
		c9  pgtype.Text
		c10 pgtype.Text
		c11 pgtype.Bool
		c12 pgtype.Bool
		c13 pgtype.Int4
		c14 pgtype.Int4
		c15 pgtype.Timestamptz
		c16 pgtype.Timestamptz
		c17 pgtype.Int4
		c18 pgtype.Text
		c19 pgtype.Text
		c20 pgtype.Bool
		c21 pgtype.Bool
		c22 pgtype.Int8
		c23 pgtype.Int8
		c24 pgtype.Timestamptz
		c25 pgtype.Timestamptz
	)

	if t.err != nil {
		return t.err
	}

	if err := t.row.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &c11, &c12, &c13, &c14, &c15, &c16, &c17, &c18, &c19, &c20, &c21, &c22, &c23, &c24, &c25); err != nil {
		return err
	}

	if err := c0.AssignTo(&t1.Field1); err != nil {
		return err
	}

	if err := c1.AssignTo(&t1.Field2); err != nil {
		return err
	}

	if err := c2.AssignTo(&t1.Field3); err != nil {
		return err
	}

	if err := c3.AssignTo(&t1.Field4); err != nil {
		return err
	}

	if err := c4.AssignTo(&t1.Field5); err != nil {
		return err
	}

	if err := c5.AssignTo(&t1.Field6); err != nil {
		return err
	}

	switch c6.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t1.Field7 = tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t1.Field7 = tmp
	default:
		if err := c6.AssignTo(&t1.Field7); err != nil {
			return err
		}
	}

	switch c7.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t1.Field8 = &tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t1.Field8 = &tmp
	default:
		if err := c7.AssignTo(&t1.Field8); err != nil {
			return err
		}
	}

	if err := c8.AssignTo(&t1.Unmappedfield); err != nil {
		return err
	}

	if err := c9.AssignTo(&t2.Field1); err != nil {
		return err
	}

	if err := c10.AssignTo(&t2.Field2); err != nil {
		return err
	}

	if err := c11.AssignTo(&t2.Field3); err != nil {
		return err
	}

	if err := c12.AssignTo(&t2.Field4); err != nil {
		return err
	}

	if err := c13.AssignTo(&t2.Field5); err != nil {
		return err
	}

	if err := c14.AssignTo(&t2.Field6); err != nil {
		return err
	}

	switch c15.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t2.Field7 = tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t2.Field7 = tmp
	default:
		if err := c15.AssignTo(&t2.Field7); err != nil {
			return err
		}
	}

	switch c16.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t2.Field8 = &tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t2.Field8 = &tmp
	default:
		if err := c16.AssignTo(&t2.Field8); err != nil {
			return err
		}
	}

	if err := c17.AssignTo(&t2.Unmappedfield); err != nil {
		return err
	}

	if err := c18.AssignTo(&t3.Field1); err != nil {
		return err
	}

	if err := c19.AssignTo(&t3.Field2); err != nil {
		return err
	}

	if err := c20.AssignTo(&t3.Field3); err != nil {
		return err
	}

	if err := c21.AssignTo(&t3.Field4); err != nil {
		return err
	}

	if err := c22.AssignTo(&t3.Field5); err != nil {
		return err
	}

	if err := c23.AssignTo(&t3.Field6); err != nil {
		return err
	}

	switch c24.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t3.Field7 = tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t3.Field7 = tmp
	default:
		if err := c24.AssignTo(&t3.Field7); err != nil {
			return err
		}
	}

	switch c25.InfinityModifier {
	case pgtype.Infinity:
		tmp := time.Unix(math.MaxInt64-62135596800, 999999999)
		t3.Field8 = &tmp
	case pgtype.NegativeInfinity:
		tmp := time.Unix(math.MinInt64, math.MinInt64)
		t3.Field8 = &tmp
	default:
		if err := c25.AssignTo(&t3.Field8); err != nil {
			return err
		}
	}

	return nil
}

// Err set an error to return by scan
func (t ComboScannerStaticRow) Err(err error) ComboScannerStaticRow {
	t.err = err
	return t
}
