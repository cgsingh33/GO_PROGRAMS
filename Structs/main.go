package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	constant  contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.constant.email = "alex@gmail.com"
	alex.constant.zipCode = 12345
	alex.newFstname("alexa")
	alex.print()
}

func (pointer *person) newFstname(newFirstname string) {
	(*pointer).firstName = newFirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
