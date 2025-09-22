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

func TestString(t *testing.T) {
	t.Run("test serialize and deserialize with string8", func(t *testing.T) {
		str := strings.Repeat("x", 200)

		data_bytes := Serialize(str)

		str_type := data_bytes[0]
		if str_type != 217 {
			t.Errorf("Expected type of string to be 217 (str8), got %d", str_type)
		}

		str_size := int(data_bytes[1])
		if str_size != 200 {
			t.Errorf("Expected size of string to be 200, got %d", str_size)
		}

		str_value := Deserialize(data_bytes)
		if str_value != strings.Repeat("x", 200) {
			t.Errorf("Expected value of string to be %s, got %s", strings.Repeat("x", 200), str_value)
		}
	})

	t.Run("test serialize and deserialize with string16", func(t *testing.T) {
		str := strings.Repeat("x", 260)
		data_bytes := Serialize(str)

		str_type := data_bytes[0]
		if str_type != 218 {
			t.Errorf("Expected type of string to be 218, got %d", str_type)
		}

		str_size := decodeSize(data_bytes[1:3])
		if str_size != 260 {
			t.Errorf("Expected size of string to be 260, got %d", str_size)
		}

		str_value := Deserialize(data_bytes)
		if str_value != strings.Repeat("x", 260) {
			t.Errorf("Expected value of string to be %s, got %s", strings.Repeat("x", 260), str_value)
		}

	})

	t.Run("test serialize  and deserialize with string32", func(t *testing.T) {
		str := strings.Repeat("x", 70000)
		data_bytes := Serialize(str)

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

		str_value := Deserialize(data_bytes)
		if str_value != strings.Repeat("x", 70000) {
			t.Errorf("Expected value of string to be %s, got %s", strings.Repeat("x", 70000), str_value)
		}
	})
}

func TestFloat(t *testing.T) {
	t.Run("test serialize and deserialize with float32", func(t *testing.T) {
		fnum32 := float32(3.14)
		data_bytes := Serialize(fnum32)

		float_type := data_bytes[0]
		if float_type != 202 {
			t.Errorf("Expected type of float to be 203 (float32), got %d", float_type)
		}

		float_value := Deserialize(data_bytes)
		if reflect.TypeOf(float_value).Kind() != reflect.Float32 {
			t.Errorf("Expected type of float to be float32, got %v", reflect.TypeOf(float_value).Kind())
		}

		if float_value != fnum32 {
			t.Errorf("Expected value of float to be %f, got %f", fnum32, float_value)
		}

	})

	t.Run("test serialize and deserialize with float64", func(t *testing.T) {
		fnum64 := float64(3.14)
		data_bytes := Serialize(fnum64)

		float_type := data_bytes[0]
		if float_type != 203 {
			t.Errorf("Expected type of float to be 203 (float64), got %d", float_type)
		}

		float_value := Deserialize(data_bytes)
		if reflect.TypeOf(float_value).Kind() != reflect.Float64 {
			t.Errorf("Expected type of float to be float64, got %v", reflect.TypeOf(float_value).Kind())
		}

		if float_value != fnum64 {
			t.Errorf("Expected value of float to be %f, got %f", fnum64, float_value)
		}
	})
}

func TestInt(t *testing.T) {
	t.Run("test serialize and deserialize with int8", func(t *testing.T) {
		snum8 := int8(-100)
		data_bytes := Serialize(snum8)

		int_type := data_bytes[0]
		if int_type != 208 {
			t.Errorf("Expected type of int to be 208 (int8), got %d", int_type)
		}

		int_value := Deserialize(data_bytes)
		if reflect.TypeOf(int_value).Kind() != reflect.Int8 {
			t.Errorf("Expected type of int to be int8, got %v", reflect.TypeOf(int_value).Kind())
		}
		if int_value != snum8 {
			t.Errorf("Expected value of int to be %d, got %d", snum8, int_value)
		}
	})

	t.Run("test serialize and deserialize with int16", func(t *testing.T) {
		snum16 := int16(300)
		data_bytes := Serialize(snum16)

		int_type := data_bytes[0]
		if int_type != 209 {
			t.Errorf("Expected type of int to be 209 (int16), got %d", int_type)
		}

		int_value := Deserialize(data_bytes)
		if reflect.TypeOf(int_value).Kind() != reflect.Int16 {
			t.Errorf("Expected type of int to be int16, got %v", reflect.TypeOf(int_value).Kind())
		}

		if int_value != snum16 {
			t.Errorf("Expected value of int to be %d, got %d", snum16, int_value)
		}
	})

	t.Run("test serialize and deserialize with int32", func(t *testing.T) {
		snum32 := int32(70000)
		data_bytes := Serialize(snum32)
		int_type := data_bytes[0]
		if int_type != 210 {
			t.Errorf("Expected type of int to be 210 (int32), got %d", int_type)
		}

		int_value := Deserialize(data_bytes)
		if reflect.TypeOf(int_value).Kind() != reflect.Int32 {
			t.Errorf("Expected type of int to be int32, got %v", reflect.TypeOf(int_value).Kind())
		}

		if int_value != snum32 {
			t.Errorf("Expected value of int to be %d, got %d", snum32, int_value)
		}
	})

	t.Run("test serialize and deserialize with int64", func(t *testing.T) {
		snum64 := int64(100)
		data_bytes := Serialize(snum64)
		int_type := data_bytes[0]
		if int_type != 211 {
			t.Errorf("Expected type of int to be 211 (int64), got %d", int_type)
		}

		int_value := Deserialize(data_bytes)
		if reflect.TypeOf(int_value).Kind() != reflect.Int64 {
			t.Errorf("Expected type of int to be int64, got %v", reflect.TypeOf(int_value).Kind())
		}

		if int_value != snum64 {
			t.Errorf("Expected value of int to be %d, got %d", snum64, int_value)
		}
	})

}

