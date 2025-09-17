package msgpack

import (
	"encoding/binary"
	"math"
	"reflect"
)

const Msgpack_Nil = 0xc0   //192
const Msgpack_False = 0xc2 //194
const Msgpack_True = 0xc3  //195

const Msgpack_Bin_8 = 0xc4  //196
const Msgpack_Bin_16 = 0xc5 //197
const Msgpack_Bin_32 = 0xc6 //198

const Msgpack_Uint_8 = 0xcc  //204
const Msgpack_Uint_16 = 0xcd //205
const Msgpack_Uint_32 = 0xce //206
const Msgpack_Uint_64 = 0xcf //207

const Msgpack_Int_8 = 0xd0  //208
const Msgpack_Int_16 = 0xd1 //209
const Msgpack_Int_32 = 0xd2 //210
const Msgpack_Int_64 = 0xd3 //211

const Msgpack_Float_32 = 0xca //202
const Msgpack_Float_64 = 0xcb //203

const Msgpack_String_8 = 0xd9  //217
const Msgpack_String_16 = 0xda //218
const Msgpack_String_32 = 0xdb //219

const Msgpack_Array_16 = 0xdc //220
const Msgpack_Array_32 = 0xdd //221

const Msgpack_Map_16 = 0xde //222
const Msgpack_Map_32 = 0xdf //223

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
	Arr      []byte
	Mapp     map[any]any
	Typed    map[float32]int16
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
		typeByte: Msgpack_Int_8,
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
		t = Msgpack_String_8
	case len(value) < string_16:
		size = encodeSize(len(value))
		t = Msgpack_String_16
	case len(value) < string_32:
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

func encodeArray(arr any) Array {
	array_value := reflect.ValueOf(arr)
	array_size := array_value.Len()

	array_8 := int(math.Pow(2, 8) - 1)
	array_16 := int(math.Pow(2, 16) - 1)
	array_32 := int(math.Pow(2, 32) - 1)
	bytes := make([]byte, 0)
	var size []byte
	var t byte
	bin_flag := false

	if array_size != 0 {
		element := array_value.Index(0).Interface()
		if element != nil {
			if reflect.TypeOf(element).Kind() == reflect.Uint8 {
				bin_flag = true
			}
		}
	}
	if array_size <= array_16 && !bin_flag {
		size = encodeSize(array_size)
		t = Msgpack_Array_16
	} else if array_size <= array_32 && !bin_flag {
		size = encodeSize(array_size)
		t = Msgpack_Array_32
	} else if array_size <= array_8 && bin_flag {
		size = make([]byte, 0)
		size = append(size, byte(array_size))
		t = Msgpack_Bin_8
	} else if array_size <= array_16 && bin_flag {
		size = encodeSize(array_size)
		t = Msgpack_Bin_16
	} else if array_size <= array_32 && bin_flag {
		size = encodeSize(array_size)
		t = Msgpack_Bin_32
	}

	for i := 0; i < array_size; i++ {
		element := array_value.Index(i).Interface()
		if element == nil {
			bytes = append(bytes, Msgpack_Nil)
			continue
		}
		objectBytes := Serialize(element)
		bytes = append(bytes, objectBytes...)
	}

	return Array{
		typeByte: t,
		value:    bytes,
		size:     size,
	}
}

