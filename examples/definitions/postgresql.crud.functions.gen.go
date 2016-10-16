package definitions

import (
	"database/sql"
	"time"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental crud -o postgresql.crud.functions.gen.go --table=example1 --scanner=DynamicExample1Scanner --unique-scanner=NewStaticRowExample1Scanner bitbucket.org/jatone/genieql/examples/definitions.Example1
// invoked by go generate @ definitions/example.go line 10

func Example1Insert(q *sql.DB, createdAt time.Time, id int, textField string, updatedAt time.Time, uuidField string) StaticRowExample1Scanner {
	const query = `INSERT INTO example1 (created_at,id,text_field,updated_at,uuid_field) VALUES ($1,$2,$3,$4,$5) RETURNING created_at,id,text_field,updated_at,uuid_field`
	return NewStaticRowExample1Scanner(q.QueryRow(query, createdAt, id, textField, updatedAt, uuidField))
}

func Example1FindByCreatedAt(q *sql.DB, createdAt time.Time) Example1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE created_at = $1`
	return DynamicExample1Scanner(q.Query(query, createdAt))
}

func Example1FindByID(q *sql.DB, id int) Example1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE id = $1`
	return DynamicExample1Scanner(q.Query(query, id))
}

func Example1FindByTextField(q *sql.DB, textField string) Example1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE text_field = $1`
	return DynamicExample1Scanner(q.Query(query, textField))
}

func Example1FindByUpdatedAt(q *sql.DB, updatedAt time.Time) Example1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE updated_at = $1`
	return DynamicExample1Scanner(q.Query(query, updatedAt))
}

func Example1FindByUUIDField(q *sql.DB, uuidField string) Example1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE uuid_field = $1`
	return DynamicExample1Scanner(q.Query(query, uuidField))
}

func Example1FindByKey(q *sql.DB, id int) StaticRowExample1Scanner {
	const query = `SELECT created_at,id,text_field,updated_at,uuid_field FROM example1 WHERE id = $1`
	return NewStaticRowExample1Scanner(q.QueryRow(query, id))
}

func Example1UpdateByID(q *sql.DB, id int) StaticRowExample1Scanner {
	const query = `UPDATE example1 SET created_at = $1, id = $2, text_field = $3, updated_at = $4, uuid_field = $5 WHERE id = $6 RETURNING created_at,id,text_field,updated_at,uuid_field`
	return NewStaticRowExample1Scanner(q.QueryRow(query, id))
}

func Example1DeleteByID(q *sql.DB, id int) StaticRowExample1Scanner {
	const query = `DELETE FROM example1 WHERE id = $1 RETURNING created_at,id,text_field,updated_at,uuid_field`
	return NewStaticRowExample1Scanner(q.QueryRow(query, id))
}
