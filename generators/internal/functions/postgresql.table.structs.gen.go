package functions

import (
	"net"
	"time"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental structure table constants -o postgresql.table.structs.gen.go
// invoked by go generate @ functions/functions.go line 3

// Example1 generated by genieql
type Example1 struct {
	BigintField          int
	BitField             []byte
	BitVaryingField      []byte
	BoolField            bool
	ByteArrayField       []byte
	CharacterField       string
	CharacterFixedField  string
	CidrField            net.IPNet
	DecimalField         float64
	DoublePrecisionField float64
	InetField            net.IP
	Int2Array            []int
	Int4Array            []int
	Int8Array            []int
	IntField             int
	IntervalField        time.Duration
	JSONField            []byte
	JsonbField           []byte
	MacaddrField         net.HardwareAddr
	NumericField         float64
	RealField            float32
	SmallintField        int
	TextField            string
	TimestampField       time.Time
	UUIDArray            []string
	UUIDField            string
}

// Example2 generated by genieql
type Example2 struct {
	BoolField      bool
	Int4Array      []int
	Int8Array      []int
	TextField      string
	TimestampField time.Time
	UUIDArray      []string
	UUIDField      string
}

// Example3 generated by genieql
type Example3 struct {
	Created time.Time
	Email   *string
	ID      int
	Updated time.Time
}

// Example4 generated by genieql
type Example4 struct {
	Created time.Time
	Email   string
	ID      string
	Updated time.Time
}
