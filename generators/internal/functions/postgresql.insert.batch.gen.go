package functions

import "bitbucket.org/jatone/genieql/internal/sqlx"

// DO NOT MODIFY: This File was auto generated by the following command:
// genieql generate insert experimental batch-function -o postgresql.insert.batch.gen.go
// invoked by go generate @ functions/functions.go line 8

// NewExample1BatchInsertFunction creates a scanner that inserts a batch of
// records into the database.
func NewExample1BatchInsertFunction(q sqlx.Queryer, p ...Example1) Example1Scanner {
	return &example1BatchInsertFunction{
		q:         q,
		remaining: p,
	}
}

type example1BatchInsertFunction struct {
	q         sqlx.Queryer
	remaining []Example1
	scanner   Example1Scanner
}

func (t *example1BatchInsertFunction) Scan(dst *Example1) error {
	return t.scanner.Scan(dst)
}

func (t *example1BatchInsertFunction) Err() error {
	if t.scanner == nil {
		return nil
	}

	return t.scanner.Err()
}

func (t *example1BatchInsertFunction) Close() error {
	if t.scanner == nil {
		return nil
	}
	return t.scanner.Close()
}

func (t *example1BatchInsertFunction) Next() bool {
	var (
		advanced bool
	)

	if t.scanner != nil && t.scanner.Next() {
		return true
	}

	// advance to the next check
	if len(t.remaining) > 0 && t.Close() == nil {
		t.scanner, t.remaining, advanced = t.advance(t.q, t.remaining...)
		return advanced && t.scanner.Next()
	}

	return false
}

