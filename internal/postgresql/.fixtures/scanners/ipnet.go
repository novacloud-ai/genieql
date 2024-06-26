package generated

import (
	"database/sql"
	"net"

	"github.com/jackc/pgtype"
)

// IPNet scanner interface.
type IPNet interface {
	Scan(arg1 *net.IPNet) error
	Next() bool
	Close() error
	Err() error
}

type errIPNet struct {
	e error
}

func (t errIPNet) Scan(arg1 *net.IPNet) error {
	return t.e
}

func (t errIPNet) Next() bool {
	return false
}

func (t errIPNet) Err() error {
	return t.e
}

func (t errIPNet) Close() error {
	return nil
}

// IPNetStaticColumns generated by genieql
const IPNetStaticColumns = `"arg1"`

// NewIPNetStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewIPNetStatic(rows *sql.Rows, err error) IPNet {
	if err != nil {
		return errIPNet{e: err}
	}

	return iPNetStatic{
		Rows: rows,
	}
}

// iPNetStatic generated by genieql
type iPNetStatic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t iPNetStatic) Scan(arg1 *net.IPNet) error {
	var (
		c0 pgtype.CIDR
	)

	if err := t.Rows.Scan(&c0); err != nil {
		return err
	}

	if err := c0.AssignTo(arg1); err != nil {
		return err
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t iPNetStatic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t iPNetStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t iPNetStatic) Next() bool {
	return t.Rows.Next()
}

// NewIPNetStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewIPNetStaticRow(row *sql.Row) IPNetStaticRow {
	return IPNetStaticRow{
		row: row,
	}
}

// IPNetStaticRow generated by genieql
type IPNetStaticRow struct {
	err error
	row *sql.Row
}

// Scan generated by genieql
func (t IPNetStaticRow) Scan(arg1 *net.IPNet) error {
	var (
		c0 pgtype.CIDR
	)

	if t.err != nil {
		return t.err
	}

	if err := t.row.Scan(&c0); err != nil {
		return err
	}

	if err := c0.AssignTo(arg1); err != nil {
		return err
	}

	return nil
}

// Err set an error to return by scan
func (t IPNetStaticRow) Err(err error) IPNetStaticRow {
	t.err = err
	return t
}
