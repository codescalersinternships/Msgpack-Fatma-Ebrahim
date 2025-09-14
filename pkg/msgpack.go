package msgpack

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

const Msgpack_Nil = 0xc0
const Msgpack_False = 0xc2
const Msgpack_True = 0xc3

const Msgpack_Uint_8 = 0xcc
const Msgpack_Uint_16 = 0xcd
const Msgpack_Uint_32 = 0xce
const Msgpack_Uint_64 = 0xcf

const Msgpack_Int_8 = 0xd0
const Msgpack_Int_16 = 0xd1
const Msgpack_Int_32 = 0xd2
const Msgpack_Int_64 = 0xd3

const Msgpack_Float_32 = 0xca
const Msgpack_Float_64 = 0xcb

const Msgpack_String_8 = 0xd9
const Msgpack_String_16 = 0xda
const Msgpack_String_32 = 0xdb

type Uint_8 struct {
	typeByte byte
	value    byte
}

type Uint struct {
	typeByte byte
	value    []byte
}

type Int_8 struct {
	typeByte byte
	value    byte
}

type Int struct {
	typeByte byte
	value    []byte
}

type Float struct {
	typeByte byte
	value    []byte
}

type String struct {
	typeByte byte
	value    []byte
	size     []byte
}

type Object struct {
	IsObject bool
	Flag     bool
	Unum     uint8
	Snum     int16
	Fnum     float32
	Str      string
}

func encodeBool(value bool) byte {
	if value {
		return byte(Msgpack_True)
	} else {
		return byte(Msgpack_False)
	}
}
func encodeUint(value uint8) Uint_8 {
	return Uint_8{
		typeByte: Msgpack_Uint_8,
		value:    value,
	}
}

func encodeUint16(value uint16) Uint {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, value)
	return Uint{
		typeByte: Msgpack_Uint_16,
		value:    bytes,
	}
}

func encodeUint32(value uint32) Uint {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, value)
	return Uint{
		typeByte: Msgpack_Uint_32,
		value:    bytes,
	}
}

func encodeUint64(value uint64) Uint {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint64(bytes, value)
	return Uint{
		typeByte: Msgpack_Uint_64,
		value:    bytes,
	}
}

func encodeInt(value int8) Int_8 {
	return Int_8{
		typeByte: Msgpack_Uint_8,
		value:    byte(value),
	}
}

func encodeInt16(value int16) Int {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, uint16(value))
	return Int{
		typeByte: Msgpack_Uint_16,
		value:    bytes,
	}
}

func encodeInt32(value int32) Int {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(value))
	return Int{
		typeByte: Msgpack_Uint_32,
		value:    bytes,
	}
}

func encodeInt64(value int64) Int {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(value))
	return Int{
		typeByte: Msgpack_Uint_64,
		value:    bytes,
	}
}

func encodeFloat32(value float32) Float {
	bin := math.Float32bits(value)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bin)
	return Float{
		typeByte: Msgpack_Float_32,
		value:    bytes,
	}
}

func encodeFloat64(value float64) Float {
	bin := math.Float64bits(value)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bin)
	return Float{
		typeByte: Msgpack_Float_64,
		value:    bytes,
	}
}
func encodeString(value string) String {
	string_8 := int(math.Pow(2, 8) - 1)
	string_16 := int(math.Pow(2, 16) - 1)
	string_32 := int(math.Pow(2, 32) - 1)
	var size []byte
	var t byte
	switch {
	case len(value) < string_8:
		size = make([]byte, 0)
		size = append(size, byte(len(value)))
		t = Msgpack_String_8
	case len(value) < string_16:
		size = make([]byte, 2)
		binary.BigEndian.PutUint16(size, uint16(len(value)))
		t = Msgpack_String_16
	case len(value) < string_32:
		size = make([]byte, 4)
		binary.BigEndian.PutUint32(size, uint32(len(value)))
		t = Msgpack_String_32
	}

	bytes := []byte(value)
	copy(bytes, value)
	return String{
		typeByte: t,
		value:    bytes,
		size:     size,
	}
}

func decodeBool(value byte) bool {
	if value == byte(Msgpack_True) {
		return true
	} else {
		return false
	}
}

func decodeUint8(value byte) uint8 {
	return uint8(value)
}

func decodeUint16(value []byte) uint16 {
	return binary.BigEndian.Uint16(value)
}

func decodeUint32(value []byte) uint32 {
	return binary.BigEndian.Uint32(value)
}
func decodeUint64(value []byte) uint64 {
	return binary.BigEndian.Uint64(value)
}

