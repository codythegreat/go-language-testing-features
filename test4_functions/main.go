package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(greeting("cody"))
	fmt.Println(getSum(2, 7))
}
