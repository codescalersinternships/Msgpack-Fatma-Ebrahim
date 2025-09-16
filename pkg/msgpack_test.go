package msgpack

import (
	"reflect"
	"strings"
	"testing"
)

type User struct {
	Name string
	Id   string
}

func TestPackString(t *testing.T) {
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
})

}