func TestUint(t *testing.T) {

	t.Run("test serialize and deserialize with uint8", func(t *testing.T) {
		unum8 := uint8(100)
		data_bytes := Serialize(unum8)

		uint_type := data_bytes[0]
		if uint_type != 204 {
			t.Errorf("Expected type of uint to be 204 (uint8), got %d", uint_type)
		}

		uint_value := Deserialize(data_bytes)
		if reflect.TypeOf(uint_value).Kind() != reflect.Uint8 {
			t.Errorf("Expected type of uint to be uint8, got %v", reflect.TypeOf(uint_value).Kind())
		}
	})

	t.Run("test serialize and deserialize with uint16", func(t *testing.T) {
		unum16 := uint16(300)
		data_bytes := Serialize(unum16)

		uint_type := data_bytes[0]
		if uint_type != 205 {
			t.Errorf("Expected type of uint to be 205 (uint16), got %d", uint_type)
		}

		uint_value := Deserialize(data_bytes)
		if reflect.TypeOf(uint_value).Kind() != reflect.Uint16 {
			t.Errorf("Expected type of uint to be uint16, got %v", reflect.TypeOf(uint_value).Kind())
		}
	})

	t.Run("test serialize and deserialize with uint32", func(t *testing.T) {
		unum32 := uint32(70000)
		data_bytes := Serialize(unum32)
		uint_type := data_bytes[0]
		if uint_type != 206 {
			t.Errorf("Expected type of uint to be 206 (uint32), got %d", uint_type)
		}

		uint_value := Deserialize(data_bytes)
		if reflect.TypeOf(uint_value).Kind() != reflect.Uint32 {
			t.Errorf("Expected type of uint to be uint32, got %v", reflect.TypeOf(uint_value).Kind())
		}
	})

	t.Run("test serialize and deserialize with uint64", func(t *testing.T) {
		unum64 := uint64(100)
		data_bytes := Serialize(unum64)
		uint_type := data_bytes[0]
		if uint_type != 207 {
			t.Errorf("Expected type of uint to be 207 (uint64), got %d", uint_type)
		}

		uint_value := Deserialize(data_bytes)
		if reflect.TypeOf(uint_value).Kind() != reflect.Uint64 {
			t.Errorf("Expected type of uint to be uint64, got %v", reflect.TypeOf(uint_value).Kind())
		}
	})
}

func TestArray(t *testing.T) {
	t.Run("test serialize and deserialize with array16", func(t *testing.T) {
		arr := []any{1.1, 2.2, 3.3}
		data_bytes := Serialize(arr)
		array_type := data_bytes[0]

		if array_type != 220 {
			t.Errorf("Expected type of array to be 220 (array16), got %d", array_type)
		}

		array_size := decodeSize(data_bytes[1:3])
		if array_size != 3 {
			t.Errorf("Expected size of array to be 3, got %d", array_size)
		}

		array_value := Deserialize(data_bytes)
		if reflect.TypeOf(array_value).Kind() != reflect.Slice {
			t.Errorf("Expected type of array to be slice, got %v", reflect.TypeOf(array_value).Kind())
		}
		if !reflect.DeepEqual(array_value, arr) {
			t.Errorf("Expected array to be %v, got %v", arr, array_value)
		}

	})

	t.Run("test serialize and deserialize with array32", func(t *testing.T) {
		arr := make([]any, 70000)
		for i := range arr {
			arr[i] = true
		}

		data_bytes := Serialize(arr)
		array_type := data_bytes[0]

		if array_type != 221 {
			t.Errorf("Expected type of array to be 221 (array32), got %d", array_type)
		}

		array_size := decodeSize(data_bytes[1:5])
		if array_size != 70000 {
			t.Errorf("Expected size of array to be 70000, got %d", array_size)
		}

		array_value := Deserialize(data_bytes)
		if reflect.TypeOf(array_value).Kind() != reflect.Slice {
			t.Errorf("Expected type of array to be slice, got %v", reflect.TypeOf(array_value).Kind())
		}
		if !reflect.DeepEqual(array_value, arr) {
			t.Errorf("Expected array to be %v, got %v", arr[:5], array_value.([]any)[:5]) // print first few elements
		}
	})

	t.Run("test serialize and deserialize with empty array", func(t *testing.T) {
		arr := make([]any, 3)
		data_bytes := Serialize(arr)
		array_type := data_bytes[0]

		if array_type != 220 {
			t.Errorf("Expected type of array to be 220 (array16), got %d", array_type)
		}

		array_size := decodeSize(data_bytes[1:3])
		if array_size != 3 {
			t.Errorf("Expected size of array to be 3, got %d", array_size)
		}

		array_value := Deserialize(data_bytes)
		if reflect.TypeOf(array_value).Kind() != reflect.Slice {
			t.Errorf("Expected type of array to be slice, got %v", reflect.TypeOf(array_value).Kind())
		}

		array := array_value.([]any)
		if array[0] != nil {
			t.Errorf("Expected array element to be nil, got %v", array[0])
		}

	})

}

