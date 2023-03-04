//go:build genieql.generate
// +build genieql.generate

package autocompile

import (
	"context"
	"time"

	"bitbucket.org/jatone/genieql/examples/autocompile/pkga"
	"bitbucket.org/jatone/genieql/internal/sqlx"
	genieql "bitbucket.org/jatone/genieql/interp"
)

// Example1 ...
func Example1(gql genieql.Structure) {
	gql.From(
		gql.Table("example1"),
	)
}

// Example2 ...
func Example2(gql genieql.Structure) {
	gql.From(
		gql.Query("SELECT * FROM example2"),
	)
}

// Example3 ...
func Example3(gql genieql.Structure) {
	gql.From(
		gql.Table("example2"),
	)
}

func Timestamp(gql genieql.Structure) {
	gql.From(
		gql.Table("timestamp_examples"),
	)
}

// generates a scanner that consumes the given parameters.
func CustomScanner(gql genieql.Scanner, pattern func(i1, i2 int, b1 bool, t1 time.Time)) {}

// generates a scanner that consumes the given parameters.
func Example1Scanner(genieql.Scanner, func(Example1)) {}

// generates a scanner that consumes the given parameters.
func Example2Scanner(genieql.Scanner, func(Example2)) {}

// generates a scanner that consumes the given parameters.
func CombinedScanner(genieql.Scanner, func(e1 Example1, e2 Example2)) {}

// generates a scanner that for types from different packages
func CombinedScanner2(genieql.Scanner, func(e1 Example1, e2 pkga.Example1)) {}

// generates a scanner that consumes the given parameters.
func TimestampScanner(genieql.Scanner, func(Timestamp)) {}

// generates a function function based on the provided functional pattern and the query.
func Example1FindByX1(
	gql genieql.Function,
	pattern func(ctx context.Context, q sqlx.Queryer, i1, i2 int) NewExample1ScannerStaticRow,
) {
	gql = gql.Query("SELECT " + Example1ScannerStaticColumns + " FROM example1 WHERE id = {i1} AND foo = {i2}")
}

func Example1FindBy(gql genieql.QueryAutogen, ctx context.Context, q sqlx.Queryer, e Example1) NewExample1ScannerStaticRow {
	gql.From("example1").Ignore("created_at", "updated_at", "id")
}

func Example1LookupBy(gql genieql.QueryAutogen, ctx context.Context, q sqlx.Queryer, e Example1) NewExample1ScannerStatic {
	gql.From("example1")
}

// insert a single example1 record.
func Example1Insert(
	gql genieql.Insert,
	pattern func(ctx context.Context, q sqlx.Queryer, e Example1) NewExample1ScannerStaticRow,
) {
	gql.Into("example1").Default("uuid_field")
}

func TimestampInsert(
	gql genieql.Insert,
	pattern func(ctx context.Context, q sqlx.Queryer, e Timestamp) NewTimestampScannerStaticRow,
) {
	gql.Into("timestamp_examples").Default("id")
}

// create a merge insert
func ConflictInsert(
	gql genieql.Insert,
	pattern func(ctx context.Context, q sqlx.Queryer, e Timestamp) NewTimestampScannerStaticRow,
) {
	gql.Into("timestamp_examples").Default("id").Conflict("ON CONFLICT (id) DO NOTHING")
}
