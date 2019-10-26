// +build !genieql.ignore

package autocompile

import (
	"context"
	"database/sql"
	"time"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql auto -o genieql.gen.go
// invoked by go generate @ autocompile/autocompile.go line 3

// Example1 structure generated by genieql.
type Example1 struct {
	CreatedAt time.Time
	ID        int
	TextField *string
	UpdatedAt time.Time
	UUIDField string
}

// Example2 structure generated by genieql.
type Example2 struct {
	BoolField bool
	CreatedAt time.Time
	TextField string
	UpdatedAt time.Time
	UUIDField string
}

// Example3 structure generated by genieql.
type Example3 struct {
	BoolField bool
	CreatedAt time.Time
	TextField string
	UpdatedAt time.Time
	UUIDField string
}

// CustomScanner scanner interface.
type CustomScanner interface {
	Scan(i1, i2 *int, b1 *bool, t1 *time.Time) error
	Next() bool
	Close() error
	Err() error
}

type errCustomScanner struct {
	e error
}

func (t errCustomScanner) Scan(i1, i2 *int, b1 *bool, t1 *time.Time) error {
	return t.e
}

func (t errCustomScanner) Next() bool {
	return false
}

func (t errCustomScanner) Err() error {
	return t.e
}

func (t errCustomScanner) Close() error {
	return nil
}

// CustomScannerStaticColumns generated by genieql
var CustomScannerStaticColumns = `"i1","i2","b1","t1"`

// NewCustomScannerStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewCustomScannerStatic(rows *sql.Rows, err error) CustomScanner {
	if err != nil {
		return errCustomScanner{e: err}
	}

	return customScannerStatic{
		Rows: rows,
	}
}

// customScannerStatic generated by genieql
type customScannerStatic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t customScannerStatic) Scan(i1, i2 *int, b1 *bool, t1 *time.Time) error {
	var (
		c0 sql.NullInt64
		c1 sql.NullInt64
		c2 sql.NullBool
		c3 sql.NullTime
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

// Err generated by genieql
func (t customScannerStatic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t customScannerStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t customScannerStatic) Next() bool {
	return t.Rows.Next()
}

// NewCustomScannerStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewCustomScannerStaticRow(row *sql.Row) CustomScannerStaticRow {
	return CustomScannerStaticRow{
		row: row,
	}
}

// CustomScannerStaticRow generated by genieql
type CustomScannerStaticRow struct {
	row *sql.Row
}

// Scan generated by genieql
func (t CustomScannerStaticRow) Scan(i1, i2 *int, b1 *bool, t1 *time.Time) error {
	var (
		c0 sql.NullInt64
		c1 sql.NullInt64
		c2 sql.NullBool
		c3 sql.NullTime
	)

	if err := t.row.Scan(&c0, &c1, &c2, &c3); err != nil {
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

	return nil
}

// NewCustomScannerDynamic creates a scanner that operates on a dynamic
// set of columns that can be returned in any subset/order.
func NewCustomScannerDynamic(rows *sql.Rows, err error) CustomScanner {
	if err != nil {
		return errCustomScanner{e: err}
	}

	return customScannerDynamic{
		Rows: rows,
	}
}

// customScannerDynamic generated by genieql
type customScannerDynamic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t customScannerDynamic) Scan(i1, i2 *int, b1 *bool, t1 *time.Time) error {
	const (
		cn0 = "i1"
		cn1 = "i2"
		cn2 = "b1"
		cn3 = "t1"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullInt64
		c1      sql.NullInt64
		c2      sql.NullBool
		c3      sql.NullTime
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
				*i1 = tmp
			}
		case cn1:
			if c1.Valid {
				tmp := int(c1.Int64)
				*i2 = tmp
			}
		case cn2:
			if c2.Valid {
				tmp := c2.Bool
				*b1 = tmp
			}
		case cn3:
			if c3.Valid {
				tmp := c3.Time
				*t1 = tmp
			}
		}
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t customScannerDynamic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t customScannerDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t customScannerDynamic) Next() bool {
	return t.Rows.Next()
}

