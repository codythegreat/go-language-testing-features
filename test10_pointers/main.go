package main

import "fmt"

func main() {
	a := 5
	b := &a // pointer

	// print a and memory address of a
	fmt.Println(a, b)

	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a) // these are the same thing

	// change val with pointer
	*b = 10
	fmt.Println(a)

	// why use pointers? when passing a lot of data through
	// functions pointers can be used to boost performance
	// (since the address is being passed.)

}
