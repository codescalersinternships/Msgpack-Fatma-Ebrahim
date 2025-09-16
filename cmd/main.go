package main

import (
	"fmt"
	msgpack "github.com/codescalersinternships/Msgpack-Fatma-Ebrahim/pkg"
)

func main() {
	arr := make([]any, 4)
	arr[0] = "hello"
	arr[1] = 7
	arr[2] = 9.8
	arr[3] = "world"

	m := make(map[any]any)
	m["hello"] = 10
	m["world"] = -20

	object := msgpack.Object{
		IsObject: true,
		Flag:     false,
		Unum:     0,
		Snum:     -1000,
		Fnum:     3.14,
		Str:      "Fatma",
		Arr:      arr,
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
	fmt.Printf("Unpacked: %+v\n", unpacked)
	fmt.Printf("Object: %+v\n", obj)

}
