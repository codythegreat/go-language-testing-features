package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [3]string

	// assign values
	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"
	// fruitArr[2] = "banana"

	// Declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	// Slices
	fruitSlice := []string{"Apple", "Orange", "Banana"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
