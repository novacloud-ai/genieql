package genieql

import (
	"database/sql"
	"fmt"
)

func ConnectDB() (*sql.DB, error) {
	dbconf := struct {
		Driver         string
		User           string
		Name           string
		SSLMode        string
		Port           int
		MaxConnections int
	}{
		Driver:         "postgres",
		User:           "jatone",
		Name:           "sso",
		SSLMode:        "disable",
		Port:           5432,
		MaxConnections: 10,
	}

	return sql.Open(dbconf.Driver, fmt.Sprintf("user=%s dbname=%s sslmode=%s port=%d", dbconf.User, dbconf.Name, dbconf.SSLMode, dbconf.Port))
}

// Mapping - structure that holds a mapping of the aliases to fields
type Mapping struct {
	Package string
	Type    string
	Fields  []Field
}

func (t Mapping) FindMatch(column string) (Field, bool) {
	for _, field := range t.Fields {
		if field.IsAlias(column) {
			return field, true
		}
	}

	return Field{}, false
}

type Field struct {
	Type        string
	StructField string
	Aliases     []string
}

func (t Field) IsAlias(s string) bool {
	for _, alias := range t.Aliases {
		if alias == s {
			return true
		}
	}

	return false
}

type Match struct {
	Mapping
	ArgPosition int
	Field
	ScanPosition int
}
