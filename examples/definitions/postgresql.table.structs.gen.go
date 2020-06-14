package definitions

import (
	"net"
	"time"
)

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate experimental structure table constants -o postgresql.table.structs.gen.go
// invoked by go generate @ definitions/example.go line 3

// Example1 structure generated by genieql.
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

// Example2 structure generated by genieql.
type Example2 struct {
	BoolField      bool
	Int4Array      []int
	Int8Array      []int
	TextField      string
	TimestampField time.Time
	UUIDArray      []string
	UUIDField      string
}
