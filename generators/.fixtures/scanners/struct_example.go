package example

import "database/sql"

// StructExample scanner interface.
type StructExample interface {
	Scan(arg *StructA) error
	Next() bool
	Close() error
	Err() error
}

type errStructExample struct {
	e error
}

func (t errStructExample) Scan(arg *StructA) error {
	return t.e
}

func (t errStructExample) Next() bool {
	return false
}

func (t errStructExample) Err() error {
	return t.e
}

func (t errStructExample) Close() error {
	return nil
}

// StructExampleStaticColumns generated by genieql
const StructExampleStaticColumns = `a,b,c,d,e,f,g,h`

// NewStructExampleStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewStructExampleStatic(rows *sql.Rows, err error) StructExample {
	if err != nil {
		return errStructExample{e: err}
	}

	return structExampleStatic{
		Rows: rows,
	}
}

// structExampleStatic generated by genieql
type structExampleStatic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t structExampleStatic) Scan(arg *StructA) error {
	var (
		c0 sql.NullInt64
		c1 sql.NullInt64
		c2 sql.NullInt64
		c3 sql.NullBool
		c4 sql.NullBool
		c5 sql.NullBool
		c6 sql.NullInt64
		c7 sql.NullBool
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7); err != nil {
		return err
	}

	if c0.Valid {
		tmp := int(c0.Int64)
		arg.A = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		arg.B = tmp
	}

	if c2.Valid {
		tmp := int(c2.Int64)
		arg.C = tmp
	}

	if c3.Valid {
		tmp := c3.Bool
		arg.D = tmp
	}

	if c4.Valid {
		tmp := c4.Bool
		arg.E = tmp
	}

	if c5.Valid {
		tmp := c5.Bool
		arg.F = tmp
	}

	if c6.Valid {
		tmp := int(c6.Int64)
		*arg.G = tmp
	}

	if c7.Valid {
		tmp := c7.Bool
		*arg.H = tmp
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t structExampleStatic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t structExampleStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t structExampleStatic) Next() bool {
	return t.Rows.Next()
}

// NewStructExampleStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewStructExampleStaticRow(row *sql.Row, err error) StructExampleStaticRow {
	return StructExampleStaticRow{
		err: err,
		row: row,
	}
}

// StructExampleStaticRow generated by genieql
type StructExampleStaticRow struct {
	err error
	row *sql.Row
}

// Scan generated by genieql
func (t StructExampleStaticRow) Scan(arg *StructA) error {
	var (
		c0 sql.NullInt64
		c1 sql.NullInt64
		c2 sql.NullInt64
		c3 sql.NullBool
		c4 sql.NullBool
		c5 sql.NullBool
		c6 sql.NullInt64
		c7 sql.NullBool
	)

	if t.err != nil {
		return err
	}

	if err := t.row.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7); err != nil {
		return err
	}

	if c0.Valid {
		tmp := int(c0.Int64)
		arg.A = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		arg.B = tmp
	}

	if c2.Valid {
		tmp := int(c2.Int64)
		arg.C = tmp
	}

	if c3.Valid {
		tmp := c3.Bool
		arg.D = tmp
	}

	if c4.Valid {
		tmp := c4.Bool
		arg.E = tmp
	}

	if c5.Valid {
		tmp := c5.Bool
		arg.F = tmp
	}

	if c6.Valid {
		tmp := int(c6.Int64)
		*arg.G = tmp
	}

	if c7.Valid {
		tmp := c7.Bool
		*arg.H = tmp
	}

	return nil
}

// NewStructExampleDynamic creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func NewStructExampleDynamic(rows *sql.Rows, err error) StructExample {
	if err != nil {
		return errStructExample{e: err}
	}

	return structExampleDynamic{
		Rows: rows,
	}
}

// structExampleDynamic generated by genieql
type structExampleDynamic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t structExampleDynamic) Scan(arg *StructA) error {
	const (
		cn0 = "a"
		cn1 = "b"
		cn2 = "c"
		cn3 = "d"
		cn4 = "e"
		cn5 = "f"
		cn6 = "g"
		cn7 = "h"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullInt64
		c1      sql.NullInt64
		c2      sql.NullInt64
		c3      sql.NullBool
		c4      sql.NullBool
		c5      sql.NullBool
		c6      sql.NullInt64
		c7      sql.NullBool
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
			if c0.Valid {
				tmp := int(c0.Int64)
				arg.A = tmp
			}
		case cn1:
			if c1.Valid {
				tmp := int(c1.Int64)
				arg.B = tmp
			}
		case cn2:
			if c2.Valid {
				tmp := int(c2.Int64)
				arg.C = tmp
			}
		case cn3:
			if c3.Valid {
				tmp := c3.Bool
				arg.D = tmp
			}
		case cn4:
			if c4.Valid {
				tmp := c4.Bool
				arg.E = tmp
			}
		case cn5:
			if c5.Valid {
				tmp := c5.Bool
				arg.F = tmp
			}
		case cn6:
			if c6.Valid {
				tmp := int(c6.Int64)
				*arg.G = tmp
			}
		case cn7:
			if c7.Valid {
				tmp := c7.Bool
				*arg.H = tmp
			}
		}
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t structExampleDynamic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t structExampleDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t structExampleDynamic) Next() bool {
	return t.Rows.Next()
}
