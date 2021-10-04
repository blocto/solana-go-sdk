# borsh-go

[![Go Reference](https://pkg.go.dev/badge/github.com/near/borsh-go.svg)](https://pkg.go.dev/github.com/near/borsh-go)

**borsh-go** is an implementation of the [Borsh] binary serialization format for Go
projects.

Borsh stands for _Binary Object Representation Serializer for Hashing_. It is
meant to be used in security-critical projects as it prioritizes consistency,
safety, speed, and comes with a strict specification.

## Features

- Based on Go Reflection. Avoids the need for create protocol file and code generation. Simply
defining `struct` and go.


## Usage

### Example

```go
package demo

import (
	"log"
	"reflect"
	"testing"

	"github.com/near/borsh-go"
)

type A struct {
	X uint64
	Y string
	Z string `borsh_skip:"true"` // will skip this field when serializing/deserializing
}

func TestSimple(t *testing.T) {
	x := A{
		X: 3301,
		Y: "liber primus",
	}
	data, err := borsh.Serialize(x)
	log.Print(data)
	if err != nil {
		t.Error(err)
	}
	y := new(A)
	err = borsh.Deserialize(y, data)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(x, *y) {
		t.Error(x, y)
	}
}
```

For more examples of usage, refer to `borsh_test.go`.

## Type Mappings

Borsh                 | Go           |  Description
--------------------- | -------------- |--------
`bool`		      | `bool`	       |
`u8` integer          | `uint8`        |
`u16` integer         | `uint16`       |
`u32` integer         | `uint32`       |
`u64` integer         | `uint64`       |
`u128` integer        | `big.Int`  |
`i8` integer          | `int8`        |
`i16` integer         | `int16`       |
`i32` integer         | `int32`       |
`i64` integer         | `int64`       |
`i128` integer        |            |  Not supported yet
`f32` float           | `float32`      |
`f64` float           | `float64`      |
fixed-size array      | `[size]type`   |  go array
dynamic-size array    |  `[]type`      |  go slice
string                | `string`       |
option                |  `*type`         |   go pointer
map                   |   `map`          |
set                   |   `map[type]struct{}`  | go map with value type set to `struct{}`
structs               |   `struct`      |
enum                  |   `borsh.Enum`  |    use `type MyEnum borsh.Enum` to define enum type
