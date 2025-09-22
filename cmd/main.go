package main

import (
	"fmt"
	msgpack "github.com/codescalersinternships/Msgpack-Fatma-Ebrahim/pkg"
)

func main() {
	arr := make([]float32, 3)
	arr[0] = 0
	arr[1] = 7.1
	arr[2] = 9.8

	m := make(map[any]any)
	m["hello"] = 10
	m["world"] = -20

	typed_map := make(map[float32]int16)
	typed_map[1.1] = 1
	typed_map[2.2] = 2

	object := msgpack.Object{
		IsObject: true,
		Flag:     false,
		Unum:     0,
		Snum:     -1000,
		Fnum:     3.14,
		Str:      "Fatma",
		Arr:      arr,
		Mapp:     m,
		Typed:    typed_map,
	}

	packed, err := msgpack.Pack(object)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Packed: %+v\n", packed)

	var obj msgpack.Object
	unpacked, err := msgpack.Unpack(packed, &obj)
	if err != nil {
		fmt.Println(err)
	}
	unpacked_val := unpacked.(map[any]any)
	fmt.Printf("Unpacked: %+v\n", unpacked_val["Typed"])

	fmt.Printf("object: %+v\n", obj)

	bytes := make([]float32, 3)
	bytes[0] = 1.1
	bytes[1] = 2.2
	bytes[2] = 3.3

	ser := msgpack.Serialize(bytes)
	fmt.Printf("Serialized: %+v\n", ser)

	des := msgpack.Deserialize(ser)
	fmt.Printf("Deserialized: %+v\n", des)

}
