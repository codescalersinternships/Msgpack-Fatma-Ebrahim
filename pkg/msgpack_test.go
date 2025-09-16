package msgpack

import (
	"strings"
	"testing"
)

type user struct {
	Name string
	Id   string
}

func TestPack(t *testing.T) {
	t.Run("test pack with string8", func(t *testing.T) {
		user := user{
			Name: "Fatma",
			Id:   "id_1",
		}
		packed, err := Pack(user)
		expected := []byte{222, 0, 2,
			217, 4, 78, 97, 109, 101, 217, 5, 70, 97, 116, 109, 97,
			217, 2, 73, 100, 217, 4, 105, 100, 95, 49}

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if string(packed) != string(expected) {
			t.Errorf("Expected %v, got %v", expected, packed)
		}

	})

	t.Run("test pack with string16", func(t *testing.T) {
		user := user{
			Id:   strings.Repeat("x", 260),
		}
		packed, err := Pack(user)
		expected := []byte{222, 0, 2,
			217, 2, 73, 100, 217, 4, 105, 100, 95, 49}

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if string(packed) != string(expected) {
			t.Errorf("Expected %v, got %v", expected, packed)
		}

	})
}
