package main

import (
	"fmt"
	msgpack "messagepack/pkg"
)

func main() {
	object := msgpack.Object{
		IsObject: true,
		Flag:  false,
		Unum: 0,
		Snum: -3,
		Fnum: 3.14,
	}
	bytes, err := msgpack.Pack(object)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bytes)
}

