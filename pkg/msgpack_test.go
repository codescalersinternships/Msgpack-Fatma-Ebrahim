package msgpack

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type User struct {
	Name string
	Id   string
}

func TestPackString(t *testing.T) {
	t.Run("test pack with string8", func(t *testing.T) {
		user := User{
			Id: strings.Repeat("x", 200),
		}
		packed, err := Pack(user)
		data_bytes := packed[15:]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(data_bytes) != 202 {
			t.Errorf("Expected size of data to be 202, got %d", len(data_bytes))
		}

		str_type := data_bytes[0]
		if str_type != 217 {
			t.Errorf("Expected type of string to be 217 (str8), got %d", str_type)
		}

		str_size := int(data_bytes[1])
		if str_size != 200 {
			t.Errorf("Expected size of string to be 200, got %d", str_size)
		}

		str_value, _ := deserializeElement(data_bytes)
		if str_value != strings.Repeat("x", 200) {
			t.Errorf("Expected value of string to be %s, got %s", strings.Repeat("x", 200), str_value)
		}
	})

	t.Run("test pack with string16", func(t *testing.T) {
		user := User{
			Id: strings.Repeat("x", 260),
		}
		packed, err := Pack(user)
		data_bytes := packed[15:]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(data_bytes) != 263 {
			t.Errorf("Expected size of data to be 263, got %d", len(data_bytes))
		}

		str_type := data_bytes[0]
		if str_type != 218 {
			t.Errorf("Expected type of string to be 218, got %d", str_type)
		}

		str_size := decodeSize(data_bytes[1:3])
		if str_size != 260 {
			t.Errorf("Expected size of string to be 260, got %d", str_size)
		}

		str_value, _ := deserializeElement(data_bytes)
		if str_value != strings.Repeat("x", 260) {
			t.Errorf("Expected value of string to be %s, got %s", strings.Repeat("x", 260), str_value)
		}

	})

	t.Run("test pack with string32", func(t *testing.T) {
		user := User{
			Id: strings.Repeat("x", 70000),
		}
		packed, err := Pack(user)
		data_bytes := packed[15:]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(data_bytes) != 70005 {
			t.Errorf("Expected size of data to be 70005, got %d", len(data_bytes))
		}

		str_type := data_bytes[0]
		if str_type != 219 {
			t.Errorf("Expected type of string to be 219 (str32), got %d", str_type)
		}

		str_size := decodeSize(data_bytes[1:5])
		if str_size != 70000 {
			t.Errorf("Expected size of string to be 70000, got %d", str_size)
		}

		str_value, _ := deserializeElement(data_bytes)
		if str_value != strings.Repeat("x", 70000) {
			t.Errorf("Expected value of string to be %s, got %s", strings.Repeat("x", 70000), str_value)
		}
	})
}

func TestPackFloat(t *testing.T) {
	type Object struct {
		Fnum32 float32
		Fnum64 float64
	}
	t.Run("test pack with float32", func(t *testing.T) {

		obj := Object{
			Fnum32: 3.14,
		}

		packed, err := Pack(obj)
		data_bytes := packed[3+8 : 16]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		float_type := data_bytes[0]
		if float_type != 202 {
			t.Errorf("Expected type of float to be 203 (float32), got %d", float_type)
		}

		float_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(float_value).Kind() != reflect.Float32 {
			t.Errorf("Expected type of float to be float32, got %v", reflect.TypeOf(float_value).Kind())
		}

	})

	t.Run("test pack with float64", func(t *testing.T) {
		obj := Object{
			Fnum64: 3.14,
		}

		packed, err := Pack(obj)
		data_bytes := packed[16+8:]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		float_type := data_bytes[0]
		if float_type != 203 {
			t.Errorf("Expected type of float to be 203 (float64), got %d", float_type)
		}

		float_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(float_value).Kind() != reflect.Float64 {
			t.Errorf("Expected type of float to be float64, got %v", reflect.TypeOf(float_value).Kind())
		}
	})
}

