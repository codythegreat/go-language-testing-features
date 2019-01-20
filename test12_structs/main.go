package main

import (
	"fmt"
	"strconv"
)

// define person struct

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	// you can also put all data types on one line like this:
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)

func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct

	person1 := Person{firstName: "Cody", lastName: "Maxie", city: "Durant", gender: "m", age: 24}

	person2 := Person{"Megan", "Claborn", "Durant", "f", 22}
	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	person2.age++
	fmt.Println(person2.age)

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person2.getMarried("Maxie")
	fmt.Println(person2.greet())
}
