package basic

import (
	"time"

	"bitbucket.org/jatone/genieql/internal/sqlx"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental crud --output=example_crud_functions.gen.go --table=example3 --unique-scanner=NewExampleScannerStaticRow --scanner=NewExampleScannerStatic example
// invoked by go generate @ basic/example.go line 12

func exampleInsert(q sqlx.Queryer, arg1 example) ExampleScannerStaticRow {
	const query = `INSERT INTO example3 ("created","email","id","updated") VALUES ($1,$2,$3,$4) RETURNING "created","email","id","updated"`
	return NewExampleScannerStaticRow(q.QueryRow(query, arg1.ID, arg1.Email, arg1.Created, arg1.Updated))
}

func exampleFindByCreated(q sqlx.Queryer, created time.Time) ExampleScannerStaticRow {
	const query = `SELECT "created","email","id","updated" FROM example3 WHERE "created" = $1`
	return NewExampleScannerStaticRow(q.QueryRow(query, created))
}

func exampleLookupByCreated(q sqlx.Queryer, created time.Time) ExampleScanner {
	const query = `SELECT "created","email","id","updated" FROM example3 WHERE "created" = $1`
	return NewExampleScannerStatic(q.Query(query, created))
}

func exampleFindByEmail(q sqlx.Queryer, email string) ExampleScannerStaticRow {
	const query = `SELECT "created","email","id","updated" FROM example3 WHERE "email" = $1`
	return NewExampleScannerStaticRow(q.QueryRow(query, email))
}

func exampleLookupByEmail(q sqlx.Queryer, email string) ExampleScanner {
	const query = `SELECT "created","email","id","updated" FROM example3 WHERE "email" = $1`
	return NewExampleScannerStatic(q.Query(query, email))
}

func exampleFindByID(q sqlx.Queryer, id int) ExampleScannerStaticRow {
	const query = `SELECT "created","email","id","updated" FROM example3 WHERE "id" = $1`
	return NewExampleScannerStaticRow(q.QueryRow(query, id))
}

func exampleLookupByID(q sqlx.Queryer, id int) ExampleScanner {
	const query = `SELECT "created","email","id","updated" FROM example3 WHERE "id" = $1`
	return NewExampleScannerStatic(q.Query(query, id))
}

func exampleFindByUpdated(q sqlx.Queryer, updated time.Time) ExampleScannerStaticRow {
	const query = `SELECT "created","email","id","updated" FROM example3 WHERE "updated" = $1`
	return NewExampleScannerStaticRow(q.QueryRow(query, updated))
}

func exampleLookupByUpdated(q sqlx.Queryer, updated time.Time) ExampleScanner {
	const query = `SELECT "created","email","id","updated" FROM example3 WHERE "updated" = $1`
	return NewExampleScannerStatic(q.Query(query, updated))
}

func exampleFindByKey(q sqlx.Queryer, id int) ExampleScannerStaticRow {
	const query = `SELECT "created","email","id","updated" FROM example3 WHERE "id" = $1`
	return NewExampleScannerStaticRow(q.QueryRow(query, id))
}

func exampleUpdateByID(q sqlx.Queryer, id int, update example) ExampleScannerStaticRow {
	const query = `UPDATE example3 SET "created" = $1, "email" = $2, "updated" = $3 WHERE "id" = $4 RETURNING "created","email","updated"`
	return NewExampleScannerStaticRow(q.QueryRow(query, update.Email, update.Created, update.Updated, id))
}

func exampleDeleteByID(q sqlx.Queryer, id int) ExampleScannerStaticRow {
	const query = `DELETE FROM example3 WHERE "id" = $1 RETURNING "created","email","id","updated"`
	return NewExampleScannerStaticRow(q.QueryRow(query, id))
}
