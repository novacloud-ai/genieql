package examples

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate crud --config=scanner-test.config --output=type1_queries.gen.go bitbucket.org/jatone/genieql/scanner/internal/examples.Type1 type1
// invoked by go generate @ examples/type1.go line 12

const Type1Insert = `INSERT INTO type1 (field1,field2,field3,field4,field5,field6,field7,field8) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING field1,field2,field3,field4,field5,field6,field7,field8`
const Type1FindByField1 = `SELECT field1,field2,field3,field4,field5,field6,field7,field8 FROM type1 WHERE field1 = $1`
const Type1FindByField2 = `SELECT field1,field2,field3,field4,field5,field6,field7,field8 FROM type1 WHERE field2 = $1`
const Type1FindByField3 = `SELECT field1,field2,field3,field4,field5,field6,field7,field8 FROM type1 WHERE field3 = $1`
const Type1FindByField4 = `SELECT field1,field2,field3,field4,field5,field6,field7,field8 FROM type1 WHERE field4 = $1`
const Type1FindByField5 = `SELECT field1,field2,field3,field4,field5,field6,field7,field8 FROM type1 WHERE field5 = $1`
const Type1FindByField6 = `SELECT field1,field2,field3,field4,field5,field6,field7,field8 FROM type1 WHERE field6 = $1`
const Type1FindByField7 = `SELECT field1,field2,field3,field4,field5,field6,field7,field8 FROM type1 WHERE field7 = $1`
const Type1FindByField8 = `SELECT field1,field2,field3,field4,field5,field6,field7,field8 FROM type1 WHERE field8 = $1`
const Type1UpdateByID = `UPDATE type1 SET field1 = $1, field2 = $2, field3 = $3, field4 = $4, field5 = $5, field6 = $6, field7 = $7, field8 = $8 WHERE field1 = $9 RETURNING field1,field2,field3,field4,field5,field6,field7,field8`
const Type1DeleteByID = `DELETE FROM type1 WHERE field1 = $1 RETURNING field1,field2,field3,field4,field5,field6,field7,field8`
