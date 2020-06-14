package drivers

import (
	"github.com/jackc/pgtype"

	"bitbucket.org/jatone/genieql"
)

// implements the lib/pq driver https://github.com/lib/pq
func init() {
	genieql.RegisterDriver(PQ, NewDriver(libpq...))
}

// PQ - driver for github.com/lib/pq
const PQ = "github.com/lib/pq"

const pqDefaultDecode = `func() {
	if err := {{ .From | expr }}.AssignTo({{.To | autoreference | expr}}); err != nil {
		return err
	}
}`

var libpq = []genieql.NullableTypeDefinition{
	// cannot support OID yet... due to no field to access.
	// {Type: "pgtype.OID", Native: stringExprString, NullType: "pgtype.OIDValue", NullField: "Text"}
	{
		Type:      "pgtype.CIDR",
		Native:    cidrExpr,
		NullType:  "pgtype.CIDR",
		NullField: "IPNet",
		Decoder:   &pgtype.Inet{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.CIDRArray",
		Native:    cidrArrayExpr,
		NullType:  "pgtype.CIDRArray",
		NullField: "Elements",
		Decoder:   &pgtype.CIDRArray{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Macaddr",
		Native:    macExpr,
		NullType:  "pgtype.Macaddr",
		NullField: "Addr",
		Decoder:   &pgtype.Macaddr{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Name",
		Native:    stringExprString,
		NullType:  "pgtype.Name",
		NullField: "Text",
		Decoder:   &pgtype.Name{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Inet",
		Native:    ipExpr,
		NullType:  "pgtype.Inet",
		NullField: "IPNet",
		Decoder:   &pgtype.Inet{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Numeric",
		Native:    float64ExprString,
		NullType:  "pgtype.Numeric",
		NullField: "Int",
		Decoder:   &pgtype.Numeric{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Bytea",
		Native:    bytesExpr,
		NullType:  "pgtype.Bytea",
		NullField: "Bytes",
		Decoder:   &pgtype.Bytea{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Bit",
		Native:    bytesExpr,
		NullType:  "pgtype.Bit",
		NullField: "Bytes",
		Decoder:   &pgtype.Bit{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Varbit",
		Native:    bytesExpr,
		NullType:  "pgtype.Varbit",
		NullField: "Bytes",
		Decoder:   &pgtype.Varbit{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Bool",
		Native:    boolExprString,
		NullType:  "pgtype.Bool",
		NullField: "Bool",
		Decoder:   &pgtype.Bool{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Float4",
		Native:    float32ExprString,
		NullType:  "pgtype.Float4",
		NullField: "Float",
		Decoder:   &pgtype.Float4{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Float8",
		Native:    float64ExprString,
		NullType:  "pgtype.Float8",
		NullField: "Float",
		Decoder:   &pgtype.Float8{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Int2",
		Native:    intExprString,
		NullType:  "pgtype.Int2",
		NullField: "Int",
		Decoder:   &pgtype.Int2{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Int2Array",
		Native:    intArrExpr,
		NullType:  "pgtype.Int2Array",
		NullField: "Elements",
		Decoder:   &pgtype.Int2Array{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Int4",
		Native:    intExprString,
		NullType:  "pgtype.Int4",
		NullField: "Int",
		Decoder:   &pgtype.Int4{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Int4Array",
		Native:    intArrExpr,
		NullType:  "pgtype.Int4Array",
		NullField: "Elements",
		Decoder:   &pgtype.Int4Array{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Int8",
		Native:    intExprString,
		NullType:  "pgtype.Int8",
		NullField: "Int",
		Decoder:   &pgtype.Int8{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Int8Array",
		Native:    intArrExpr,
		NullType:  "pgtype.Int8Array",
		NullField: "Elements",
		Decoder:   &pgtype.Int8Array{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Text",
		Native:    stringExprString,
		NullType:  "pgtype.Text",
		NullField: "String",
		Decoder:   &pgtype.Text{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Varchar",
		Native:    stringExprString,
		NullType:  "pgtype.Varchar",
		NullField: "String",
		Decoder:   &pgtype.Varchar{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.BPChar",
		Native:    stringExprString,
		NullType:  "pgtype.BPChar",
		NullField: "String",
		Decoder:   &pgtype.BPChar{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Date",
		Native:    timeExprString,
		NullType:  "pgtype.Date",
		NullField: "Time",
		Decoder:   &pgtype.Date{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Timestamp",
		Native:    timeExprString,
		NullType:  "pgtype.Timestamp",
		NullField: "Time",
		Decoder:   &pgtype.Timestamp{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Timestamptz",
		Native:    timeExprString,
		NullType:  "pgtype.Timestamptz",
		NullField: "Time",
		Decoder:   &pgtype.Timestamptz{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.Interval",
		Native:    durationExpr,
		NullType:  "pgtype.Interval",
		NullField: "Microseconds",
		Decoder:   &pgtype.Interval{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.UUID",
		Native:    stringExprString,
		NullType:  "pgtype.UUID",
		NullField: "Bytes",
		Decoder:   &pgtype.UUID{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.UUIDArray",
		Native:    stringArrExpr,
		NullType:  "pgtype.UUIDArray",
		NullField: "Elements",
		Decoder:   &pgtype.UUIDArray{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.JSONB",
		Native:    bytesExpr,
		NullType:  "pgtype.JSONB",
		NullField: "Bytes",
		Decoder:   &pgtype.JSONB{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "pgtype.JSON",
		Native:    bytesExpr,
		NullType:  "pgtype.JSON",
		NullField: "Bytes",
		Decoder:   &pgtype.JSON{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "json.RawMessage",
		Native:    bytesExpr,
		NullType:  "pgtype.JSON",
		NullField: "Bytes",
		Decoder:   &pgtype.JSON{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "*json.RawMessage",
		Nullable:  true,
		Native:    bytesExpr,
		NullType:  "pgtype.JSON",
		NullField: "Bytes",
		Decoder:   &pgtype.JSON{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "net.IPNet",
		Native:    cidrExpr,
		NullType:  "pgtype.CIDR",
		NullField: "IPNet",
		Decoder:   &pgtype.Inet{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "*net.IPNet",
		Nullable:  true,
		Native:    cidrExpr,
		NullType:  "pgtype.CIDR",
		NullField: "IPNet",
		Decoder:   &pgtype.Inet{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "[]net.IPNet",
		Native:    cidrArrayExpr,
		NullType:  "pgtype.CIDRArray",
		NullField: "Elements",
		Decoder:   &pgtype.CIDRArray{},
		Decode:    pgxDefaultDecode,
	},
	{
		Type:      "*[]net.IPNet",
		Nullable:  true,
		Native:    cidrArrayExpr,
		NullType:  "pgtype.CIDRArray",
		NullField: "Elements",
		Decoder:   &pgtype.CIDRArray{},
		Decode:    pgxDefaultDecode,
	},
}