func TestPackInt(t *testing.T) {
	type Object struct {
		Snum8  int8
		Snum16 int16
		Snum32 int32
		Snum64 int64
	}
	t.Run("test pack with int8", func(t *testing.T) {

		obj := Object{
			Snum8: 3,
		}

		packed, err := Pack(obj)
		data_bytes := packed[3+2+5 : 10+2]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		int_type := data_bytes[0]
		if int_type != 208 {
			t.Errorf("Expected type of int to be 208 (int8), got %d", int_type)
		}

		int_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(int_value).Kind() != reflect.Int8 {
			t.Errorf("Expected type of int to be int8, got %v", reflect.TypeOf(int_value).Kind())
		}
	})

	t.Run("test pack with int16", func(t *testing.T) {
		obj := Object{
			Snum16: -300,
		}

		packed, err := Pack(obj)
		data_bytes := packed[12+2+6 : 20+3]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		int_type := data_bytes[0]
		if int_type != 209 {
			t.Errorf("Expected type of int to be 209 (int16), got %d", int_type)
		}

		int_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(int_value).Kind() != reflect.Int16 {
			t.Errorf("Expected type of int to be int16, got %v", reflect.TypeOf(int_value).Kind())
		}
	})

	t.Run("test pack with int32", func(t *testing.T) {
		obj := Object{
			Snum32: 70000,
		}

		packed, err := Pack(obj)
		data_bytes := packed[23+2+6 : 31+5]
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		int_type := data_bytes[0]
		if int_type != 210 {
			t.Errorf("Expected type of int to be 210 (int32), got %d", int_type)
		}

		int_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(int_value).Kind() != reflect.Int32 {
			t.Errorf("Expected type of int to be int32, got %v", reflect.TypeOf(int_value).Kind())
		}
	})

	t.Run("test pack with int64", func(t *testing.T) {
		obj := Object{
			Snum64: 5000000000,
		}

		packed, err := Pack(obj)
		data_bytes := packed[36+2+6:]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		int_type := data_bytes[0]
		if int_type != 211 {
			t.Errorf("Expected type of int to be 211 (int64), got %d", int_type)
		}

		int_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(int_value).Kind() != reflect.Int64 {
			t.Errorf("Expected type of int to be int64, got %v", reflect.TypeOf(int_value).Kind())
		}
	})

}

func TestPackUint(t *testing.T) {
	type Object struct {
		Unum8  uint8
		Unum16 uint16
		Unum32 uint32
		Unum64 uint64
	}

	t.Run("test pack with uint8", func(t *testing.T) {
		obj := Object{
			Unum8: 250,
		}

		packed, err := Pack(obj)
		data_bytes := packed[3+2+5 : 10+2]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		uint_type := data_bytes[0]
		if uint_type != 204 {
			t.Errorf("Expected type of uint to be 204 (uint8), got %d", uint_type)
		}

		uint_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(uint_value).Kind() != reflect.Uint8 {
			t.Errorf("Expected type of uint to be uint8, got %v", reflect.TypeOf(uint_value).Kind())
		}
	})

	t.Run("test pack with uint16", func(t *testing.T) {
		obj := Object{
			Unum16: 700,
		}

		packed, err := Pack(obj)
		data_bytes := packed[12+2+6 : 20+3]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		uint_type := data_bytes[0]
		if uint_type != 205 {
			t.Errorf("Expected type of uint to be 205 (uint16), got %d", uint_type)
		}

		uint_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(uint_value).Kind() != reflect.Uint16 {
			t.Errorf("Expected type of uint to be uint16, got %v", reflect.TypeOf(uint_value).Kind())
		}
	})

	t.Run("test pack with uint32", func(t *testing.T) {
		obj := Object{
			Unum32: 70000,
		}

		packed, err := Pack(obj)
		data_bytes := packed[23+2+6 : 31+5]

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		uint_type := data_bytes[0]
		if uint_type != 206 {
			t.Errorf("Expected type of uint to be 206 (uint32), got %d", uint_type)
		}

		uint_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(uint_value).Kind() != reflect.Uint32 {
			t.Errorf("Expected type of uint to be uint32, got %v", reflect.TypeOf(uint_value).Kind())
		}
	})

	t.Run("test pack with uint64", func(t *testing.T) {
		obj := Object{
			Unum64: 5000000000,
		}

		packed, err := Pack(obj)
		data_bytes := packed[36+2+6:]
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		uint_type := data_bytes[0]
		if uint_type != 207 {
			t.Errorf("Expected type of uint to be 207 (uint64), got %d", uint_type)
		}

		uint_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(uint_value).Kind() != reflect.Uint64 {
			t.Errorf("Expected type of uint to be uint64, got %v", reflect.TypeOf(uint_value).Kind())
		}
	})
}

func TestPackArray(t *testing.T) {
	type Object struct {
		Arr []any
	}
	t.Run("test pack with array16", func(t *testing.T) {
		arr := []any{1.1, 2.2, 3.3}
		obj := Object{
			Arr: arr,
		}
		packed, err := Pack(obj)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		data_bytes := packed[3+5:]

		array_type := data_bytes[0]

		if array_type != 220 {
			t.Errorf("Expected type of array to be 220 (array), got %d", array_type)
		}

		array_size := decodeSize(data_bytes[1:3])
		if array_size != 3 {
			t.Errorf("Expected size of array to be 3, got %d", array_size)
		}

		array_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(array_value).Kind() != reflect.Slice {
			t.Errorf("Expected type of array to be slice, got %v", reflect.TypeOf(array_value).Kind())
		}
		if !reflect.DeepEqual(array_value, arr) {
			t.Errorf("Expected array to be %v, got %v", arr, array_value)
		}

	})

	t.Run("test pack with array32", func(t *testing.T) {
		arr := make([]any, 70000)
		for i := range arr {
			arr[i] = true
		}
		obj := Object{
			Arr: arr,
		}
		packed, err := Pack(obj)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		data_bytes := packed[3+5:]

		array_type := data_bytes[0]

		if array_type != 221 {
			t.Errorf("Expected type of array to be 221 (array32), got %d", array_type)
		}

		array_size := decodeSize(data_bytes[1:5])
		if array_size != 70000 {
			t.Errorf("Expected size of array to be 70000, got %d", array_size)
		}

		array_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(array_value).Kind() != reflect.Slice {
			t.Errorf("Expected type of array to be slice, got %v", reflect.TypeOf(array_value).Kind())
		}
		if !reflect.DeepEqual(array_value, arr) {
			t.Errorf("Expected array to be %v, got %v", arr[:5], array_value.([]any)[:5]) // print first few elements
		}
	})

	t.Run("test pack with empty array", func(t *testing.T) {
		arr := make([]any, 3)
		obj := Object{
			Arr: arr,
		}
		packed, err := Pack(obj)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		data_bytes := packed[3+5:]

		array_type := data_bytes[0]

		if array_type != 220 {
			t.Errorf("Expected type of array to be 220 (array), got %d", array_type)
		}

		array_size := decodeSize(data_bytes[1:3])
		if array_size != 3 {
			t.Errorf("Expected size of array to be 3, got %d", array_size)
		}

		array_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(array_value).Kind() != reflect.Slice {
			t.Errorf("Expected type of array to be slice, got %v", reflect.TypeOf(array_value).Kind())
		}

		array := array_value.([]any)
		if array[0] != nil {
			t.Errorf("Expected array element to be nil, got %v", array[0])
		}

	})

}