func (t *example1BatchInsertFunction) advance(q sqlx.Queryer, p ...Example1) (Example1Scanner, []Example1, bool) {
	switch len(p) {
	case 0:
		return nil, []Example1(nil), false
	case 1:
		const query = `INSERT INTO example1 ("bigint_field","bit_field","bit_varying_field","bool_field","byte_array_field","character_field","character_fixed_field","cidr_field","decimal_field","double_precision_field","inet_field","int2_array","int4_array","int8_array","int_field","interval_field","json_field","jsonb_field","macaddr_field","numeric_field","real_field","smallint_field","text_field","timestamp_field","uuid_array","uuid_field") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26) RETURNING "bigint_field","bit_field","bit_varying_field","bool_field","byte_array_field","character_field","character_fixed_field","cidr_field","decimal_field","double_precision_field","inet_field","int2_array","int4_array","int8_array","int_field","interval_field","json_field","jsonb_field","macaddr_field","numeric_field","real_field","smallint_field","text_field","timestamp_field","uuid_array","uuid_field"`
		exploder := func(p ...Example1) (r [26]interface{}) {
			for idx, v := range p[:1] {
				r[idx*26+0], r[idx*26+1], r[idx*26+2], r[idx*26+3], r[idx*26+4], r[idx*26+5], r[idx*26+6], r[idx*26+7], r[idx*26+8], r[idx*26+9], r[idx*26+10], r[idx*26+11], r[idx*26+12], r[idx*26+13], r[idx*26+14], r[idx*26+15], r[idx*26+16], r[idx*26+17], r[idx*26+18], r[idx*26+19], r[idx*26+20], r[idx*26+21], r[idx*26+22], r[idx*26+23], r[idx*26+24], r[idx*26+25] = v.BigintField, v.BitField, v.BitVaryingField, v.BoolField, v.ByteArrayField, v.CharacterField, v.CharacterFixedField, v.CidrField, v.DecimalField, v.DoublePrecisionField, v.InetField, v.Int2Array, v.Int4Array, v.Int8Array, v.IntField, v.IntervalField, v.JSONField, v.JsonbField, v.MacaddrField, v.NumericField, v.RealField, v.SmallintField, v.TextField, v.TimestampField, v.UUIDArray, v.UUIDField
			}
			return
		}
		tmp := exploder(p...)
		return NewExample1ScannerStatic(q.Query(query, tmp[:]...)), []Example1(nil), true
	case 2:
		const query = `INSERT INTO example1 ("bigint_field","bit_field","bit_varying_field","bool_field","byte_array_field","character_field","character_fixed_field","cidr_field","decimal_field","double_precision_field","inet_field","int2_array","int4_array","int8_array","int_field","interval_field","json_field","jsonb_field","macaddr_field","numeric_field","real_field","smallint_field","text_field","timestamp_field","uuid_array","uuid_field") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26),($27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52) RETURNING "bigint_field","bit_field","bit_varying_field","bool_field","byte_array_field","character_field","character_fixed_field","cidr_field","decimal_field","double_precision_field","inet_field","int2_array","int4_array","int8_array","int_field","interval_field","json_field","jsonb_field","macaddr_field","numeric_field","real_field","smallint_field","text_field","timestamp_field","uuid_array","uuid_field"`
		exploder := func(p ...Example1) (r [52]interface{}) {
			for idx, v := range p[:2] {
				r[idx*26+0], r[idx*26+1], r[idx*26+2], r[idx*26+3], r[idx*26+4], r[idx*26+5], r[idx*26+6], r[idx*26+7], r[idx*26+8], r[idx*26+9], r[idx*26+10], r[idx*26+11], r[idx*26+12], r[idx*26+13], r[idx*26+14], r[idx*26+15], r[idx*26+16], r[idx*26+17], r[idx*26+18], r[idx*26+19], r[idx*26+20], r[idx*26+21], r[idx*26+22], r[idx*26+23], r[idx*26+24], r[idx*26+25] = v.BigintField, v.BitField, v.BitVaryingField, v.BoolField, v.ByteArrayField, v.CharacterField, v.CharacterFixedField, v.CidrField, v.DecimalField, v.DoublePrecisionField, v.InetField, v.Int2Array, v.Int4Array, v.Int8Array, v.IntField, v.IntervalField, v.JSONField, v.JsonbField, v.MacaddrField, v.NumericField, v.RealField, v.SmallintField, v.TextField, v.TimestampField, v.UUIDArray, v.UUIDField
			}
			return
		}
		tmp := exploder(p...)
		return NewExample1ScannerStatic(q.Query(query, tmp[:]...)), []Example1(nil), true
	case 3:
		const query = `INSERT INTO example1 ("bigint_field","bit_field","bit_varying_field","bool_field","byte_array_field","character_field","character_fixed_field","cidr_field","decimal_field","double_precision_field","inet_field","int2_array","int4_array","int8_array","int_field","interval_field","json_field","jsonb_field","macaddr_field","numeric_field","real_field","smallint_field","text_field","timestamp_field","uuid_array","uuid_field") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26),($27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52),($53,$54,$55,$56,$57,$58,$59,$60,$61,$62,$63,$64,$65,$66,$67,$68,$69,$70,$71,$72,$73,$74,$75,$76,$77,$78) RETURNING "bigint_field","bit_field","bit_varying_field","bool_field","byte_array_field","character_field","character_fixed_field","cidr_field","decimal_field","double_precision_field","inet_field","int2_array","int4_array","int8_array","int_field","interval_field","json_field","jsonb_field","macaddr_field","numeric_field","real_field","smallint_field","text_field","timestamp_field","uuid_array","uuid_field"`
		exploder := func(p ...Example1) (r [78]interface{}) {
			for idx, v := range p[:3] {
				r[idx*26+0], r[idx*26+1], r[idx*26+2], r[idx*26+3], r[idx*26+4], r[idx*26+5], r[idx*26+6], r[idx*26+7], r[idx*26+8], r[idx*26+9], r[idx*26+10], r[idx*26+11], r[idx*26+12], r[idx*26+13], r[idx*26+14], r[idx*26+15], r[idx*26+16], r[idx*26+17], r[idx*26+18], r[idx*26+19], r[idx*26+20], r[idx*26+21], r[idx*26+22], r[idx*26+23], r[idx*26+24], r[idx*26+25] = v.BigintField, v.BitField, v.BitVaryingField, v.BoolField, v.ByteArrayField, v.CharacterField, v.CharacterFixedField, v.CidrField, v.DecimalField, v.DoublePrecisionField, v.InetField, v.Int2Array, v.Int4Array, v.Int8Array, v.IntField, v.IntervalField, v.JSONField, v.JsonbField, v.MacaddrField, v.NumericField, v.RealField, v.SmallintField, v.TextField, v.TimestampField, v.UUIDArray, v.UUIDField
			}
			return
		}
		tmp := exploder(p...)
		return NewExample1ScannerStatic(q.Query(query, tmp[:]...)), []Example1(nil), true
	case 4:
		const query = `INSERT INTO example1 ("bigint_field","bit_field","bit_varying_field","bool_field","byte_array_field","character_field","character_fixed_field","cidr_field","decimal_field","double_precision_field","inet_field","int2_array","int4_array","int8_array","int_field","interval_field","json_field","jsonb_field","macaddr_field","numeric_field","real_field","smallint_field","text_field","timestamp_field","uuid_array","uuid_field") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26),($27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52),($53,$54,$55,$56,$57,$58,$59,$60,$61,$62,$63,$64,$65,$66,$67,$68,$69,$70,$71,$72,$73,$74,$75,$76,$77,$78),($79,$80,$81,$82,$83,$84,$85,$86,$87,$88,$89,$90,$91,$92,$93,$94,$95,$96,$97,$98,$99,$100,$101,$102,$103,$104) RETURNING "bigint_field","bit_field","bit_varying_field","bool_field","byte_array_field","character_field","character_fixed_field","cidr_field","decimal_field","double_precision_field","inet_field","int2_array","int4_array","int8_array","int_field","interval_field","json_field","jsonb_field","macaddr_field","numeric_field","real_field","smallint_field","text_field","timestamp_field","uuid_array","uuid_field"`
		exploder := func(p ...Example1) (r [104]interface{}) {
			for idx, v := range p[:4] {
				r[idx*26+0], r[idx*26+1], r[idx*26+2], r[idx*26+3], r[idx*26+4], r[idx*26+5], r[idx*26+6], r[idx*26+7], r[idx*26+8], r[idx*26+9], r[idx*26+10], r[idx*26+11], r[idx*26+12], r[idx*26+13], r[idx*26+14], r[idx*26+15], r[idx*26+16], r[idx*26+17], r[idx*26+18], r[idx*26+19], r[idx*26+20], r[idx*26+21], r[idx*26+22], r[idx*26+23], r[idx*26+24], r[idx*26+25] = v.BigintField, v.BitField, v.BitVaryingField, v.BoolField, v.ByteArrayField, v.CharacterField, v.CharacterFixedField, v.CidrField, v.DecimalField, v.DoublePrecisionField, v.InetField, v.Int2Array, v.Int4Array, v.Int8Array, v.IntField, v.IntervalField, v.JSONField, v.JsonbField, v.MacaddrField, v.NumericField, v.RealField, v.SmallintField, v.TextField, v.TimestampField, v.UUIDArray, v.UUIDField
			}
			return
		}
		tmp := exploder(p...)
		return NewExample1ScannerStatic(q.Query(query, tmp[:]...)), []Example1(nil), true
	default:
		const query = `INSERT INTO example1 ("bigint_field","bit_field","bit_varying_field","bool_field","byte_array_field","character_field","character_fixed_field","cidr_field","decimal_field","double_precision_field","inet_field","int2_array","int4_array","int8_array","int_field","interval_field","json_field","jsonb_field","macaddr_field","numeric_field","real_field","smallint_field","text_field","timestamp_field","uuid_array","uuid_field") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26),($27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52),($53,$54,$55,$56,$57,$58,$59,$60,$61,$62,$63,$64,$65,$66,$67,$68,$69,$70,$71,$72,$73,$74,$75,$76,$77,$78),($79,$80,$81,$82,$83,$84,$85,$86,$87,$88,$89,$90,$91,$92,$93,$94,$95,$96,$97,$98,$99,$100,$101,$102,$103,$104),($105,$106,$107,$108,$109,$110,$111,$112,$113,$114,$115,$116,$117,$118,$119,$120,$121,$122,$123,$124,$125,$126,$127,$128,$129,$130) RETURNING "bigint_field","bit_field","bit_varying_field","bool_field","byte_array_field","character_field","character_fixed_field","cidr_field","decimal_field","double_precision_field","inet_field","int2_array","int4_array","int8_array","int_field","interval_field","json_field","jsonb_field","macaddr_field","numeric_field","real_field","smallint_field","text_field","timestamp_field","uuid_array","uuid_field"`
		exploder := func(p ...Example1) (r [130]interface{}) {
			for idx, v := range p[:5] {
				r[idx*26+0], r[idx*26+1], r[idx*26+2], r[idx*26+3], r[idx*26+4], r[idx*26+5], r[idx*26+6], r[idx*26+7], r[idx*26+8], r[idx*26+9], r[idx*26+10], r[idx*26+11], r[idx*26+12], r[idx*26+13], r[idx*26+14], r[idx*26+15], r[idx*26+16], r[idx*26+17], r[idx*26+18], r[idx*26+19], r[idx*26+20], r[idx*26+21], r[idx*26+22], r[idx*26+23], r[idx*26+24], r[idx*26+25] = v.BigintField, v.BitField, v.BitVaryingField, v.BoolField, v.ByteArrayField, v.CharacterField, v.CharacterFixedField, v.CidrField, v.DecimalField, v.DoublePrecisionField, v.InetField, v.Int2Array, v.Int4Array, v.Int8Array, v.IntField, v.IntervalField, v.JSONField, v.JsonbField, v.MacaddrField, v.NumericField, v.RealField, v.SmallintField, v.TextField, v.TimestampField, v.UUIDArray, v.UUIDField
			}
			return
		}
		tmp := exploder(p[:5]...)
		return NewExample1ScannerStatic(q.Query(query, tmp[:]...)), p[5:], true
	}
}
