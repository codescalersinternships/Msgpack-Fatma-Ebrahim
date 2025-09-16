# Msgpack-Fatma-Ebrahim
a Go library that implements MessagePack serialization and deserialization.

Check this link to learn more: https://github.com/msgpack/msgpack/blob/master/spec.md

## Functions:
### `Serialize(elementVal any) []byte`
serializes a single element into a byte array

### `Deserialize(value []byte) any`
deserializes a byte array into a single element

### `Pack(obj interface{}) ([]byte, error)`
packs an object into a byte array

### `Unpack(bytes []byte, obj interface{}) (any, error)`
unpacks a byte array into an object


## How to Use:
### Step 1: Install the library using `go get`

  ```bash
  go get github.com/codescalersinternships/Msgpack-Fatma-Ebrahim
  ```

This command fetches the library and adds it to your project's `go.mod` file.

### Step 2: Import and use the library in your code

  After running `go get`, you can import the library into your project and use the functions as described:


```
package main

import (
	"fmt"
	msgpack "github.com/codescalersinternships/Msgpack-Fatma-Ebrahim/pkg"
)

type Object struct {
	IsObject bool
	Flag     bool
}

func main() {
	object := Object{
		IsObject: true,
		Flag:     false,
	}

	packed, err := msgpack.Pack(object)
	if err != nil {
		fmt.Println(err)
	}

	var obj Object
	_, err = msgpack.Unpack(packed, &obj)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Object: %+v\n", obj)
}
```