func encodeMap(m any) Map {
	map_value := reflect.ValueOf(m)
	map_size := map_value.Len()
	map_16 := int(math.Pow(2, 16) - 1)
	map_32 := int(math.Pow(2, 32) - 1)

	bytes := make([]byte, 0)
	var size []byte
	var t byte

	if map_size <= map_16 {
		size = encodeSize(map_size)
		t = Msgpack_Map_16
	} else if map_size <= map_32 {
		size = encodeSize(map_size)
		t = Msgpack_Map_32
	}

	iter := map_value.MapRange()
	for iter.Next() {
		key := iter.Key().Interface()
		val := iter.Value().Interface()

		objectBytes := Serialize(key)
		bytes = append(bytes, objectBytes...)

		objectBytes = Serialize(val)
		bytes = append(bytes, objectBytes...)
	}

	return Map{
		typeByte: t,
		value:    bytes,
		size:     size,
	}
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

func decodeArray16(value []byte) ([]any, int) {
	size := decodeSize(value[:2])
	data := value[2:]
	array := make([]any, size)
	overallOffset := 0
	for i := 0; i < size; i++ {
		element, offset := deserialize_helper(data)
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
		element, offset := deserialize_helper(data)
		array[i] = element
		data = data[offset:]
		overallOffset += offset
	}
	return array, overallOffset + 5
}

func decodeMap16(value []byte) (any, int) {
	size := decodeSize(value[:2])
	data := value[2:]
	m := make(map[any]any)
	overallOffset := 0

	for i := 0; i < size; i++ {
		key, offset := deserialize_helper(data)
		data = data[offset:]
		overallOffset += offset
		val, offset := deserialize_helper(data)
		data = data[offset:]
		m[key] = val
		overallOffset += offset
	}
	return m, overallOffset + 3
}

func decodeMap32(value []byte) (any, int) {
	size := decodeSize(value[:4])
	data := value[4:]
	m := make(map[any]any)
	overallOffset := 0

	for i := 0; i < size; i++ {
		key, offset := deserialize_helper(data)
		data = data[offset:]
		overallOffset += offset
		val, offset := deserialize_helper(data)
		data = data[offset:]
		m[key] = val
		overallOffset += offset
	}
	return m, overallOffset + 5
}

func decodeBin8(value []byte) ([]byte, int) {
	size := int(value[0])
	data := value[1:]
	array := make([]byte, size)
	overallOffset := 0
	for i := 0; i < size; i++ {
		element, offset := deserialize_helper(data)
		array[i] = element.(byte)
		data = data[offset:]
		overallOffset += offset
	}

	return array, overallOffset + 2
}

func decodeBin16(value []byte) ([]byte, int) {
	size := decodeSize(value[:2])
	data := value[2:]
	array := make([]byte, size)
	overallOffset := 0
	for i := 0; i < size; i++ {
		element, offset := deserialize_helper(data)
		array[i] = element.(byte)
		data = data[offset:]
		overallOffset += offset
	}

	return array, overallOffset + 3
}

func decodeBin32(value []byte) ([]byte, int) {
	size := decodeSize(value[:4])
	data := value[4:]
	array := make([]byte, size)
	overallOffset := 0
	for i := 0; i < size; i++ {
		element, offset := deserialize_helper(data)
		array[i] = element.(byte)
		data = data[offset:]
		overallOffset += offset
	}

	return array, overallOffset + 5
}

func deserialize_helper(value []byte) (any, int) {
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
	case Msgpack_Bin_8:
		return decodeBin8(value[1:])
	case Msgpack_Bin_16:
		return decodeBin16(value[1:])
	case Msgpack_Bin_32:
		return decodeBin32(value[1:])
	case Msgpack_Nil:
		return nil, 1
	}

	return nil, 0
}

