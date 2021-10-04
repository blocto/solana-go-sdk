package borsh

// Simple Enum type in Go.
//  type MyEnum borsh.Enum
//  const (
//    A MyEnum = iota
//    B
//    C
//  )
//
// Complex Enum type in Go.
//  type MyEnum struct {
//    Enum borsh.Enum `borsh_enum:"true"`
//    Foo  Foo
//    Bar  Bar
//  }
//
//  type Foo struct {
//	  FooA int32
//	  FooB string
//  }
//
//  type Bar struct {
//	  BarA int64
//	  BarB string
//  }
type Enum uint8
