package alternate1

import (
	"database/sql"

	"github.com/jackc/pgtype"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental scanners types --config=generators-test.config -o postgresql.scanners.gen.go
// invoked by go generate @ alternate1/10_genieql.go line 4

// Type1Scanner scanner interface.
type Type1Scanner interface {
	Scan(sp0 *Type1) error
	Next() bool
	Close() error
	Err() error
}

type errType1Scanner struct {
	e error
}

func (t errType1Scanner) Scan(sp0 *Type1) error {
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
const Type1ScannerStaticColumns = `"field1","field2","field3","field4","field5","field6","field7","field8","unmappedfield"`

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
func (t type1ScannerStatic) Scan(sp0 *Type1) error {
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

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8); err != nil {
		return err
	}

	if err := c0.AssignTo(&sp0.Field1); err != nil {
		return err
	}

	if err := c1.AssignTo(&sp0.Field2); err != nil {
		return err
	}

	if err := c2.AssignTo(&sp0.Field3); err != nil {
		return err
	}

	if err := c3.AssignTo(&sp0.Field4); err != nil {
		return err
	}

	if err := c4.AssignTo(&sp0.Field5); err != nil {
		return err
	}

	if err := c5.AssignTo(&sp0.Field6); err != nil {
		return err
	}

	if err := c6.AssignTo(&sp0.Field7); err != nil {
		return err
	}

	if err := c7.AssignTo(&sp0.Field8); err != nil {
		return err
	}

	if err := c8.AssignTo(&sp0.Unmappedfield); err != nil {
		return err
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
	row *sql.Row
}

// Scan generated by genieql
func (t Type1ScannerStaticRow) Scan(sp0 *Type1) error {
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

	if err := t.row.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8); err != nil {
		return err
	}

	if err := c0.AssignTo(&sp0.Field1); err != nil {
		return err
	}

	if err := c1.AssignTo(&sp0.Field2); err != nil {
		return err
	}

	if err := c2.AssignTo(&sp0.Field3); err != nil {
		return err
	}

	if err := c3.AssignTo(&sp0.Field4); err != nil {
		return err
	}

	if err := c4.AssignTo(&sp0.Field5); err != nil {
		return err
	}

	if err := c5.AssignTo(&sp0.Field6); err != nil {
		return err
	}

	if err := c6.AssignTo(&sp0.Field7); err != nil {
		return err
	}

	if err := c7.AssignTo(&sp0.Field8); err != nil {
		return err
	}

	if err := c8.AssignTo(&sp0.Unmappedfield); err != nil {
		return err
	}

	return nil
}

// NewType1ScannerDynamic creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func NewType1ScannerDynamic(rows *sql.Rows, err error) Type1Scanner {
	if err != nil {
		return errType1Scanner{e: err}
	}

	return type1ScannerDynamic{
		Rows: rows,
	}
}

// type1ScannerDynamic generated by genieql
type type1ScannerDynamic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t type1ScannerDynamic) Scan(sp0 *Type1) error {
	const (
		cn0 = "field1"
		cn1 = "field2"
		cn2 = "field3"
		cn3 = "field4"
		cn4 = "field5"
		cn5 = "field6"
		cn6 = "field7"
		cn7 = "field8"
		cn8 = "unmappedfield"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      pgtype.Text
		c1      pgtype.Text
		c2      pgtype.Bool
		c3      pgtype.Bool
		c4      pgtype.Int4
		c5      pgtype.Int4
		c6      pgtype.Timestamptz
		c7      pgtype.Timestamptz
		c8      pgtype.Int4
	)

	if columns, err = t.Rows.Columns(); err != nil {
		return err
	}

	dst = make([]interface{}, 0, len(columns))

	for _, column := range columns {
		switch column {
		case cn0:
			dst = append(dst, &c0)
		case cn1:
			dst = append(dst, &c1)
		case cn2:
			dst = append(dst, &c2)
		case cn3:
			dst = append(dst, &c3)
		case cn4:
			dst = append(dst, &c4)
		case cn5:
			dst = append(dst, &c5)
		case cn6:
			dst = append(dst, &c6)
		case cn7:
			dst = append(dst, &c7)
		case cn8:
			dst = append(dst, &c8)
		default:
			dst = append(dst, &ignored)
		}
	}

	if err := t.Rows.Scan(dst...); err != nil {
		return err
	}

	for _, column := range columns {
		switch column {
		case cn0:
			if err := c0.AssignTo(&sp0.Field1); err != nil {
				return err
			}
		case cn1:
			if err := c1.AssignTo(&sp0.Field2); err != nil {
				return err
			}
		case cn2:
			if err := c2.AssignTo(&sp0.Field3); err != nil {
				return err
			}
		case cn3:
			if err := c3.AssignTo(&sp0.Field4); err != nil {
				return err
			}
		case cn4:
			if err := c4.AssignTo(&sp0.Field5); err != nil {
				return err
			}
		case cn5:
			if err := c5.AssignTo(&sp0.Field6); err != nil {
				return err
			}
		case cn6:
			if err := c6.AssignTo(&sp0.Field7); err != nil {
				return err
			}
		case cn7:
			if err := c7.AssignTo(&sp0.Field8); err != nil {
				return err
			}
		case cn8:
			if err := c8.AssignTo(&sp0.Unmappedfield); err != nil {
				return err
			}
		}
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t type1ScannerDynamic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t type1ScannerDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t type1ScannerDynamic) Next() bool {
	return t.Rows.Next()
}