func TestBinary(t *testing.T) {
	t.Run("test serialize and deserialize with binary8", func(t *testing.T) {
		arr := []byte{1, 2, 3}
		data_bytes := Serialize(arr)
		bin_type := data_bytes[0]

		if bin_type != 196 {
			t.Errorf("Expected type of binary to be 196 (binary8), got %d", bin_type)
		}

		bin_size := int(data_bytes[1])
		if bin_size != 3 {
			t.Errorf("Expected size of binary to be 3, got %d", bin_size)
		}

		bin_value := Deserialize(data_bytes)
		if reflect.TypeOf(bin_value).Kind() != reflect.Slice {
			t.Errorf("Expected type of binary to be slice, got %v", reflect.TypeOf(bin_value).Kind())
		}
		if !reflect.DeepEqual(bin_value, arr) {
			t.Errorf("Expected binary to be %v, got %v", arr, bin_value)
		}

	})

	t.Run("test serialize and deserialize with binary16", func(t *testing.T) {
		arr := make([]byte, 500)
		data_bytes := Serialize(arr)
		bin_type := data_bytes[0]

		if bin_type != 197 {
			t.Errorf("Expected type of binary to be 197 (binary16), got %d", bin_type)
		}

		bin_size := decodeSize(data_bytes[1:3])
		if bin_size != 500 {
			t.Errorf("Expected size of binary to be 500, got %d", bin_size)
		}

		bin_value := Deserialize(data_bytes)
		if reflect.TypeOf(bin_value).Kind() != reflect.Slice {
			t.Errorf("Expected type of binary to be slice, got %v", reflect.TypeOf(bin_value).Kind())
		}
		if !reflect.DeepEqual(bin_value, arr) {
			t.Errorf("Expected binary to be %v, got %v", arr, bin_value)
		}

	})

	t.Run("test serialize and deserialize with binary32", func(t *testing.T) {
		arr := make([]byte, 70000)
		data_bytes := Serialize(arr)
		bin_type := data_bytes[0]

		if bin_type != 198 {
			t.Errorf("Expected type of binary to be 198 (binary32), got %d", bin_type)
		}

		bin_size := decodeSize(data_bytes[1:5])
		if bin_size != 70000 {
			t.Errorf("Expected size of binary to be 70000, got %d", bin_size)
		}

		bin_value := Deserialize(data_bytes)
		if reflect.TypeOf(bin_value).Kind() != reflect.Slice {
			t.Errorf("Expected type of binary to be slice, got %v", reflect.TypeOf(bin_value).Kind())
		}
		if !reflect.DeepEqual(bin_value, arr) {
			t.Errorf("Expected binary to be %v, got %v", arr, bin_value)
		}

	})

	t.Run("test serialize and deserialize with empty binary", func(t *testing.T) {
		arr := make([]byte, 1)
		data_bytes := Serialize(arr)
		bin_type := data_bytes[0]

		if bin_type != 196 {
			t.Errorf("Expected type of binary to be 196 (binary8), got %d", bin_type)
		}

		bin_size := int(data_bytes[1])
		if bin_size != 1 {
			t.Errorf("Expected size of binary to be 1, got %d", bin_size)
		}

		bin_value := Deserialize(data_bytes)
		if reflect.TypeOf(bin_value).Kind() != reflect.Slice {
			t.Errorf("Expected type of binary to be slice, got %v", reflect.TypeOf(bin_value).Kind())
		}

		bin := bin_value.([]byte)
		if bin[0] != 0 {
			t.Errorf("Expected binary element to be 0, got %v", bin[0])
		}

	})
}