// Example1Scanner scanner interface.
type Example1Scanner interface {
	Scan(sp0 *Example1) error
	Next() bool
	Close() error
	Err() error
}

type errExample1Scanner struct {
	e error
}

func (t errExample1Scanner) Scan(sp0 *Example1) error {
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

// Example1ScannerStaticColumns generated by genieql
var Example1ScannerStaticColumns = `"created_at","id","text_field","updated_at","uuid_field"`

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

// example1ScannerStatic generated by genieql
type example1ScannerStatic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t example1ScannerStatic) Scan(sp0 *Example1) error {
	var (
		c0 sql.NullTime
		c1 sql.NullInt64
		c2 sql.NullString
		c3 sql.NullTime
		c4 sql.NullString
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Time
		sp0.CreatedAt = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		sp0.ID = tmp
	}

	if c2.Valid {
		tmp := c2.String
		sp0.TextField = &tmp
	}

	if c3.Valid {
		tmp := c3.Time
		sp0.UpdatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.String
		sp0.UUIDField = tmp
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t example1ScannerStatic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t example1ScannerStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
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

// Example1ScannerStaticRow generated by genieql
type Example1ScannerStaticRow struct {
	row *sql.Row
}

// Scan generated by genieql
func (t Example1ScannerStaticRow) Scan(sp0 *Example1) error {
	var (
		c0 sql.NullTime
		c1 sql.NullInt64
		c2 sql.NullString
		c3 sql.NullTime
		c4 sql.NullString
	)

	if err := t.row.Scan(&c0, &c1, &c2, &c3, &c4); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Time
		sp0.CreatedAt = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		sp0.ID = tmp
	}

	if c2.Valid {
		tmp := c2.String
		sp0.TextField = &tmp
	}

	if c3.Valid {
		tmp := c3.Time
		sp0.UpdatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.String
		sp0.UUIDField = tmp
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

// example1ScannerDynamic generated by genieql
type example1ScannerDynamic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t example1ScannerDynamic) Scan(sp0 *Example1) error {
	const (
		cn0 = "created_at"
		cn1 = "id"
		cn2 = "text_field"
		cn3 = "updated_at"
		cn4 = "uuid_field"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullTime
		c1      sql.NullInt64
		c2      sql.NullString
		c3      sql.NullTime
		c4      sql.NullString
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
				tmp := c0.Time
				sp0.CreatedAt = tmp
			}
		case cn1:
			if c1.Valid {
				tmp := int(c1.Int64)
				sp0.ID = tmp
			}
		case cn2:
			if c2.Valid {
				tmp := c2.String
				sp0.TextField = &tmp
			}
		case cn3:
			if c3.Valid {
				tmp := c3.Time
				sp0.UpdatedAt = tmp
			}
		case cn4:
			if c4.Valid {
				tmp := c4.String
				sp0.UUIDField = tmp
			}
		}
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t example1ScannerDynamic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t example1ScannerDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t example1ScannerDynamic) Next() bool {
	return t.Rows.Next()
}

// Example2Scanner scanner interface.
type Example2Scanner interface {
	Scan(sp0 *Example2) error
	Next() bool
	Close() error
	Err() error
}

type errExample2Scanner struct {
	e error
}

func (t errExample2Scanner) Scan(sp0 *Example2) error {
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

// Example2ScannerStaticColumns generated by genieql
var Example2ScannerStaticColumns = `"bool_field","created_at","text_field","updated_at","uuid_field"`

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

// example2ScannerStatic generated by genieql
type example2ScannerStatic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t example2ScannerStatic) Scan(sp0 *Example2) error {
	var (
		c0 sql.NullBool
		c1 sql.NullTime
		c2 sql.NullString
		c3 sql.NullTime
		c4 sql.NullString
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Bool
		sp0.BoolField = tmp
	}

	if c1.Valid {
		tmp := c1.Time
		sp0.CreatedAt = tmp
	}

	if c2.Valid {
		tmp := c2.String
		sp0.TextField = tmp
	}

	if c3.Valid {
		tmp := c3.Time
		sp0.UpdatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.String
		sp0.UUIDField = tmp
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t example2ScannerStatic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t example2ScannerStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
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

// Example2ScannerStaticRow generated by genieql
type Example2ScannerStaticRow struct {
	row *sql.Row
}

// Scan generated by genieql
func (t Example2ScannerStaticRow) Scan(sp0 *Example2) error {
	var (
		c0 sql.NullBool
		c1 sql.NullTime
		c2 sql.NullString
		c3 sql.NullTime
		c4 sql.NullString
	)

	if err := t.row.Scan(&c0, &c1, &c2, &c3, &c4); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Bool
		sp0.BoolField = tmp
	}

	if c1.Valid {
		tmp := c1.Time
		sp0.CreatedAt = tmp
	}

	if c2.Valid {
		tmp := c2.String
		sp0.TextField = tmp
	}

	if c3.Valid {
		tmp := c3.Time
		sp0.UpdatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.String
		sp0.UUIDField = tmp
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

// example2ScannerDynamic generated by genieql
type example2ScannerDynamic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t example2ScannerDynamic) Scan(sp0 *Example2) error {
	const (
		cn0 = "bool_field"
		cn1 = "created_at"
		cn2 = "text_field"
		cn3 = "updated_at"
		cn4 = "uuid_field"
	)
	var (
		ignored sql.RawBytes
		err     error
		columns []string
		dst     []interface{}
		c0      sql.NullBool
		c1      sql.NullTime
		c2      sql.NullString
		c3      sql.NullTime
		c4      sql.NullString
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
				sp0.BoolField = tmp
			}
		case cn1:
			if c1.Valid {
				tmp := c1.Time
				sp0.CreatedAt = tmp
			}
		case cn2:
			if c2.Valid {
				tmp := c2.String
				sp0.TextField = tmp
			}
		case cn3:
			if c3.Valid {
				tmp := c3.Time
				sp0.UpdatedAt = tmp
			}
		case cn4:
			if c4.Valid {
				tmp := c4.String
				sp0.UUIDField = tmp
			}
		}
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t example2ScannerDynamic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t example2ScannerDynamic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t example2ScannerDynamic) Next() bool {
	return t.Rows.Next()
}

// CombinedScanner scanner interface.
type CombinedScanner interface {
	Scan(e1 *Example1, e2 *Example2) error
	Next() bool
	Close() error
	Err() error
}

type errCombinedScanner struct {
	e error
}

func (t errCombinedScanner) Scan(e1 *Example1, e2 *Example2) error {
	return t.e
}

func (t errCombinedScanner) Next() bool {
	return false
}

func (t errCombinedScanner) Err() error {
	return t.e
}

func (t errCombinedScanner) Close() error {
	return nil
}

// NewCombinedScannerStatic creates a scanner that operates on a static
// set of columns that are always returned in the same order.
func NewCombinedScannerStatic(rows *sql.Rows, err error) CombinedScanner {
	if err != nil {
		return errCombinedScanner{e: err}
	}

	return combinedScannerStatic{
		Rows: rows,
	}
}

// combinedScannerStatic generated by genieql
type combinedScannerStatic struct {
	Rows *sql.Rows
}

// Scan generated by genieql
func (t combinedScannerStatic) Scan(e1 *Example1, e2 *Example2) error {
	var (
		c0 sql.NullTime
		c1 sql.NullInt64
		c2 sql.NullString
		c3 sql.NullTime
		c4 sql.NullString
		c5 sql.NullBool
		c6 sql.NullTime
		c7 sql.NullString
		c8 sql.NullTime
		c9 sql.NullString
	)

	if err := t.Rows.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Time
		e1.CreatedAt = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		e1.ID = tmp
	}

	if c2.Valid {
		tmp := c2.String
		e1.TextField = &tmp
	}

	if c3.Valid {
		tmp := c3.Time
		e1.UpdatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.String
		e1.UUIDField = tmp
	}

	if c5.Valid {
		tmp := c5.Bool
		e2.BoolField = tmp
	}

	if c6.Valid {
		tmp := c6.Time
		e2.CreatedAt = tmp
	}

	if c7.Valid {
		tmp := c7.String
		e2.TextField = tmp
	}

	if c8.Valid {
		tmp := c8.Time
		e2.UpdatedAt = tmp
	}

	if c9.Valid {
		tmp := c9.String
		e2.UUIDField = tmp
	}

	return t.Rows.Err()
}

// Err generated by genieql
func (t combinedScannerStatic) Err() error {
	return t.Rows.Err()
}

// Close generated by genieql
func (t combinedScannerStatic) Close() error {
	if t.Rows == nil {
		return nil
	}
	return t.Rows.Close()
}

// Next generated by genieql
func (t combinedScannerStatic) Next() bool {
	return t.Rows.Next()
}

// NewCombinedScannerStaticRow creates a scanner that operates on a static
// set of columns that are always returned in the same order, only scans a single row.
func NewCombinedScannerStaticRow(row *sql.Row) CombinedScannerStaticRow {
	return CombinedScannerStaticRow{
		row: row,
	}
}

// CombinedScannerStaticRow generated by genieql
type CombinedScannerStaticRow struct {
	row *sql.Row
}

// Scan generated by genieql
func (t CombinedScannerStaticRow) Scan(e1 *Example1, e2 *Example2) error {
	var (
		c0 sql.NullTime
		c1 sql.NullInt64
		c2 sql.NullString
		c3 sql.NullTime
		c4 sql.NullString
		c5 sql.NullBool
		c6 sql.NullTime
		c7 sql.NullString
		c8 sql.NullTime
		c9 sql.NullString
	)

	if err := t.row.Scan(&c0, &c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9); err != nil {
		return err
	}

	if c0.Valid {
		tmp := c0.Time
		e1.CreatedAt = tmp
	}

	if c1.Valid {
		tmp := int(c1.Int64)
		e1.ID = tmp
	}

	if c2.Valid {
		tmp := c2.String
		e1.TextField = &tmp
	}

	if c3.Valid {
		tmp := c3.Time
		e1.UpdatedAt = tmp
	}

	if c4.Valid {
		tmp := c4.String
		e1.UUIDField = tmp
	}

	if c5.Valid {
		tmp := c5.Bool
		e2.BoolField = tmp
	}

	if c6.Valid {
		tmp := c6.Time
		e2.CreatedAt = tmp
	}

	if c7.Valid {
		tmp := c7.String
		e2.TextField = tmp
	}

	if c8.Valid {
		tmp := c8.Time
		e2.UpdatedAt = tmp
	}

	if c9.Valid {
		tmp := c9.String
		e2.UUIDField = tmp
	}

	return nil
}

func Example1FindByX1(ctx context.Context, q sqlx.Queryer, i1, i2 int) Example1ScannerStaticRow {
	const query = `SELECT "created_at","id","text_field","updated_at","uuid_field" FROM example1 WHERE id = $1 AND foo = $2`
	return NewExample1ScannerStaticRow(q.QueryRowContext(ctx, query, i1, i2))
}

// Example1InsertStaticColumns generated by genieql
var Example1InsertStaticColumns = `$1,$2,$3,$4,DEFAULT`

func Example1InsertExplode(arg1 *Example1) []interface{} {
	return []interface{}{arg1.CreatedAt, arg1.ID, arg1.TextField, arg1.UpdatedAt}
}

func Example1Insert(ctx context.Context, q sqlx.Queryer, e Example1) Example1ScannerStaticRow {
	const query = `INSERT INTO example1 ("created_at","id","text_field","updated_at","uuid_field") VALUES ($1,$2,$3,$4,DEFAULT),($5,$6,$7,$8,DEFAULT),($9,$10,$11,$12,DEFAULT),($13,$14,$15,$16,DEFAULT),($17,$18,$19,$20,DEFAULT),($21,$22,$23,$24,DEFAULT),($25,$26,$27,$28,DEFAULT),($29,$30,$31,$32,DEFAULT),($33,$34,$35,$36,DEFAULT),($37,$38,$39,$40,DEFAULT) RETURNING "created_at","id","text_field","updated_at","uuid_field"`
	return NewExample1ScannerStaticRow(q.QueryRowContext(ctx, query, e.CreatedAt, e.ID, e.TextField, e.UpdatedAt))
}


