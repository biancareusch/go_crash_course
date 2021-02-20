package main

import (
	"fmt"
	"strconv"
)

//Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

//Greeting method (value receiver) [getter]
func (p Person) greet() string {
	return "Hello, my Name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

//hasBirthday (pointer receiver) [setter]
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getsMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//Init person using struct
	person1 := Person{firstName: "Bianca", lastName: "Reusch", city: "SA", gender: "f", age: 27}
	person2 := Person{"Bianca", "Reusch", "SA", "f", 27}
	person3 := Person{"Bob", "Johnson", "SA", "m", 27}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)

	person1.hasBirthday()
	person1.getsMarried("Williams")
	person3.getsMarried("Thompson")
	fmt.Println(person1.greet())
	fmt.Println(person3.greet())

}