func TestPackMap(t *testing.T) {
	type Object struct {
		Dict map[any]any
	}

	t.Run("test pack with map16", func(t *testing.T) {
		m := map[any]any{
			"Lang": "Go",
			"Ver":  "1.19",
		}
		obj := Object{
			Dict: m,
		}
		packed, err := Pack(obj)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		data_bytes := packed[3+6:]

		map_type := data_bytes[0]
		if map_type != 222 {
			t.Errorf("Expected type of map to be 222 (map16), got %d", map_type)
		}

		map_size := decodeSize(data_bytes[1:3])
		if map_size != 2 {
			t.Errorf("Expected size of map to be 2, got %d", map_size)
		}

		map_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(map_value).Kind() != reflect.Map {
			t.Errorf("Expected type of map to be map, got %v", reflect.TypeOf(map_value).Kind())
		}

		if !reflect.DeepEqual(map_value, m) {
			t.Errorf("Expected map to be %v, got %v", m, map_value)
		}
	})

	t.Run("test pack with map32", func(t *testing.T) {
		m := make(map[any]any, 70000)
		for i := 0; i < 70000; i++ {
			m[fmt.Sprintf("k%d", i)] = fmt.Sprintf("v%d", i)
		}
		obj := Object{
			Dict: m,
		}
		packed, err := Pack(obj)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		data_bytes := packed[3+6:]

		map_type := data_bytes[0]
		if map_type != 223 {
			t.Errorf("Expected type of map to be 223 (map32), got %d", map_type)
		}

		map_size := decodeSize(data_bytes[1:5])
		if map_size != 70000 {
			t.Errorf("Expected size of map to be 70000, got %d", map_size)
		}

		map_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(map_value).Kind() != reflect.Map {
			t.Errorf("Expected type of map to be map, got %v", reflect.TypeOf(map_value).Kind())
		}

		if !reflect.DeepEqual(map_value, m) {
			t.Errorf("Expected map to be %v, got %v", m, map_value)
		}
	})

	t.Run("test pack with empty map", func(t *testing.T) {
		m := make(map[any]any)
		obj := Object{
			Dict: m,
		}
		packed, err := Pack(obj)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		data_bytes := packed[3+6:]

		map_type := data_bytes[0]
		if map_type != 222 {
			t.Errorf("Expected type of map to be 222 (map16), got %d", map_type)
		}

		map_size := decodeSize(data_bytes[1:3])
		if map_size != 0 {
			t.Errorf("Expected size of map to be 0, got %d", map_size)
		}

		map_value, _ := deserializeElement(data_bytes)
		if reflect.TypeOf(map_value).Kind() != reflect.Map {
			t.Errorf("Expected type of map to be map, got %v", reflect.TypeOf(map_value).Kind())
		}

		if !reflect.DeepEqual(map_value, m) {
			t.Errorf("Expected map to be %v, got %v", m, map_value)
		}
	})
}

func TestAll(t *testing.T) {
	t.Run("test pack and unpack with string8", func(t *testing.T) {
		user_to_pack := User{
			Name: "Fatma",
			Id:   "id_1",
		}
		packed, err := Pack(user_to_pack)
		expected := []byte{222, 0, 2,
			217, 4, 78, 97, 109, 101, 217, 5, 70, 97, 116, 109, 97,
			217, 2, 73, 100, 217, 4, 105, 100, 95, 49}
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if string(packed) != string(expected) {
			t.Errorf("Expected %v, got %v", expected, packed)
		}

		var user_unpacked User
		_, err = Unpack(packed, &user_unpacked)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if !reflect.DeepEqual(user_to_pack, user_unpacked) {
			t.Errorf("Expected %v, got %v", user_to_pack, user_unpacked)
		}

	})
}
