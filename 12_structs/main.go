package main

import (
	"fmt"
	"strconv"
)

type Person struct{
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName, lastName, city, gender string
	age int
}

// Value receiver function

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Pointer receiver function (change the person attribute value)

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	}
	p.lastName = spouseLastName
}

func main() {
	// person1 := Person{firstName: "Vincent", lastName: "Andrian", city:"Singapore", gender:"Male", age:25}

	person1 := Person{"Samantha", "Smith", "Singapore", "Female", 25}

	// fmt.Println(person1)
	// fmt.Println(person1.age)
	// person1.age = 30
	// fmt.Println(person1.age)

	person1.hasBirthday()
	person1.getMarried("Andrian")
	fmt.Println(person1.greet())
	
}