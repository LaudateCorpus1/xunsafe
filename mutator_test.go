package xunsafe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestField_Mutator(t *testing.T) {

	type Bar struct {
		ID int
	}

	type Struct1 struct {
		I   int
		I64 int64
		I32 int32
		I16 int16
		I8  int8

		UI   uint
		UI64 uint64
		UI32 uint32
		UI16 uint16
		UI8  uint8

		F64 float64
		F32 float32

		B    bool
		S    string
		Bs   []byte
		T    time.Time
		Bars []Bar
	}

	type Struct2 struct {
		I   *int
		I64 *int64
		I32 *int32
		I16 *int16
		I8  *int8

		UI   *uint
		UI64 *uint64
		UI32 *uint32
		UI16 *uint16
		UI8  *uint8

		F64 *float64
		F32 *float32

		B  *bool
		S  *string
		Bs *[]byte
		T  *time.Time
	}

	aStruct1 := &Struct1{
		I:   1,
		I64: 2,
		I32: 3,
		I16: 4,
		I8:  5,

		UI:   6,
		UI64: 7,
		UI32: 8,
		UI16: 9,
		UI8:  10,

		F64: 11,
		F32: 12,

		B:  true,
		S:  "test",
		Bs: []byte("abc"),
		T:  time.Now(),
		Bars: []Bar{
			{
				ID: 1,
			},
		},
	}

	var testCases = []struct {
		description string
		expect      interface{}
		actual      func() interface{}
		name        string
	}{
		{
			description: "int",
			expect:      100 + aStruct1.I,
			name:        "I",
		},
		{
			description: "int64",
			expect:      100 + aStruct1.I64,
			name:        "I64",
		},
		{
			description: "int32",
			expect:      100 + aStruct1.I32,
			name:        "I32",
		},
		{
			description: "int16",
			expect:      100 + aStruct1.I16,
			name:        "I16",
		},
		{
			description: "int8",
			expect:      100 + aStruct1.I8,
			name:        "I8",
		},
		{
			description: "uint",
			expect:      100 + aStruct1.UI,
			name:        "UI",
		},
		{
			description: "uint64",
			expect:      100 + aStruct1.UI64,
			name:        "UI64",
		},
		{
			description: "uint32",
			expect:      100 + aStruct1.UI32,
			name:        "UI32",
		},
		{
			description: "uint16",
			expect:      100 + aStruct1.UI16,
			name:        "UI16",
		},
		{
			description: "uint8",
			expect:      100 + aStruct1.UI8,
			name:        "UI8",
		},
		{
			description: "string",
			expect:      "pre" + aStruct1.S,
			name:        "S",
		},
		{
			description: "bool",
			expect:      aStruct1.B,
			name:        "B",
		},
		{
			description: "float64",
			expect:      100 + aStruct1.F64,
			name:        "F64",
		},
		{
			description: "float32",
			expect:      100 + aStruct1.F32,
			name:        "F32",
		},
		{
			description: "bytes",
			expect:      aStruct1.Bs,
			name:        "Bs",
		},
		{
			description: "time",
			expect:      aStruct1.T,
			name:        "T",
		},
	}

	aStruct1Type := reflect.TypeOf(Struct1{})
	aStruct1Addr := Addr(aStruct1)

	for _, testCase := range testCases {

		var field *Field
		var aStructAddr uintptr
		var holderVal reflect.Value
		field = FieldByName(aStruct1Type, testCase.name)
		aStructAddr = aStruct1Addr
		holderVal = reflect.ValueOf(aStruct1)

		//var actual interface{}
		switch val := testCase.expect.(type) {
		case int:
			field.SetInt(aStructAddr, val)
		case int64:
			field.SetInt64(aStructAddr, val)
		case int32:
			field.SetInt32(aStructAddr, val)
		case int16:
			field.SetInt16(aStructAddr, val)
		case int8:
			field.SetInt8(aStructAddr, val)
		case uint:
			field.SetUint(aStructAddr, val)
		case uint64:
			field.SetUint64(aStructAddr, val)
		case uint32:
			field.SetUint32(aStructAddr, val)
		case uint16:
			field.SetUint16(aStructAddr, val)
		case uint8:
			field.SetUint8(aStructAddr, val)
		case bool:
			field.SetBool(aStructAddr, val)
		case string:
			field.SetString(aStructAddr, val)
		case []byte:
			field.SetBytes(aStructAddr, val)
		case float64:
			field.SetFloat64(aStructAddr, val)
		case float32:
			field.SetFloat32(aStructAddr, val)
		case time.Time:
			field.SetTime(aStructAddr, val)
		case []Bar:
			field.SetValue(aStructAddr, val)
		case *int:
			field.SetIntPtr(aStructAddr, val)
		case *int64:
			field.SetInt64Ptr(aStructAddr, val)
		case *int32:
			field.SetInt32Ptr(aStructAddr, val)
		case *int16:
			field.SetInt16Ptr(aStructAddr, val)
		case *int8:
			field.SetInt8Ptr(aStructAddr, val)
		case *uint:
			field.SetUintPtr(aStructAddr, val)
		case *uint64:
			field.SetUint64Ptr(aStructAddr, val)
		case *uint32:
			field.SetUint32Ptr(aStructAddr, val)
		case *uint16:
			field.SetUint16Ptr(aStructAddr, val)
		case *uint8:
			field.SetUint8Ptr(aStructAddr, val)
		case *bool:
			field.SetBoolPtr(aStructAddr, val)
		case *string:
			field.SetStringPtr(aStructAddr, val)
		case *[]byte:
			field.SetBytesPtr(aStructAddr, val)
		case *float64:
			field.SetFloat64Ptr(aStructAddr, val)
		case *float32:
			field.SetFloat32Ptr(aStructAddr, val)
		case *time.Time:
			field.SetTimePtr(aStructAddr, val)
		}
		actual := holderVal.Elem().FieldByName(testCase.name).Interface()
		assert.EqualValues(t, testCase.expect, actual, testCase.description)
	}

}

