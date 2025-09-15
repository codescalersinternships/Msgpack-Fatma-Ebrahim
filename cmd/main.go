package main

import (
	"fmt"
	msgpack "github.com/codescalersinternships/Msgpack-Fatma-Ebrahim/pkg"
)

func main() {
	arr := make([]any, 3)
	arr[0] = "hello"
	arr[1] = 7
	arr[2] = 9.8

	object := msgpack.Object{
		IsObject: true,
		Flag:  false,
		Unum: 0,
		Snum: -1000,
		Fnum: 3.14,
		Str:  "Fatma",
		Arr:  arr,
	}
	bytes, err := msgpack.Pack(object)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bytes)
}

