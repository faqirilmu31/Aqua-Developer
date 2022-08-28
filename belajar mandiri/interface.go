package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(h HasName) {
	fmt.Println("Hello", h.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"

	SayHello(eko)

	cat := Animal{
		Name: "Kocheng",
	}
	SayHello(cat)
}