package main

import (
	"fmt"
	msgpack "github.com/codescalersinternships/Msgpack-Fatma-Ebrahim/pkg"
)

func main() {
	object := msgpack.Object{
		IsObject: true,
		Flag:  false,
		Unum: 0,
		Snum: -3,
		Fnum: 3.14,
		Str:  "Fatma",
	}
	bytes, err := msgpack.Pack(object)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bytes)
}

