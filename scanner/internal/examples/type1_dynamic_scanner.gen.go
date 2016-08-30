package examples

import (
	"database/sql"

	"github.com/lib/pq"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql scanner dynamic --config=scanner-test.config --output=type1_dynamic_scanner.gen.go bitbucket.org/jatone/genieql/scanner/internal/examples.Type1 type1
// invoked by go generate @ examples/type1.go line 11

// DynamicType1DynamicScanner creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func DynamicType1DynamicScanner(rows *sql.Rows, err error) Type1Scanner {
	if err != nil {
		return errType1Scanner{e: err}
	}

	return dynamicType1DynamicScanner{
		Rows: rows,
	}
}

type dynamicType1DynamicScanner struct {
	Rows *sql.Rows
}

func (t dynamicType1DynamicScanner) Scan(arg0 *Type1) error {
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullString
		c1      sql.NullString
		c2      sql.NullBool
		c3      sql.NullBool
		c4      sql.NullInt64
		c5      sql.NullInt64
		c6      pq.NullTime
		c7      pq.NullTime
	)

	if columns, err = t.Rows.Columns(); err != nil {
		return err
	}

	dst = make([]interface{}, 0, len(columns))

	for _, column := range columns {
		switch column {
		case "field1":
			dst = append(dst, &c0)
		case "field2":
			dst = append(dst, &c1)
		case "field3":
			dst = append(dst, &c2)
		case "field4":
			dst = append(dst, &c3)
		case "field5":
			dst = append(dst, &c4)
		case "field6":
			dst = append(dst, &c5)
		case "field7":
			dst = append(dst, &c6)
		case "field8":
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
		case "field1":
			if c0.Valid {
				tmp := c0.String
				arg0.Field1 = tmp
			}
		case "field2":
			if c1.Valid {
				tmp := c1.String
				arg0.Field2 = &tmp
			}
		case "field3":
			if c2.Valid {
				tmp := c2.Bool
				arg0.Field3 = tmp
			}
		case "field4":
			if c3.Valid {
				tmp := c3.Bool
				arg0.Field4 = &tmp
			}
		case "field5":
			if c4.Valid {
				tmp := int(c4.Int64)
				arg0.Field5 = tmp
			}
		case "field6":
			if c5.Valid {
				tmp := int(c5.Int64)
				arg0.Field6 = &tmp
			}
		case "field7":
			if c6.Valid {
				tmp := c6.Time
				arg0.Field7 = tmp
			}
		case "field8":
			if c7.Valid {
				tmp := c7.Time
				arg0.Field8 = &tmp
			}
		}
	}

	return t.Rows.Err()
}

func (t dynamicType1DynamicScanner) Err() error {
	return t.Rows.Err()
}

func (t dynamicType1DynamicScanner) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t dynamicType1DynamicScanner) Next() bool {
	return t.Rows.Next()
}
