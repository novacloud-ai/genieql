package example

import "bitbucket.org/jatone/genieql/internal/sqlx"

func queryFunction2(q sqlx.Queryer, query string, arg1 int) ExampleScanner {
	return StaticExampleScanner(q.Query(query, arg1))
}
