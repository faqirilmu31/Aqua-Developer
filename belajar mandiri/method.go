package main

// import "strings"
import "fmt"

type student struct {
	name string
	grade int
}

func main() {
	var student1 = student{"Reni", 10}
	student1.SayHello()

	student1.changeName("Reni Setyaningsih")
	student1.SayHello()
}

func (s student) SayHello() {
	fmt.Println("Hello", s.name)
}

func (s *student) changeName(name string) {
	s.name = name
}