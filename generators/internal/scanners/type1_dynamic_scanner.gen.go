package scanners

import (
	"database/sql"
	"math"
	"time"

	"github.com/jackc/pgtype"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql scanner dynamic --config=generators-test.config --output=type1_dynamic_scanner.gen.go Type1 type1
// invoked by go generate @ scanners/type1.go line 8

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
func (t type1ScannerDynamic) Scan(arg0 *Type1) error {
	const (
		cn0 = "field1"
		cn1 = "field2"
		cn2 = "field3"
		cn3 = "field4"
		cn4 = "field5"
		cn5 = "field6"
		cn6 = "field7"
		cn7 = "field8"
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
			if err := c0.AssignTo(&arg0.Field1); err != nil {
				return err
			}
		case cn1:
			if err := c1.AssignTo(&arg0.Field2); err != nil {
				return err
			}
		case cn2:
			if err := c2.AssignTo(&arg0.Field3); err != nil {
				return err
			}
		case cn3:
			if err := c3.AssignTo(&arg0.Field4); err != nil {
				return err
			}
		case cn4:
			if err := c4.AssignTo(&arg0.Field5); err != nil {
				return err
			}
		case cn5:
			if err := c5.AssignTo(&arg0.Field6); err != nil {
				return err
			}
		case cn6:
			switch c6.InfinityModifier {
			case pgtype.Infinity:
				tmp := time.Unix(math.MaxInt64, math.MaxInt64)
				arg0.Field7 = tmp
			case pgtype.NegativeInfinity:
				tmp := time.Unix(math.MinInt64, math.MinInt64)
				arg0.Field7 = tmp
			default:
				if err := c6.AssignTo(&arg0.Field7); err != nil {
					return err
				}
			}
		case cn7:
			switch c7.InfinityModifier {
			case pgtype.Infinity:
				tmp := time.Unix(math.MaxInt64, math.MaxInt64)
				arg0.Field8 = &tmp
			case pgtype.NegativeInfinity:
				tmp := time.Unix(math.MinInt64, math.MinInt64)
				arg0.Field8 = &tmp
			default:
				if err := c7.AssignTo(&arg0.Field8); err != nil {
					return err
				}
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
