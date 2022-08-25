package main

import "fmt"

func main() {
	var student map[string]int
	student = map[string]int{}

	student["Reni"] = 13
	student["Setya"] = 14
	fmt.Println(student)
}