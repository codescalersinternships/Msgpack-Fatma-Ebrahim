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
const Msgpack_Int_8 = 0xd0
const Msgpack_Float_32 = 0xca

const Msgpack_String_8 = 0xd9
const Msgpack_String_16 = 0xda
const Msgpack_String_32 = 0xdb

type Uint_8 struct {
	typeByte byte
	value    byte
}

type Int_8 struct {
	typeByte byte
	value    byte
}

type Float_32 struct {
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
	Snum     int8
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

func encodeInt(value int8) Int_8 {
	return Int_8{
		typeByte: Msgpack_Uint_8,
		value:    byte(value),
	}
}

func encodeFloat(value float32) Float_32 {
	bin := math.Float32bits(value)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bin)
	return Float_32{
		typeByte: Msgpack_Float_32,
		value:    bytes,
	}
}

func encodeString(value string) String {
	string_8 := int(math.Pow(2, 8) - 1)
	string_16 := int(math.Pow(2, 16) - 1)
	string_32 := int(math.Pow(2, 32) - 1)
	var size []byte
	numBytes := 8
	var t byte
	switch {
	case len(value) < string_8:
		numBytes = 8
		size = make([]byte, 0)
		size = append(size, byte(len(value)))
		t = Msgpack_String_8
	case len(value) < string_16:
		numBytes = 16
		size = make([]byte, 2)
		binary.BigEndian.PutUint16(size, uint16(len(value)))
		t = Msgpack_String_16
	case len(value) < string_32:
		numBytes = 32
		size = make([]byte, 4)
		binary.BigEndian.PutUint32(size, uint32(len(value)))
		t = Msgpack_String_32
	}

	bytes := make([]byte, int(math.Pow(2, float64(numBytes))-1))
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

func decodeUint(value byte) uint8 {
	return uint8(value)
}

func decodeInt(value byte) int8 {
	return int8(value)
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
			fmt.Println(decodeUint(encodedUint.value))
		case reflect.Int8:
			encodedInt := encodeInt(int8(memberValue.Int()))
			objectBytes = append(objectBytes, encodedInt.typeByte, encodedInt.value)
			fmt.Println(decodeInt(encodedInt.value))
		case reflect.Float32:
			encodedFloat := encodeFloat(float32(memberValue.Float()))
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
