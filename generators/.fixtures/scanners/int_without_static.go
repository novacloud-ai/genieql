package example

import "database/sql"

// IntNoStatic scanner interface.
type IntNoStatic interface {
	Scan(arg *int) error
	Next() bool
	Close() error
	Err() error
}

type errIntNoStatic struct {
	e error
}

func (t errIntNoStatic) Scan(arg *int) error {
	return t.e
}

func (t errIntNoStatic) Next() bool {
	return false
}

func (t errIntNoStatic) Err() error {
	return t.e
}

func (t errIntNoStatic) Close() error {
	return nil
}

// NewIntNoStaticDynamic creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func NewIntNoStaticDynamic(rows *sql.Rows, err error) IntNoStatic {
	if err != nil {
		return errIntNoStatic{e: err}
	}

	return intNoStaticDynamic{
		Rows: rows,
	}
}

type intNoStaticDynamic struct {
	Rows *sql.Rows
}

func (t intNoStaticDynamic) Scan(arg *int) error {
	const (
		arg0 = "arg"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullInt64
	)

	if columns, err = t.Rows.Columns(); err != nil {
		return err
	}

	dst = make([]interface{}, 0, len(columns))

	for _, column := range columns {
		switch column {
		case arg0:
			dst = append(dst, &c0)
		default:
			dst = append(dst, &ignored)
		}
	}

	if err := t.Rows.Scan(dst...); err != nil {
		return err
	}

	for _, column := range columns {
		switch column {
		case arg0:
			if c0.Valid {
				tmp := int(c0.Int64)
				*arg = tmp
			}
		}
	}

	return t.Rows.Err()
}

func (t intNoStaticDynamic) Err() error {
	return t.Rows.Err()
}

func (t intNoStaticDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t intNoStaticDynamic) Next() bool {
	return t.Rows.Next()
}
