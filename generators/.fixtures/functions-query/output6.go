package example

import "bitbucket.org/jatone/genieql/internal/sqlx"

// queryFunction1 generated by genieql
func queryFunction1(q sqlx.Queryer, arg1 int) ExampleScanner {
	var query = HelloWorld
	return StaticExampleScanner(q.Query(query, arg1))
}
