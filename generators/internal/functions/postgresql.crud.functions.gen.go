package functions

import (
	"database/sql"
	"time"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental crud -o postgresql.crud.functions.gen.go --table=example1 --scanner=NewExample1ScannerDynamic --unique-scanner=NewExample1ScannerStaticRow bitbucket.org/jatone/genieql/generators/internal/functions.Example1
// invoked by go generate @ functions/functions.go line 5

func Example1Insert(q *sql.DB, arg1 Example1) Example1ScannerStaticRow {
	const query = `INSERT INTO example1 (created_at,id,text_field,updated_at,uuid_field) VALUES ($1,$2,$3,$4,$5) RETURNING created_at,id,text_field,updated_at,uuid_field`
	return NewExample1ScannerStaticRow(q.QueryRow(query, arg1.CreatedAt, arg1.ID, arg1.TextField, arg1.UpdatedAt, arg1.UUIDField))
}

func Example1FindByCreatedAt(q *sql.DB, createdAt time.Time) Example1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE created_at = $1`
	return NewExample1ScannerDynamic(q.Query(query, createdAt))
}

func Example1FindByID(q *sql.DB, id int) Example1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE id = $1`
	return NewExample1ScannerDynamic(q.Query(query, id))
}

func Example1FindByTextField(q *sql.DB, textField string) Example1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE text_field = $1`
	return NewExample1ScannerDynamic(q.Query(query, textField))
}

func Example1FindByUpdatedAt(q *sql.DB, updatedAt time.Time) Example1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE updated_at = $1`
	return NewExample1ScannerDynamic(q.Query(query, updatedAt))
}

func Example1FindByUUIDField(q *sql.DB, uuidField string) Example1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE uuid_field = $1`
	return NewExample1ScannerDynamic(q.Query(query, uuidField))
}

func Example1FindByKey(q *sql.DB, id int) Example1ScannerStaticRow {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE id = $1`
	return NewExample1ScannerStaticRow(q.QueryRow(query, id))
}

func Example1UpdateByID(q *sql.DB, id int, update Example1) Example1ScannerStaticRow {
	const query = `UPDATE example1 SET created_at = $1, text_field = $2, updated_at = $3, uuid_field = $4 WHERE id = $5 RETURNING created_at,id,text_field,updated_at,uuid_field`
	return NewExample1ScannerStaticRow(q.QueryRow(query, update.CreatedAt, update.TextField, update.UpdatedAt, update.UUIDField, id))
}

func Example1DeleteByID(q *sql.DB, id int) Example1ScannerStaticRow {
	const query = `DELETE FROM example1 WHERE id = $1 RETURNING created_at,id,text_field,updated_at,uuid_field`
	return NewExample1ScannerStaticRow(q.QueryRow(query, id))
}
