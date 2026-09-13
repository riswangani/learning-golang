package main

import (
	"fmt"
)

type HasName interface {
	GetName() string
	GetAge() int
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
	fmt.Println("Age", value.GetAge())
}

type Person struct {
	Name string
	// Address string
	Age int
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAge() int {
	return person.Age
}


type Animal struct {
	Name string
	Age  int
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (animal Animal) GetAge() int {
	return animal.Age
}

func main() {
	person := Person{Name: "Gani", Age: 30}
	animal := Animal{Name: "Kucing", Age: 0}
	SayHello(person)
	SayHello(animal)
}
