package functions

import (
	"database/sql"

	"github.com/lib/pq"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental scanners types -o postgresql.scanners.gen.go
// invoked by go generate @ functions/functions.go line 4

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

const Example1ScannerStaticColumns = `"created_at","id","text_field","updated_at","uuid_field"`

// NewExample1ScannerStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewExample1ScannerStatic(rows *sql.Rows, err error) Example1Scanner {
	if err != nil {
		return errExample1Scanner{e: err}
	}

	return example1ScannerStatic{
		Rows: rows,
	}
}

type example1ScannerStatic struct {
	Rows *sql.Rows
}

func (t example1ScannerStatic) Scan(e *Example1) error {
	var (
		c0 pq.NullTime
		c1 sql.NullInt64
		c2 sql.NullString
		c3 pq.NullTime
		c4 sql.NullString
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Time
		e.CreatedAt = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		e.ID = tmp
	}

	if c2.Valid {
		tmp := c2.String
		e.TextField = &tmp
	}

	if c3.Valid {
		tmp := c3.Time
		e.UpdatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.String
		e.UUIDField = tmp
	}

	return t.Rows.Err()
}

func (t example1ScannerStatic) Err() error {
	return t.Rows.Err()
}

func (t example1ScannerStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t example1ScannerStatic) Next() bool {
	return t.Rows.Next()
}

// NewExample1ScannerStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewExample1ScannerStaticRow(row *sql.Row) Example1ScannerStaticRow {
	return Example1ScannerStaticRow{
		row: row,
	}
}

type Example1ScannerStaticRow struct {
	row *sql.Row
}

func (t Example1ScannerStaticRow) Scan(e *Example1) error {
	var (
		c0 pq.NullTime
		c1 sql.NullInt64
		c2 sql.NullString
		c3 pq.NullTime
		c4 sql.NullString
	)

	if err := t.row.Scan(&c0, &c1, &c2, &c3, &c4); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Time
		e.CreatedAt = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		e.ID = tmp
	}

	if c2.Valid {
		tmp := c2.String
		e.TextField = &tmp
	}

	if c3.Valid {
		tmp := c3.Time
		e.UpdatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.String
		e.UUIDField = tmp
	}

	return nil
}

// NewExample1ScannerDynamic creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func NewExample1ScannerDynamic(rows *sql.Rows, err error) Example1Scanner {
	if err != nil {
		return errExample1Scanner{e: err}
	}

	return example1ScannerDynamic{
		Rows: rows,
	}
}

type example1ScannerDynamic struct {
	Rows *sql.Rows
}

func (t example1ScannerDynamic) Scan(e *Example1) error {
	const (
		created_at0 = "created_at"
		id1         = "id"
		text_field2 = "text_field"
		updated_at3 = "updated_at"
		uuid_field4 = "uuid_field"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      pq.NullTime
		c1      sql.NullInt64
		c2      sql.NullString
		c3      pq.NullTime
		c4      sql.NullString
	)

	if columns, err = t.Rows.Columns(); err != nil {
		return err
	}

	dst = make([]interface{}, 0, len(columns))

	for _, column := range columns {
		switch column {
		case created_at0:
			dst = append(dst, &c0)
		case id1:
			dst = append(dst, &c1)
		case text_field2:
			dst = append(dst, &c2)
		case updated_at3:
			dst = append(dst, &c3)
		case uuid_field4:
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
		case created_at0:
			if c0.Valid {
				tmp := c0.Time
				e.CreatedAt = tmp
			}
		case id1:
			if c1.Valid {
				tmp := int(c1.Int64)
				e.ID = tmp
			}
		case text_field2:
			if c2.Valid {
				tmp := c2.String
				e.TextField = &tmp
			}
		case updated_at3:
			if c3.Valid {
				tmp := c3.Time
				e.UpdatedAt = tmp
			}
		case uuid_field4:
			if c4.Valid {
				tmp := c4.String
				e.UUIDField = tmp
			}
		}
	}

	return t.Rows.Err()
}

func (t example1ScannerDynamic) Err() error {
	return t.Rows.Err()
}

func (t example1ScannerDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t example1ScannerDynamic) Next() bool {
	return t.Rows.Next()
}

// Example2Scanner scanner interface.
type Example2Scanner interface {
	Scan(e *Example2) error
	Next() bool
	Close() error
	Err() error
}

type errExample2Scanner struct {
	e error
}

