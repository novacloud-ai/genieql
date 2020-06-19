package drivers

import (
	"bitbucket.org/jatone/genieql"
	"github.com/jackc/pgtype"
)

// implements the pgx driver https://github.com/jackc/pgx
func init() {
	genieql.RegisterDriver(PGX, NewDriver(pgx...))
}

// PGX - driver for github.com/jackc/pgx
const PGX = "github.com/jackc/pgx"

const pgxDefaultDecode = `func() {
	if err := {{ .From | expr }}.AssignTo({{.To | autoreference | expr}}); err != nil {
		return err
	}
}`

const pgxDefaultEncode = `func() {
	if err := {{ .To | expr }}.Set({{.From | autoreference | expr}}); err != nil {
		{{ error "err" | ast }}
	}
}`

var pgx = []genieql.NullableTypeDefinition{
	// cannot support OID yet... due to no field to access.
	// {Type: "pgtype.OID", Native: stringExprString, NullType: "pgtype.OIDValue", NullField: "Text"}
	{
		Type:      "pgtype.CIDR",
		Native:    cidrExpr,
		NullType:  "pgtype.CIDR",
		NullField: "IPNet",
		Decoder:   &pgtype.Inet{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.CIDRArray",
		Native:    cidrArrayExpr,
		NullType:  "pgtype.CIDRArray",
		NullField: "Elements",
		Decoder:   &pgtype.CIDRArray{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Macaddr",
		Native:    macExpr,
		NullType:  "pgtype.Macaddr",
		NullField: "Addr",
		Decoder:   &pgtype.Macaddr{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Name",
		Native:    stringExprString,
		NullType:  "pgtype.Name",
		NullField: "Text",
		Decoder:   &pgtype.Name{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Inet",
		Native:    ipExpr,
		NullType:  "pgtype.Inet",
		NullField: "IPNet",
		Decoder:   &pgtype.Inet{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Numeric",
		Native:    float64ExprString,
		NullType:  "pgtype.Numeric",
		NullField: "Int",
		Decoder:   &pgtype.Numeric{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Bytea",
		Native:    bytesExpr,
		NullType:  "pgtype.Bytea",
		NullField: "Bytes",
		Decoder:   &pgtype.Bytea{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Bit",
		Native:    bytesExpr,
		NullType:  "pgtype.Bit",
		NullField: "Bytes",
		Decoder:   &pgtype.Bit{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Varbit",
		Native:    bytesExpr,
		NullType:  "pgtype.Varbit",
		NullField: "Bytes",
		Decoder:   &pgtype.Varbit{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Bool",
		Native:    boolExprString,
		NullType:  "pgtype.Bool",
		NullField: "Bool",
		Decoder:   &pgtype.Bool{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Float4",
		Native:    float32ExprString,
		NullType:  "pgtype.Float4",
		NullField: "Float",
		Decoder:   &pgtype.Float4{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Float8",
		Native:    float64ExprString,
		NullType:  "pgtype.Float8",
		NullField: "Float",
		Decoder:   &pgtype.Float8{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Int2",
		Native:    intExprString,
		NullType:  "pgtype.Int2",
		NullField: "Int",
		Decoder:   &pgtype.Int2{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Int2Array",
		Native:    intArrExpr,
		NullType:  "pgtype.Int2Array",
		NullField: "Elements",
		Decoder:   &pgtype.Int2Array{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Int4",
		Native:    intExprString,
		NullType:  "pgtype.Int4",
		NullField: "Int",
		Decoder:   &pgtype.Int4{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Int4Array",
		Native:    intArrExpr,
		NullType:  "pgtype.Int4Array",
		NullField: "Elements",
		Decoder:   &pgtype.Int4Array{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Int8",
		Native:    intExprString,
		NullType:  "pgtype.Int8",
		NullField: "Int",
		Decoder:   &pgtype.Int8{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Int8Array",
		Native:    intArrExpr,
		NullType:  "pgtype.Int8Array",
		NullField: "Elements",
		Decoder:   &pgtype.Int8Array{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Text",
		Native:    stringExprString,
		NullType:  "pgtype.Text",
		NullField: "String",
		Decoder:   &pgtype.Text{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Varchar",
		Native:    stringExprString,
		NullType:  "pgtype.Varchar",
		NullField: "String",
		Decoder:   &pgtype.Varchar{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.BPChar",
		Native:    stringExprString,
		NullType:  "pgtype.BPChar",
		NullField: "String",
		Decoder:   &pgtype.BPChar{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Date",
		Native:    timeExprString,
		NullType:  "pgtype.Date",
		NullField: "Time",
		Decoder:   &pgtype.Date{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Timestamp",
		Native:    timeExprString,
		NullType:  "pgtype.Timestamp",
		NullField: "Time",
		Decoder:   &pgtype.Timestamp{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Timestamptz",
		Native:    timeExprString,
		NullType:  "pgtype.Timestamptz",
		NullField: "Time",
		Decoder:   &pgtype.Timestamptz{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.Interval",
		Native:    durationExpr,
		NullType:  "pgtype.Interval",
		NullField: "Microseconds",
		Decoder:   &pgtype.Interval{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.UUID",
		Native:    stringExprString,
		NullType:  "pgtype.UUID",
		NullField: "Bytes",
		Decoder:   &pgtype.UUID{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.UUIDArray",
		Native:    stringArrExpr,
		NullType:  "pgtype.UUIDArray",
		NullField: "Elements",
		Decoder:   &pgtype.UUIDArray{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.JSONB",
		Native:    bytesExpr,
		NullType:  "pgtype.JSONB",
		NullField: "Bytes",
		Decoder:   &pgtype.JSONB{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "pgtype.JSON",
		Native:    bytesExpr,
		NullType:  "pgtype.JSON",
		NullField: "Bytes",
		Decoder:   &pgtype.JSON{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "json.RawMessage",
		Native:    bytesExpr,
		NullType:  "pgtype.JSON",
		NullField: "Bytes",
		Decoder:   &pgtype.JSON{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "*json.RawMessage",
		Nullable:  true,
		Native:    bytesExpr,
		NullType:  "pgtype.JSON",
		NullField: "Bytes",
		Decoder:   &pgtype.JSON{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "net.IPNet",
		Native:    cidrExpr,
		NullType:  "pgtype.CIDR",
		NullField: "IPNet",
		Decoder:   &pgtype.Inet{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "*net.IPNet",
		Nullable:  true,
		Native:    cidrExpr,
		NullType:  "pgtype.CIDR",
		NullField: "IPNet",
		Decoder:   &pgtype.Inet{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "[]net.IPNet",
		Native:    cidrArrayExpr,
		NullType:  "pgtype.CIDRArray",
		NullField: "Elements",
		Decoder:   &pgtype.CIDRArray{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "*[]net.IPNet",
		Nullable:  true,
		Native:    cidrArrayExpr,
		NullType:  "pgtype.CIDRArray",
		NullField: "Elements",
		Decoder:   &pgtype.CIDRArray{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "net.IP",
		Native:    ipExpr,
		NullType:  "pgtype.Inet",
		NullField: "IPNet",
		Decoder:   &pgtype.Inet{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "*net.IP",
		Nullable:  true,
		Native:    ipExpr,
		NullType:  "pgtype.Inet",
		NullField: "IPNet",
		Decoder:   &pgtype.Inet{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "[]byte",
		Native:    bytesExpr,
		NullType:  "pgtype.Bytea",
		NullField: "Bytes",
		Decoder:   &pgtype.Bytea{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "*[]byte",
		Native:    bytesExpr,
		NullType:  "pgtype.Bytea",
		NullField: "Bytes",
		Decoder:   &pgtype.Bytea{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "[]string",
		Native:    stringArrExpr,
		NullType:  "pgtype.TextArray",
		NullField: "Elements",
		Decoder:   &pgtype.TextArray{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "*[]string",
		Native:    stringArrExpr,
		NullType:  "pgtype.TextArray",
		NullField: "Elements",
		Decoder:   &pgtype.TextArray{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "[]int",
		Native:    intArrExpr,
		NullType:  "pgtype.Int8Array",
		NullField: "Elements",
		Decoder:   &pgtype.Int8Array{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "*[]int",
		Native:    intArrExpr,
		NullType:  "pgtype.Int8Array",
		NullField: "Elements",
		Decoder:   &pgtype.Int8Array{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "time.Duration",
		Native:    durationExpr,
		NullType:  "pgtype.Interval",
		NullField: "Microseconds",
		Decoder:   &pgtype.Interval{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "*time.Duration",
		Native:    durationExpr,
		NullType:  "pgtype.Interval",
		NullField: "Microseconds",
		Decoder:   &pgtype.Interval{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "net.HardwareAddr",
		Native:    macExpr,
		NullType:  "pgtype.Macaddr",
		NullField: "Addr",
		Decoder:   &pgtype.Macaddr{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
	{
		Type:      "*net.HardwareAddr",
		Native:    macExpr,
		NullType:  "pgtype.Macaddr",
		NullField: "Addr",
		Decoder:   &pgtype.Macaddr{},
		Decode:    pgxDefaultDecode,
		Encode:    pgxDefaultEncode,
	},
}
