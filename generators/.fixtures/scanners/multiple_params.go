package example

import "database/sql"

// MultipleParam scanner interface.
type MultipleParam interface {
	Scan(arg1, arg2 *int, arg3 *bool, arg4 *string) error
	Next() bool
	Close() error
	Err() error
}

type errMultipleParam struct {
	e error
}

func (t errMultipleParam) Scan(arg1, arg2 *int, arg3 *bool, arg4 *string) error {
	return t.e
}

func (t errMultipleParam) Next() bool {
	return false
}

func (t errMultipleParam) Err() error {
	return t.e
}

func (t errMultipleParam) Close() error {
	return nil
}

const MultipleParamStaticColumns = "arg1,arg2,arg3,arg4"

// NewMultipleParamStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewMultipleParamStatic(rows *sql.Rows, err error) MultipleParam {
	if err != nil {
		return errMultipleParam{e: err}
	}

	return multipleParamStatic{
		Rows: rows,
	}
}

type multipleParamStatic struct {
	Rows *sql.Rows
}

func (t multipleParamStatic) Scan(arg1, arg2 *int, arg3 *bool, arg4 *string) error {
	var (
		c0 sql.NullInt64
		c1 sql.NullInt64
		c2 sql.NullBool
		c3 sql.NullString
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3); err != nil {
		return err
	}

	if c0.Valid {
		tmp := int(c0.Int64)
		*arg1 = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		*arg2 = tmp
	}

	if c2.Valid {
		tmp := c2.Bool
		*arg3 = tmp
	}

	if c3.Valid {
		tmp := c3.String
		*arg4 = tmp
	}

	return t.Rows.Err()
}

func (t multipleParamStatic) Err() error {
	return t.Rows.Err()
}

func (t multipleParamStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t multipleParamStatic) Next() bool {
	return t.Rows.Next()
}

// NewMultipleParamStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewMultipleParamStaticRow(row *sql.Row) MultipleParamStaticRow {
	return MultipleParamStaticRow{
		row: row,
	}
}

type MultipleParamStaticRow struct {
	row *sql.Row
}

func (t MultipleParamStaticRow) Scan(arg1, arg2 *int, arg3 *bool, arg4 *string) error {
	var (
		c0 sql.NullInt64
		c1 sql.NullInt64
		c2 sql.NullBool
		c3 sql.NullString
	)

	if err := t.row.Scan(&c0, &c1, &c2, &c3); err != nil {
		return err
	}

	if c0.Valid {
		tmp := int(c0.Int64)
		*arg1 = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		*arg2 = tmp
	}

	if c2.Valid {
		tmp := c2.Bool
		*arg3 = tmp
	}

	if c3.Valid {
		tmp := c3.String
		*arg4 = tmp
	}

	return nil
}

// NewMultipleParamDynamic creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func NewMultipleParamDynamic(rows *sql.Rows, err error) MultipleParam {
	if err != nil {
		return errMultipleParam{e: err}
	}

	return multipleParamDynamic{
		Rows: rows,
	}
}

type multipleParamDynamic struct {
	Rows *sql.Rows
}

func (t multipleParamDynamic) Scan(arg1, arg2 *int, arg3 *bool, arg4 *string) error {
	const (
		arg10 = "arg1"
		arg21 = "arg2"
		arg32 = "arg3"
		arg43 = "arg4"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullInt64
		c1      sql.NullInt64
		c2      sql.NullBool
		c3      sql.NullString
	)

	if columns, err = t.Rows.Columns(); err != nil {
		return err
	}

	dst = make([]interface{}, 0, len(columns))

	for _, column := range columns {
		switch column {
		case arg10:
			dst = append(dst, &c0)
		case arg21:
			dst = append(dst, &c1)
		case arg32:
			dst = append(dst, &c2)
		case arg43:
			dst = append(dst, &c3)
		default:
			dst = append(dst, &ignored)
		}
	}

	if err := t.Rows.Scan(dst...); err != nil {
		return err
	}

	for _, column := range columns {
		switch column {
		case arg10:
			if c0.Valid {
				tmp := int(c0.Int64)
				*arg1 = tmp
			}
		case arg21:
			if c1.Valid {
				tmp := int(c1.Int64)
				*arg2 = tmp
			}
		case arg32:
			if c2.Valid {
				tmp := c2.Bool
				*arg3 = tmp
			}
		case arg43:
			if c3.Valid {
				tmp := c3.String
				*arg4 = tmp
			}
		}
	}

	return t.Rows.Err()
}

func (t multipleParamDynamic) Err() error {
	return t.Rows.Err()
}

func (t multipleParamDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t multipleParamDynamic) Next() bool {
	return t.Rows.Next()
}
