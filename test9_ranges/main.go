package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 32, 11, 5}

	// loop through ids

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using index
	// underscore is used to not cause unused var error
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum", sum)

	// range with map
	emails := map[string]string{"dwight": "dwight@org.com", "megan": "megan@org.com", "cody": "cody@org.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("name: ", k)
	}
}