type MutBenchStruct struct {
	ID   int
	Name string
	Val  float32
}

var _mutBenchInstance = &MutBenchStruct{
	ID:   102,
	Name: "test",
	Val:  232.2,
}

var _mutIDField = FieldByName(reflect.TypeOf(AccBenchStruct{}), "ID")
var _mutNameField = FieldByName(reflect.TypeOf(AccBenchStruct{}), "Name")
var _mutValField = FieldByName(reflect.TypeOf(AccBenchStruct{}), "value")

func BenchmarkField_Mutator_Native(b *testing.B) {
	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_mutBenchInstance.ID = id
		_mutBenchInstance.Name = name
		_mutBenchInstance.Val = val
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}

func Benchmark_Mutator_Fast(b *testing.B) {
	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)

	addr := Addr(_mutBenchInstance)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_mutIDField.SetInt(addr, id)
		_mutNameField.SetString(addr, name)
		_mutValField.SetFloat32(addr, val)
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}

func Benchmark_Mutator_Fast_Ptr(b *testing.B) {
	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)

	holderPtr := Addr(_mutBenchInstance)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		*(_mutIDField.Addr(holderPtr).(*int)) = id
		*(_mutNameField.Addr(holderPtr).(*string)) = name
		*(_mutValField.Addr(holderPtr).(*float32)) = val
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}

func BenchmarkField_Mutator_Reflect(b *testing.B) {
	aType := reflect.TypeOf(AccBenchStruct{})
	var idFieldIdx, nameFiledIdx, valFieldIdx int

	if field, ok := aType.FieldByName("ID"); ok {
		idFieldIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Name"); ok {
		nameFiledIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("value"); ok {
		valFieldIdx = field.Index[0]
	}

	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)

	instanceVal := reflect.ValueOf(_mutBenchInstance).Elem()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		instanceVal.Field(idFieldIdx).Set(reflect.ValueOf(id))
		instanceVal.Field(nameFiledIdx).Set(reflect.ValueOf(name))
		instanceVal.Field(valFieldIdx).Set(reflect.ValueOf(val))
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}

func BenchmarkField_Mutator_Reflect_Ptr(b *testing.B) {
	aType := reflect.TypeOf(AccBenchStruct{})
	var idFieldIdx, nameFiledIdx, valFieldIdx int

	if field, ok := aType.FieldByName("ID"); ok {
		idFieldIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("Name"); ok {
		nameFiledIdx = field.Index[0]
	}
	if field, ok := aType.FieldByName("value"); ok {
		valFieldIdx = field.Index[0]
	}

	var id = 1000
	var name = "test 1000"
	var val = float32(43.4)

	instanceVal := reflect.ValueOf(_mutBenchInstance).Elem()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		*(instanceVal.Field(idFieldIdx).Addr().Interface().(*int)) = id
		*(instanceVal.Field(nameFiledIdx).Addr().Interface().(*string)) = name
		*(instanceVal.Field(valFieldIdx).Addr().Interface().(*float32)) = val
	}
	assert.EqualValues(b, _mutBenchInstance.ID, id)
	assert.EqualValues(b, _mutBenchInstance.Name, name)
	assert.EqualValues(b, _mutBenchInstance.Val, val)
}
