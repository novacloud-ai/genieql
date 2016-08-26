package definitions

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental scanners types -o postgresql.scanners.gen.go
// invoked by go generate @ definitions/example.go line 9

// ProfileScanner scanner interface.
type ProfileScanner interface {
	Scan(i1, i2 *int, b1 *bool, t1 *time.Time) error
	Next() bool
	Close() error
	Err() error
}

type errProfileScanner struct {
	e error
}

func (t errProfileScanner) Scan(i1, i2 *int, b1 *bool, t1 *time.Time) error {
	return t.e
}

func (t errProfileScanner) Next() bool {
	return false
}

func (t errProfileScanner) Err() error {
	return t.e
}

func (t errProfileScanner) Close() error {
	return nil
}

// StaticProfileScanner creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func StaticProfileScanner(rows *sql.Rows, err error) ProfileScanner {
	if err != nil {
		return errProfileScanner{e: err}
	}

	return staticProfileScanner{
		Rows: rows,
	}
}

type staticProfileScanner struct {
	Rows *sql.Rows
}

func (t staticProfileScanner) Scan(i1, i2 *int, b1 *bool, t1 *time.Time) error {
	var (
		c0 sql.NullInt64
		c1 sql.NullInt64
		c2 sql.NullBool
		c3 pq.NullTime
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3); err != nil {
		return err
	}

	if c0.Valid {
		tmp := int(c0.Int64)
		*i1 = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		*i2 = tmp
	}

	if c2.Valid {
		tmp := c2.Bool
		*b1 = tmp
	}

	if c3.Valid {
		tmp := c3.Time
		*t1 = tmp
	}

	return t.Rows.Err()
}

func (t staticProfileScanner) Err() error {
	return t.Rows.Err()
}

func (t staticProfileScanner) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t staticProfileScanner) Next() bool {
	return t.Rows.Next()
}

// DynamicProfileScanner creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func DynamicProfileScanner(rows *sql.Rows, err error) ProfileScanner {
	if err != nil {
		return errProfileScanner{e: err}
	}

	return dynamicProfileScanner{
		Rows: rows,
	}
}

type dynamicProfileScanner struct {
	Rows *sql.Rows
}

func (t dynamicProfileScanner) Scan(i1, i2 *int, b1 *bool, t1 *time.Time) error {
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullInt64
		c1      sql.NullInt64
		c2      sql.NullBool
		c3      pq.NullTime
	)

	if columns, err = t.Rows.Columns(); err != nil {
		return err
	}

	dst = make([]interface{}, 0, len(columns))

	for _, column := range columns {
		switch column {
		case "i1":
			dst = append(dst, &c0)
		case "i2":
			dst = append(dst, &c1)
		case "b1":
			dst = append(dst, &c2)
		case "t1":
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
		case "i1":
			if c0.Valid {
				tmp := int(c0.Int64)
				*i1 = tmp
			}
		case "i2":
			if c1.Valid {
				tmp := int(c1.Int64)
				*i2 = tmp
			}
		case "b1":
			if c2.Valid {
				tmp := c2.Bool
				*b1 = tmp
			}
		case "t1":
			if c3.Valid {
				tmp := c3.Time
				*t1 = tmp
			}
		}
	}

	return t.Rows.Err()
}
func (t dynamicProfileScanner) Err() error {
	return t.Rows.Err()
}

func (t dynamicProfileScanner) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t dynamicProfileScanner) Next() bool {
	return t.Rows.Next()
}

// Example1Scanner scanner interface.
type Example1Scanner interface {
	Scan(e *Example1) error
	Next() bool
	Close() error
	Err() error
}

type errExample1Scanner struct {
	e error
}

func (t errExample1Scanner) Scan(e *Example1) error {
	return t.e
}

func (t errExample1Scanner) Next() bool {
	return false
}

func (t errExample1Scanner) Err() error {
	return t.e
}

func (t errExample1Scanner) Close() error {
	return nil
}

// StaticExample1Scanner creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func StaticExample1Scanner(rows *sql.Rows, err error) Example1Scanner {
	if err != nil {
		return errExample1Scanner{e: err}
	}

	return staticExample1Scanner{
		Rows: rows,
	}
}

type staticExample1Scanner struct {
	Rows *sql.Rows
}

func (t staticExample1Scanner) Scan(e *Example1) error {
	var (
		c0 sql.NullInt64
		c1 sql.NullString
		c2 sql.NullString
		c3 pq.NullTime
		c4 pq.NullTime
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4); err != nil {
		return err
	}

	if c0.Valid {
		tmp := int(c0.Int64)
		e.ID = tmp
	}

	if c1.Valid {
		tmp := c1.String
		e.TextField = &tmp
	}

	if c2.Valid {
		tmp := c2.String
		e.UUIDField = tmp
	}

	if c3.Valid {
		tmp := c3.Time
		e.CreatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.Time
		e.UpdatedAt = tmp
	}

	return t.Rows.Err()
}

func (t staticExample1Scanner) Err() error {
	return t.Rows.Err()
}

func (t staticExample1Scanner) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t staticExample1Scanner) Next() bool {
	return t.Rows.Next()
}

// DynamicExample1Scanner creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func DynamicExample1Scanner(rows *sql.Rows, err error) Example1Scanner {
	if err != nil {
		return errExample1Scanner{e: err}
	}

	return dynamicExample1Scanner{
		Rows: rows,
	}
}

type dynamicExample1Scanner struct {
	Rows *sql.Rows
}

func (t dynamicExample1Scanner) Scan(e *Example1) error {
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullInt64
		c1      sql.NullString
		c2      sql.NullString
		c3      pq.NullTime
		c4      pq.NullTime
	)

	if columns, err = t.Rows.Columns(); err != nil {
		return err
	}

	dst = make([]interface{}, 0, len(columns))

	for _, column := range columns {
		switch column {
		case "id":
			dst = append(dst, &c0)
		case "text_field":
			dst = append(dst, &c1)
		case "uuid_field":
			dst = append(dst, &c2)
		case "created_at":
			dst = append(dst, &c3)
		case "updated_at":
			dst = append(dst, &c4)
		default:
			dst = append(dst, &ignored)
		}
	}

	if err := t.Rows.Scan(dst...); err != nil {
		return err
	}

	for _, column := range columns {
		switch column {
		case "id":
			if c0.Valid {
				tmp := int(c0.Int64)
				e.ID = tmp
			}
		case "text_field":
			if c1.Valid {
				tmp := c1.String
				e.TextField = &tmp
			}
		case "uuid_field":
			if c2.Valid {
				tmp := c2.String
				e.UUIDField = tmp
			}
		case "created_at":
			if c3.Valid {
				tmp := c3.Time
				e.CreatedAt = tmp
			}
		case "updated_at":
			if c4.Valid {
				tmp := c4.Time
				e.UpdatedAt = tmp
			}
		}
	}

	return t.Rows.Err()
}
func (t dynamicExample1Scanner) Err() error {
	return t.Rows.Err()
}

func (t dynamicExample1Scanner) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t dynamicExample1Scanner) Next() bool {
	return t.Rows.Next()
}