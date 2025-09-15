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

const Msgpack_Array_16 = 0xdc
const Msgpack_Array_32 = 0xdd
const Msgpack_Map_16 = 0xde
const Msgpack_Map_32 = 0xdf

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

type Array struct {
	typeByte byte
	value    []byte
	size     []byte
}

type Map struct {
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
	Arr      []any
	Mapp     map[any]any
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
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, value)
	return Uint{
		typeByte: Msgpack_Uint_64,
		value:    bytes,
	}
}

func encodeInt(value int) Int {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(value))
	return Int{
		typeByte: Msgpack_Int_64,
		value:    bytes,
	}
}

func encodeInt8(value int8) Int_8 {
	return Int_8{
		typeByte: Msgpack_Uint_8,
		value:    byte(value),
	}
}

func encodeInt16(value int16) Int {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, uint16(value))
	return Int{
		typeByte: Msgpack_Int_16,
		value:    bytes,
	}
}

func encodeInt32(value int32) Int {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(value))
	return Int{
		typeByte: Msgpack_Int_32,
		value:    bytes,
	}
}

func encodeInt64(value int64) Int {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(value))
	return Int{
		typeByte: Msgpack_Int_64,
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
		// size = encodeSize(len(value))
		t = Msgpack_String_8
	case len(value) < string_16:
		// size = make([]byte, 2)
		// binary.BigEndian.PutUint16(size, uint16(len(value)))
		size = encodeSize(len(value))
		t = Msgpack_String_16
	case len(value) < string_32:
		// size = make([]byte, 4)
		// binary.BigEndian.PutUint32(size, uint32(len(value)))
		size = encodeSize(len(value))
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

func encodeArray(arr []any) Array {
	array_16 := int(math.Pow(2, 16) - 1)
	array_32 := int(math.Pow(2, 32) - 1)
	bytes := make([]byte, 0)
	var size []byte
	var t byte

	if len(arr) <= array_16 {
		// size = make([]byte, 2)
		// binary.BigEndian.PutUint16(size, uint16(len(arr)))
		size = encodeSize(len(arr))
		t = Msgpack_Array_16
	} else if len(arr) <= array_32 {
		// size = make([]byte, 4)
		// binary.BigEndian.PutUint32(size, uint32(len(arr)))
		size = encodeSize(len(arr))
		t = Msgpack_Array_32
	}

	for _, element := range arr {
		elementType := reflect.TypeOf(element).Kind()
		elementValue := reflect.ValueOf(element)
		objectBytes := make([]byte, 0)
		serializeElement(&objectBytes, elementType, elementValue)
		bytes = append(bytes, objectBytes...)
	}

	return Array{
		typeByte: t,
		value:    bytes,
		size:     size,
	}
}

func encodeMap(m map[any]any) Map {
	map_16 := int(math.Pow(2, 16) - 1)
	map_32 := int(math.Pow(2, 32) - 1)
	bytes := make([]byte, 0)
	var size []byte
	var t byte

	if len(m) <= map_16 {
		// size = make([]byte, 2)
		// binary.BigEndian.PutUint16(size, uint16(len(m)))
		size = encodeSize(len(m))
		t = Msgpack_Map_16
	} else if len(m) <= map_32 {
		// size = make([]byte, 4)
		// binary.BigEndian.PutUint32(size, uint32(len(m)))
		size = encodeSize(len(m))
		t = Msgpack_Map_32
	}

	for key, val := range m {
		keyType := reflect.TypeOf(key).Kind()
		keyValue := reflect.ValueOf(key)
		objectBytes := make([]byte, 0)
		serializeElement(&objectBytes, keyType, keyValue)
		bytes = append(bytes, objectBytes...)

		valueType := reflect.TypeOf(val).Kind()
		valueValue := reflect.ValueOf(val)
		objectBytes = make([]byte, 0)
		serializeElement(&objectBytes, valueType, valueValue)
		bytes = append(bytes, objectBytes...)
	}

	return Map{
		typeByte: t,
		value:    bytes,
		size:     size,
	}
}

func decodeBool(value byte) (bool, int) {
	if value == byte(Msgpack_True) {
		return true, 1
	} else {
		return false, 1
	}
}

func decodeUint8(value byte) (uint8, int) {
	return uint8(value), 2
}

func decodeUint16(value []byte) (uint16, int) {
	return binary.BigEndian.Uint16(value[:2]), 3
}

func decodeUint32(value []byte) (uint32, int) {
	return binary.BigEndian.Uint32(value[:4]), 5
}

func decodeUint64(value []byte) (uint64, int) {
	return binary.BigEndian.Uint64(value[:8]), 9
}

func decodeInt8(value byte) (int8, int) {
	return int8(value), 2
}

func decodeInt16(value []byte) (int16, int) {
	return int16(binary.BigEndian.Uint16(value[:2])), 3
}

func decodeInt32(value []byte) (int32, int) {
	return int32(binary.BigEndian.Uint32(value[:4])), 5
}

func decodeInt64(value []byte) (int64, int) {
	return int64(binary.BigEndian.Uint64(value[:8])), 9
}

func decodeFloat32(value []byte) (float32, int) {
	bin := binary.BigEndian.Uint32(value[:4])
	return math.Float32frombits(bin), 5
}

func decodeFloat64(value []byte) (float64, int) {
	bin := binary.BigEndian.Uint64(value[:8])
	return math.Float64frombits(bin), 9
}

func decodeString8(value []byte) (string, int) {
	size := int(value[0])
	return string(value[1 : size+1]), 1 + 1 + size
}

func decodeString16(value []byte) (string, int) {
	size := decodeSize(value[:2])
	return string(value[2 : size+2]), 1 + 2 + size
}

func decodeString32(value []byte) (string, int) {
	size := decodeSize(value[:4])
	return string(value[4 : size+4]), 1 + 4 + size
}

func decodeSize(value []byte) int {
	if len(value) == 2 {
		return int(binary.BigEndian.Uint16(value))
	} else if len(value) == 4 {
		return int(binary.BigEndian.Uint32(value))
	} else if len(value) == 8 {
		return int(binary.BigEndian.Uint64(value))
	}
	return 0
}

func encodeSize(size int) []byte {
	size_16 := int(math.Pow(2, 16) - 1)
	size_32 := int(math.Pow(2, 32) - 1)
	var sizeBytes []byte
	if size <= size_16 {
		sizeBytes = make([]byte, 2)
		binary.BigEndian.PutUint16(sizeBytes, uint16(size))
	} else if size <= size_32 {
		sizeBytes = make([]byte, 4)
		binary.BigEndian.PutUint32(sizeBytes, uint32(size))
	} else {
		sizeBytes = make([]byte, 8)
		binary.BigEndian.PutUint64(sizeBytes, uint64(size))
	}
	return sizeBytes
}

func decodeArray16(value []byte) ([]any, int) {
	size := decodeSize(value[:2])
	data := value[2:]
	array := make([]any, size)
	overallOffset := 0
	for i := 0; i < size; i++ {
		element, offset := deserializeElement(data)
		array[i] = element
		data = data[offset:]
		overallOffset += offset
	}

	return array, overallOffset + 3
}

func decodeArray32(value []byte) ([]any, int) {
	size := decodeSize(value[:4])
	data := value[4:]
	array := make([]any, size)
	overallOffset := 0

	for i := 0; i < size; i++ {
		element, offset := deserializeElement(data)
		array[i] = element
		data = data[offset:]
		overallOffset += offset
	}
	return array, overallOffset + 5
}

func decodeMap16(value []byte) (map[any]any, int) {
	size := decodeSize(value[:2])
	data := value[2:]
	m := make(map[any]any)
	overallOffset := 0

	for i := 0; i < size; i++ {
		key, offset := deserializeElement(data)
		data = data[offset:]
		value, offset := deserializeElement(data)
		data = data[offset:]
		m[key] = value
		overallOffset += offset
	}
	return m, overallOffset + 3
}

func decodeMap32(value []byte) (map[any]any, int) {
	size := decodeSize(value[:4])
	data := value[4:]
	m := make(map[any]any)
	overallOffset := 0

	for i := 0; i < size; i++ {
		key, offset := deserializeElement(data)
		data = data[offset:]
		value, offset := deserializeElement(data)
		data = data[offset:]
		m[key] = value
		overallOffset += offset
	}
	return m, overallOffset + 5
}

func deserializeElement(value []byte) (any, int) {
	elementType := value[0]
	switch elementType {
	case Msgpack_True:
		return decodeBool(value[0])
	case Msgpack_False:
		return decodeBool(value[0])
	case Msgpack_Uint_8:
		return decodeUint8(value[1])
	case Msgpack_Uint_16:
		return decodeUint16(value[1:])
	case Msgpack_Uint_32:
		return decodeUint32(value[1:])
	case Msgpack_Uint_64:
		return decodeUint64(value[1:])
	case Msgpack_Int_8:
		return decodeInt8(value[1])
	case Msgpack_Int_16:
		return decodeInt16(value[1:])
	case Msgpack_Int_32:
		return decodeInt32(value[1:])
	case Msgpack_Int_64:
		return decodeInt64(value[1:])
	case Msgpack_Float_32:
		return decodeFloat32(value[1:])
	case Msgpack_Float_64:
		return decodeFloat64(value[1:])
	case Msgpack_String_8:
		return decodeString8(value[1:])
	case Msgpack_String_16:
		return decodeString16(value[1:])
	case Msgpack_String_32:
		return decodeString32(value[1:])
	case Msgpack_Array_16:
		return decodeArray16(value[1:])
	case Msgpack_Array_32:
		return decodeArray32(value[1:])
	case Msgpack_Map_16:
		return decodeMap16(value[1:])
	case Msgpack_Map_32:
		return decodeMap32(value[1:])
	}

	return nil, 0
}

func serializeElement(objectBytes *[]byte, elementType reflect.Kind, elementValue reflect.Value) {
	switch elementType {
	case reflect.Bool:
		*objectBytes = append(*objectBytes, encodeBool(elementValue.Bool()))

	case reflect.Uint8:
		encodedUint := encodeUint(uint8(elementValue.Uint()))
		*objectBytes = append(*objectBytes, encodedUint.typeByte, encodedUint.value)

	case reflect.Uint16:
		encodedUint := encodeUint16(uint16(elementValue.Uint()))
		*objectBytes = append(*objectBytes, encodedUint.typeByte)
		*objectBytes = append(*objectBytes, encodedUint.value...)

	case reflect.Uint32:
		encodedUint := encodeUint32(uint32(elementValue.Uint()))
		*objectBytes = append(*objectBytes, encodedUint.typeByte)
		*objectBytes = append(*objectBytes, encodedUint.value...)

	case reflect.Uint64:
		encodedUint := encodeUint64(uint64(elementValue.Uint()))
		*objectBytes = append(*objectBytes, encodedUint.typeByte)
		*objectBytes = append(*objectBytes, encodedUint.value...)

	case reflect.Int8:
		encodedInt := encodeInt8(int8(elementValue.Int()))
		*objectBytes = append(*objectBytes, encodedInt.typeByte, encodedInt.value)

	case reflect.Int16:
		encodedInt := encodeInt16(int16(elementValue.Int()))
		*objectBytes = append(*objectBytes, encodedInt.typeByte)
		*objectBytes = append(*objectBytes, encodedInt.value...)

	case reflect.Int32:
		encodedInt := encodeInt32(int32(elementValue.Int()))
		*objectBytes = append(*objectBytes, encodedInt.typeByte)
		*objectBytes = append(*objectBytes, encodedInt.value...)

	case reflect.Int64:
		encodedInt := encodeInt64(int64(elementValue.Int()))
		*objectBytes = append(*objectBytes, encodedInt.typeByte)
		*objectBytes = append(*objectBytes, encodedInt.value...)

	case reflect.Int:
		encodedInt := encodeInt(int(elementValue.Int()))
		*objectBytes = append(*objectBytes, encodedInt.typeByte)
		*objectBytes = append(*objectBytes, encodedInt.value...)

	case reflect.Float32:
		encodedFloat := encodeFloat32(float32(elementValue.Float()))
		*objectBytes = append(*objectBytes, encodedFloat.typeByte)
		*objectBytes = append(*objectBytes, encodedFloat.value...)

	case reflect.Float64:
		encodedFloat := encodeFloat64(elementValue.Float())
		*objectBytes = append(*objectBytes, encodedFloat.typeByte)
		*objectBytes = append(*objectBytes, encodedFloat.value...)

	case reflect.String:
		encodedString := encodeString(elementValue.String())
		*objectBytes = append(*objectBytes, encodedString.typeByte)
		*objectBytes = append(*objectBytes, encodedString.size...)
		*objectBytes = append(*objectBytes, encodedString.value...)

	case reflect.Slice:
		encodedArray := encodeArray(elementValue.Interface().([]any))
		*objectBytes = append(*objectBytes, encodedArray.typeByte)
		*objectBytes = append(*objectBytes, encodedArray.size...)
		*objectBytes = append(*objectBytes, encodedArray.value...)

	case reflect.Map:
		encodedMap := encodeMap(elementValue.Interface().(map[any]any))
		*objectBytes = append(*objectBytes, encodedMap.typeByte)
		*objectBytes = append(*objectBytes, encodedMap.size...)
		*objectBytes = append(*objectBytes, encodedMap.value...)

	default:
		*objectBytes = append(*objectBytes, Msgpack_Nil)
	}

}

func Pack(obj interface{}) ([]byte, error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	bytes := make([]byte, 0)
	size := make([]byte, 0)
	var mapType byte
	if t.NumField() <= 16 {
		size = encodeSize(t.NumField())
		mapType = Msgpack_Map_16
	} else if t.NumField() <= 32 {
		size = encodeSize(t.NumField())
		mapType = Msgpack_Map_32
	}
	bytes = append(bytes, mapType)
	bytes = append(bytes, size...)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		memberType := field.Type.Kind()
		memberName := field.Name
		memberValue := v.Field(i)
		fmt.Printf("struct member %d: %s %s %v \n", i, memberType, memberName, memberValue)

		objectBytes := make([]byte, 0)
		serializeElement(&objectBytes, memberType, memberValue)
		fmt.Println("serialized:", objectBytes)

		data, offset := deserializeElement(objectBytes)
		fmt.Println("deserialized:", data, offset)

		nameBytes := make([]byte, 0)
		serializeElement(&nameBytes, reflect.String, reflect.ValueOf(memberName))
		fmt.Println("name bytes:", nameBytes, memberName)

		bytes = append(bytes, nameBytes...)
		bytes = append(bytes, objectBytes...)
	}
	return bytes, nil
}

func Unpack(bytes []byte) (any, error) {
	unpacked, _ := deserializeElement(bytes)
	return unpacked, nil
}
