package main

import "fmt"

func main() {
	// define a map
	emails := make(map[string]string)

	// assign key/value pairs

	emails["bob"] = "bob@org.com"
	emails["cody"] = "cody@org.com"

	fmt.Println(emails)
	fmt.Println(emails["cody"])
	fmt.Println(len(emails))

	// delete from map
	delete(emails, "bob")
	fmt.Println(len(emails))

	// assign upon declaration
	emailsAssigned := map[string]string{"bob": "bob@org.com", "cody": "cody@org.com"}
	emailsAssigned["mike"] = "mike@org.com"
	fmt.Println(emailsAssigned)
}