// a function that serializes an element into a byte array
func Serialize(elementVal any) []byte {
	elementValue := reflect.ValueOf(elementVal)
	objectBytes := make([]byte, 0)
	elementType := reflect.TypeOf(elementVal).Kind()
	switch elementType {
	case reflect.Bool:
		objectBytes = append(objectBytes, encodeBool(elementValue.Bool()))

	case reflect.Uint8:
		encodedUint := encodeUint(uint8(elementValue.Uint()))
		objectBytes = append(objectBytes, encodedUint.typeByte, encodedUint.value)

	case reflect.Uint16:
		encodedUint := encodeUint16(uint16(elementValue.Uint()))
		objectBytes = append(objectBytes, encodedUint.typeByte)
		objectBytes = append(objectBytes, encodedUint.value...)

	case reflect.Uint32:
		encodedUint := encodeUint32(uint32(elementValue.Uint()))
		objectBytes = append(objectBytes, encodedUint.typeByte)
		objectBytes = append(objectBytes, encodedUint.value...)

	case reflect.Uint64:
		encodedUint := encodeUint64(uint64(elementValue.Uint()))
		objectBytes = append(objectBytes, encodedUint.typeByte)
		objectBytes = append(objectBytes, encodedUint.value...)

	case reflect.Int8:
		encodedInt := encodeInt8(int8(elementValue.Int()))
		objectBytes = append(objectBytes, encodedInt.typeByte, encodedInt.value)

	case reflect.Int16:
		encodedInt := encodeInt16(int16(elementValue.Int()))
		objectBytes = append(objectBytes, encodedInt.typeByte)
		objectBytes = append(objectBytes, encodedInt.value...)

	case reflect.Int32:
		encodedInt := encodeInt32(int32(elementValue.Int()))
		objectBytes = append(objectBytes, encodedInt.typeByte)
		objectBytes = append(objectBytes, encodedInt.value...)

	case reflect.Int64:
		encodedInt := encodeInt64(int64(elementValue.Int()))
		objectBytes = append(objectBytes, encodedInt.typeByte)
		objectBytes = append(objectBytes, encodedInt.value...)

	case reflect.Int:
		encodedInt := encodeInt(int(elementValue.Int()))
		objectBytes = append(objectBytes, encodedInt.typeByte)
		objectBytes = append(objectBytes, encodedInt.value...)

	case reflect.Float32:
		encodedFloat := encodeFloat32(float32(elementValue.Float()))
		objectBytes = append(objectBytes, encodedFloat.typeByte)
		objectBytes = append(objectBytes, encodedFloat.value...)

	case reflect.Float64:
		encodedFloat := encodeFloat64(elementValue.Float())
		objectBytes = append(objectBytes, encodedFloat.typeByte)
		objectBytes = append(objectBytes, encodedFloat.value...)

	case reflect.String:
		encodedString := encodeString(elementValue.String())
		objectBytes = append(objectBytes, encodedString.typeByte)
		objectBytes = append(objectBytes, encodedString.size...)
		objectBytes = append(objectBytes, encodedString.value...)

	case reflect.Slice:
		encodedArray := encodeArray(elementValue.Interface())
		objectBytes = append(objectBytes, encodedArray.typeByte)
		objectBytes = append(objectBytes, encodedArray.size...)
		objectBytes = append(objectBytes, encodedArray.value...)

	case reflect.Map:
		encodedMap := encodeMap(elementValue.Interface())
		objectBytes = append(objectBytes, encodedMap.typeByte)
		objectBytes = append(objectBytes, encodedMap.size...)
		objectBytes = append(objectBytes, encodedMap.value...)

	default:
		objectBytes = append(objectBytes, Msgpack_Nil)
	}

	return objectBytes

}

// a function that deserializes a byte array into an element
func Deserialize(value []byte) any {
	res, _ := deserialize_helper(value)
	return res
}

// a function that packs an object into a byte array
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
		memberName := field.Name
		memberValue := v.Field(i).Interface()

		objectBytes := Serialize(memberValue)

		nameBytes := Serialize(memberName)

		bytes = append(bytes, nameBytes...)
		bytes = append(bytes, objectBytes...)
	}
	return bytes, nil
}

func convertMap(m map[any]any, key_type reflect.Type, value_type reflect.Type) any {
	new_map := reflect.MakeMap(reflect.MapOf(key_type, value_type))

	for key, val := range m {
		new_map.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	}
	return new_map.Interface()
}

// a function that unpacks a byte array into an object
func Unpack(bytes []byte, obj interface{}) (any, error) {
	unpacked := Deserialize(bytes)

	v := reflect.ValueOf(obj).Elem()

	for key, val := range unpacked.(map[any]any) {
		field := v.FieldByName(key.(string))
		packed_map_type := reflect.TypeOf(val)
		object_map_type := reflect.TypeOf(field.Interface())

		if field.Kind() == reflect.Map && object_map_type != packed_map_type {
			object_map_key_type := object_map_type.Key()
			object_map_val_type := object_map_type.Elem()
			new_map := convertMap(val.(map[any]any), object_map_key_type, object_map_val_type)
			field.Set(reflect.ValueOf(new_map))
			continue
		}

		field.Set(reflect.ValueOf(val))

	}

	return unpacked, nil
}
