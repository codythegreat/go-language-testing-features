package main

import "fmt"

func main() {
	x := 10
	y := 10

	// conditional if/else statement
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Printf("%s is the color", color)
	} else if color == "blue" {
		fmt.Println("red is not the color")
	} else {
		fmt.Println("I don't know what the color is.")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")
	}
}
