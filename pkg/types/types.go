package types

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func differentTypes() {
	var yes bool = true
	displayWithType(yes)

	var text string = "bla"
	text2 := "other text"
	displayWithType(text)
	displayWithType(text2)

	// this will be the same as var i int = 32
	i := 32
	var i8 int8 = 8
	var i16 int16 = 16
	var i32 int32 = 32
	var i64 int64 = 64
	displayWithType(i)
	displayWithType(i8)
	displayWithType(i16)
	displayWithType(i32)
	displayWithType(i64)

	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // alias for uint8

	// rune // alias for int32
	// represents a Unicode code point

	var f32 float32 = 1.4512
	var f64 float64 = 1.00000000001
	displayWithType(f32)
	displayWithType(f64)

	var c128 complex128 = 17 + 4i
	displayWithType(c128)
}

func displayWithType(a interface{}) {
	fmt.Printf("Type: %v \tValue: %v\n", reflect.TypeOf(a), a)
}

func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	// %q	a double-quoted string safely escaped with Go syntax
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func conversions() {
	// Type inference (x and y are ints)
	x, y := 3, 4

	f := math.Sqrt(float64(x*x + y*y))

	var z uint = uint(f)

	fmt.Printf("sqrt(%v^2 + %v^2) = %v\n", x, y, z)
}

const Pi = 3.14

func constants() {
	const helloWorld = "hello 世界"

	fmt.Printf("%v: Pi %v is of the type %T\n", helloWorld, Pi, Pi)
}
