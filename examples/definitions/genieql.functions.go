//+build genieql,generate,functions

package definitions

import "bitbucket.org/jatone/genieql/sqlx"

type customQueryFunction func(queryer sqlx.Queryer, x1, x2, x3 int) NewProfileScannerDynamic

func customQueryFucntion2(queryer sqlx.Queryer, x1, x2, x3 int) NewProfileScannerDynamic {
	const query = query1
	return nil
}