func (t errExample2Scanner) Scan(e *Example2) error {
	return t.e
}

func (t errExample2Scanner) Next() bool {
	return false
}

func (t errExample2Scanner) Err() error {
	return t.e
}

func (t errExample2Scanner) Close() error {
	return nil
}

const Example2ScannerStaticColumns = `"bool_field","created_at","text_field","updated_at","uuid_field"`

// NewExample2ScannerStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewExample2ScannerStatic(rows *sql.Rows, err error) Example2Scanner {
	if err != nil {
		return errExample2Scanner{e: err}
	}

	return example2ScannerStatic{
		Rows: rows,
	}
}

type example2ScannerStatic struct {
	Rows *sql.Rows
}

func (t example2ScannerStatic) Scan(e *Example2) error {
	var (
		c0 sql.NullBool
		c1 pq.NullTime
		c2 sql.NullString
		c3 pq.NullTime
		c4 sql.NullString
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Bool
		e.BoolField = tmp
	}

	if c1.Valid {
		tmp := c1.Time
		e.CreatedAt = tmp
	}

	if c2.Valid {
		tmp := c2.String
		e.TextField = tmp
	}

	if c3.Valid {
		tmp := c3.Time
		e.UpdatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.String
		e.UUIDField = tmp
	}

	return t.Rows.Err()
}

func (t example2ScannerStatic) Err() error {
	return t.Rows.Err()
}

func (t example2ScannerStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t example2ScannerStatic) Next() bool {
	return t.Rows.Next()
}

// NewExample2ScannerStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewExample2ScannerStaticRow(row *sql.Row) Example2ScannerStaticRow {
	return Example2ScannerStaticRow{
		row: row,
	}
}

type Example2ScannerStaticRow struct {
	row *sql.Row
}

func (t Example2ScannerStaticRow) Scan(e *Example2) error {
	var (
		c0 sql.NullBool
		c1 pq.NullTime
		c2 sql.NullString
		c3 pq.NullTime
		c4 sql.NullString
	)

	if err := t.row.Scan(&c0, &c1, &c2, &c3, &c4); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Bool
		e.BoolField = tmp
	}

	if c1.Valid {
		tmp := c1.Time
		e.CreatedAt = tmp
	}

	if c2.Valid {
		tmp := c2.String
		e.TextField = tmp
	}

	if c3.Valid {
		tmp := c3.Time
		e.UpdatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.String
		e.UUIDField = tmp
	}

	return nil
}

// NewExample2ScannerDynamic creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func NewExample2ScannerDynamic(rows *sql.Rows, err error) Example2Scanner {
	if err != nil {
		return errExample2Scanner{e: err}
	}

	return example2ScannerDynamic{
		Rows: rows,
	}
}

type example2ScannerDynamic struct {
	Rows *sql.Rows
}

func (t example2ScannerDynamic) Scan(e *Example2) error {
	const (
		bool_field0 = "bool_field"
		created_at1 = "created_at"
		text_field2 = "text_field"
		updated_at3 = "updated_at"
		uuid_field4 = "uuid_field"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullBool
		c1      pq.NullTime
		c2      sql.NullString
		c3      pq.NullTime
		c4      sql.NullString
	)

	if columns, err = t.Rows.Columns(); err != nil {
		return err
	}

	dst = make([]interface{}, 0, len(columns))

	for _, column := range columns {
		switch column {
		case bool_field0:
			dst = append(dst, &c0)
		case created_at1:
			dst = append(dst, &c1)
		case text_field2:
			dst = append(dst, &c2)
		case updated_at3:
			dst = append(dst, &c3)
		case uuid_field4:
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
		case bool_field0:
			if c0.Valid {
				tmp := c0.Bool
				e.BoolField = tmp
			}
		case created_at1:
			if c1.Valid {
				tmp := c1.Time
				e.CreatedAt = tmp
			}
		case text_field2:
			if c2.Valid {
				tmp := c2.String
				e.TextField = tmp
			}
		case updated_at3:
			if c3.Valid {
				tmp := c3.Time
				e.UpdatedAt = tmp
			}
		case uuid_field4:
			if c4.Valid {
				tmp := c4.String
				e.UUIDField = tmp
			}
		}
	}

	return t.Rows.Err()
}

func (t example2ScannerDynamic) Err() error {
	return t.Rows.Err()
}

func (t example2ScannerDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

func (t example2ScannerDynamic) Next() bool {
	return t.Rows.Next()
}
