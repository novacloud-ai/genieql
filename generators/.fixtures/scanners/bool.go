package example

import "database/sql"

// Bool scanner interface.
type Bool interface {
	Scan(arg *bool) error
	Next() bool
	Close() error
	Err() error
}

type errBool struct {
	e error
}

func (t errBool) Scan(arg *bool) error {
	return t.e
}

func (t errBool) Next() bool {
	return false
}

func (t errBool) Err() error {
	return t.e
}

func (t errBool) Close() error {
	return nil
}

// BoolStaticColumns generated by genieql
const BoolStaticColumns = `arg`

// NewBoolStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewBoolStatic(rows *sql.Rows, err error) Bool {
	if err != nil {
		return errBool{e: err}
	}

	return boolStatic{
		Rows: rows,
	}
}

// boolStatic generated by genieql
type boolStatic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t boolStatic) Scan(arg *bool) error {
	var (
		c0 sql.NullBool
	)

	if err := t.Rows.Scan(&c0); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Bool
		*arg = tmp
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t boolStatic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t boolStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t boolStatic) Next() bool {
	return t.Rows.Next()
}

// NewBoolStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewBoolStaticRow(row *sql.Row) BoolStaticRow {
	return BoolStaticRow{
		row: row,
	}
}

// BoolStaticRow generated by genieql
type BoolStaticRow struct {
	row *sql.Row
}

// Scan generated by genieql
func (t BoolStaticRow) Scan(arg *bool) error {
	var (
		c0 sql.NullBool
	)

	if err := t.row.Scan(&c0); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Bool
		*arg = tmp
	}

	return nil
}

// NewBoolDynamic creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func NewBoolDynamic(rows *sql.Rows, err error) Bool {
	if err != nil {
		return errBool{e: err}
	}

	return boolDynamic{
		Rows: rows,
	}
}

// boolDynamic generated by genieql
type boolDynamic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t boolDynamic) Scan(arg *bool) error {
	const (
		cn0 = "arg"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullBool
	)

	if columns, err = t.Rows.Columns(); err != nil {
		return err
	}

	dst = make([]interface{}, 0, len(columns))

	for _, column := range columns {
		switch column {
		case cn0:
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
		case cn0:
			if c0.Valid {
				tmp := c0.Bool
				*arg = tmp
			}
		}
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t boolDynamic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t boolDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t boolDynamic) Next() bool {
	return t.Rows.Next()
}
