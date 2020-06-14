package generated

import "database/sql"

// BoolStaticColumns generated by genieql
const BoolStaticColumns = `"arg1"`

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
func (t boolStatic) Scan(arg1 *bool) error {
	var (
		c0 sql.NullBool
	)

	if err := t.Rows.Scan(&c0); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Bool
		*arg1 = tmp
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
func NewBoolStaticRow(row *sql.Row, err error) BoolStaticRow {
	return BoolStaticRow{
		err: err,
		row: row,
	}
}

// BoolStaticRow generated by genieql
type BoolStaticRow struct {
	err error
	row *sql.Row
}

// Scan generated by genieql
func (t BoolStaticRow) Scan(arg1 *bool) error {
	var (
		c0 sql.NullBool
	)

	if t.err != nil {
		return t.err
	}

	if err := t.row.Scan(&c0); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Bool
		*arg1 = tmp
	}

	return nil
}
