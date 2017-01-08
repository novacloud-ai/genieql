package definitions

import "time"

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental structure table constants -o postgresql.table.structs.gen.go
// invoked by go generate @ definitions/example.go line 3

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