func decodeInt(value byte) int8 {
	return int8(value)
}

func decodeInt16(value []byte) int16 {
	return int16(binary.BigEndian.Uint16(value))
}
func decodeInt32(value []byte) int32 {
	return int32(binary.BigEndian.Uint32(value))
}

func decodeInt64(value []byte) int64 {
	return int64(binary.BigEndian.Uint64(value))
}

func decodeFloat(value []byte) float32 {
	bin := binary.BigEndian.Uint32(value)
	return math.Float32frombits(bin)
}

func decodeString(value []byte) string {
	return string(value)
}

func Pack(obj interface{}) ([]byte, error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	bytes := make([]byte, 0)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		memberType := field.Type.Kind()
		memberName := field.Name
		memberValue := v.Field(i)
		fmt.Println(memberType, memberName, memberValue)
		objectBytes := make([]byte, 0)

		switch memberType {
		case reflect.Bool:
			objectBytes = append(objectBytes, encodeBool(memberValue.Bool()))
			fmt.Println(decodeBool(objectBytes[0]))

		case reflect.Uint8:
			encodedUint := encodeUint(uint8(memberValue.Uint()))
			objectBytes = append(objectBytes, encodedUint.typeByte, encodedUint.value)
			fmt.Println(decodeUint8(encodedUint.value))
		case reflect.Uint16:
			encodedUint := encodeUint16(uint16(memberValue.Uint()))
			objectBytes = append(objectBytes, encodedUint.typeByte)
			objectBytes = append(objectBytes, encodedUint.value...)
			fmt.Println(decodeUint16(encodedUint.value))
		case reflect.Uint32:
			encodedUint := encodeUint32(uint32(memberValue.Uint()))
			objectBytes = append(objectBytes, encodedUint.typeByte)
			objectBytes = append(objectBytes, encodedUint.value...)
			fmt.Println(decodeUint32(encodedUint.value))
		case reflect.Uint64:
			encodedUint := encodeUint64(uint64(memberValue.Uint()))
			objectBytes = append(objectBytes, encodedUint.typeByte)
			objectBytes = append(objectBytes, encodedUint.value...)
			fmt.Println(decodeUint64(encodedUint.value))

		case reflect.Int8:
			encodedInt := encodeInt(int8(memberValue.Int()))
			objectBytes = append(objectBytes, encodedInt.typeByte, encodedInt.value)
			fmt.Println(decodeInt(encodedInt.value))
		case reflect.Int16:
			encodedInt := encodeInt16(int16(memberValue.Int()))
			objectBytes = append(objectBytes, encodedInt.typeByte)
			objectBytes = append(objectBytes, encodedInt.value...)
			fmt.Println(decodeInt16(encodedInt.value))
		case reflect.Int32:
			encodedInt := encodeInt32(int32(memberValue.Int()))
			objectBytes = append(objectBytes, encodedInt.typeByte)
			objectBytes = append(objectBytes, encodedInt.value...)
			fmt.Println(decodeInt16(encodedInt.value))
		case reflect.Int64:
			encodedInt := encodeInt64(int64(memberValue.Int()))
			objectBytes = append(objectBytes, encodedInt.typeByte)
			objectBytes = append(objectBytes, encodedInt.value...)
			fmt.Println(decodeInt16(encodedInt.value))

		case reflect.Float32:
			encodedFloat := encodeFloat32(float32(memberValue.Float()))
			objectBytes = append(objectBytes, encodedFloat.typeByte)
			objectBytes = append(objectBytes, encodedFloat.value...)
			fmt.Println(decodeFloat(encodedFloat.value))
		case reflect.Float64:
			encodedFloat := encodeFloat64(memberValue.Float())
			objectBytes = append(objectBytes, encodedFloat.typeByte)
			objectBytes = append(objectBytes, encodedFloat.value...)
			fmt.Println(decodeFloat(encodedFloat.value))

		case reflect.String:
			encodedString := encodeString(memberValue.String())
			objectBytes = append(objectBytes, encodedString.typeByte)
			objectBytes = append(objectBytes, encodedString.size...)
			objectBytes = append(objectBytes, encodedString.value...)
			fmt.Println(decodeString(encodedString.value))

		default:
			objectBytes = append(objectBytes, Msgpack_Nil)
		}

		objectBytes = append(objectBytes, memberName...)
		fmt.Println(objectBytes)
		bytes = append(bytes, objectBytes...) // to store the name
		bytes = append(bytes, byte('\n'))     // to store the separator == 10

	}

	return bytes, nil
}

func Unpack(data []byte) (interface{}, error) {
	return nil, nil
}
