package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int

	// short verion: firstName, lastName, city, gender string
	// age	int
}
// method does not go into the struct
// Greeting method (value receiver, didnt change anything)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Sam", lastName: "Smith", city: "Boston", gender: "F", age: 25}
	fmt.Println(person1)

	// Alternative
	person2 := Person{"Bob", "Johnson", "Atlanta", "M", 40}
	fmt.Println(person2)

	fmt.Println(person1.firstName)

	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	person2.getMarried("Toms")
	fmt.Println(person2.greet())
	
}