func TestMap(t *testing.T) {
	t.Run("test serialize and deserialize with map16", func(t *testing.T) {
		m := make(map[any]any)
		m["Lang"] = "Go"
		m["Ver"] = "1.19"
		data_bytes := Serialize(m)

		map_type := data_bytes[0]
		if map_type != 222 {
			t.Errorf("Expected type of map to be 222 (map16), got %d", map_type)
		}

		map_size := decodeSize(data_bytes[1:3])
		if map_size != 2 {
			t.Errorf("Expected size of map to be 2, got %d", map_size)
		}

		map_value := Deserialize(data_bytes)
		if reflect.TypeOf(map_value).Kind() != reflect.Map {
			t.Errorf("Expected type of map to be map, got %v", reflect.TypeOf(map_value).Kind())
		}

		if !reflect.DeepEqual(map_value, m) {
			t.Errorf("Expected map to be %v, got %v", m, map_value)
		}
	})

	t.Run("test serialize and deserialize with map32", func(t *testing.T) {
		m := make(map[any]any, 70000)
		for i := 0; i < 70000; i++ {
			m[fmt.Sprintf("k%d", i)] = fmt.Sprintf("v%d", i)
		}
		data_bytes := Serialize(m)

		map_type := data_bytes[0]
		if map_type != 223 {
			t.Errorf("Expected type of map to be 223 (map32), got %d", map_type)
		}

		map_size := decodeSize(data_bytes[1:5])
		if map_size != 70000 {
			t.Errorf("Expected size of map to be 70000, got %d", map_size)
		}

		map_value := Deserialize(data_bytes)
		if reflect.TypeOf(map_value).Kind() != reflect.Map {
			t.Errorf("Expected type of map to be map, got %v", reflect.TypeOf(map_value).Kind())
		}

		if !reflect.DeepEqual(map_value, m) {
			t.Errorf("Expected map to be %v, got %v", m, map_value)
		}
	})

	t.Run("test serialize and deserialize with empty map", func(t *testing.T) {
		m := make(map[any]any)
		data_bytes := Serialize(m)

		map_type := data_bytes[0]
		if map_type != 222 {
			t.Errorf("Expected type of map to be 222 (map16), got %d", map_type)
		}

		map_size := decodeSize(data_bytes[1:3])
		if map_size != 0 {
			t.Errorf("Expected size of map to be 0, got %d", map_size)
		}

		map_value := Deserialize(data_bytes)
		if reflect.TypeOf(map_value).Kind() != reflect.Map {
			t.Errorf("Expected type of map to be map, got %v", reflect.TypeOf(map_value).Kind())
		}

		if !reflect.DeepEqual(map_value, m) {
			t.Errorf("Expected map to be %v, got %v", m, map_value)
		}
	})
}

func TestAll(t *testing.T) {
	t.Run("test pack and unpack with simple struct", func(t *testing.T) {
		user_to_pack := User{
			Name: "Fatma",
			Id:   "id_1",
		}
		packed, err := Pack(user_to_pack)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
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

	t.Run("test pack and unpack with complex struct", func(t *testing.T) {
		arr := make([]float32, 3)
		arr[0] = 1.1
		arr[1] = 2.2
		arr[2] = 3.3

		m := make(map[any]any)
		m["hello"] = 10.0
		m["world"] = -20.0

		obj_to_pack := Object{
			IsObject: true,
			Flag:     false,
			Unum:     0,
			Snum:     -1000,
			Fnum:     3.14,
			Str:      "Fatma",
			Arr:      arr,
			Mapp:     m,
		}

		packed, err := Pack(obj_to_pack)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		var obj_unpacked Object
		_, err = Unpack(packed, &obj_unpacked)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		packed_unpacked, err := Pack(obj_unpacked)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(packed, packed_unpacked) {
			t.Errorf("Expected %v, got %v", packed, packed_unpacked)
		}

	})

	type ComplexUser struct {
		ID       uint32
		Name     string
		Age      int8
		Balance  float64
		IsActive bool
		Tags     []any
		Metadata map[string]any
	}

	t.Run("test pack and unpack with more complex struct", func(t *testing.T) {
		user_to_pack := ComplexUser{
			ID:       123456,
			Name:     "Fatma Ebrahim",
			Age:      27,
			Balance:  1050.75,
			IsActive: true,
			Tags:     []any{"golang", "backend", "testing"},
			Metadata: map[string]any{
				"country": "Egypt",
				"city":    "Cairo",
			},
		}

		packed, err := Pack(user_to_pack)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		var user_unpacked ComplexUser
		_, err = Unpack(packed, &user_unpacked)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if !reflect.DeepEqual(user_to_pack, user_unpacked) {
			t.Errorf("Expected %+v, got %+v", user_to_pack, user_unpacked)
		}
	})
}
