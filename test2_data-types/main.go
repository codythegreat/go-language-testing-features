package main

import "fmt"

// you can make vars outsite of a function
var codyName string = "Name is Cody"

// but you cannot use shorthand

func main() {
	// DATA TYPES:
	// string
	// bool
	// int
	// int int8 int16 int32 int 64
	// uint uint8 uint16 uint32 uint 64 uintptr
	// byte - alias for uint8
	// rune - aliad for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name string = "Cody"
	var age byte = 24
	var isCool = true

	fmt.Println(name, age, isCool, codyName)
	fmt.Printf("%T\n", name)
}
