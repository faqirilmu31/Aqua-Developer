package main

import "fmt"

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return
}

func main() {
	var a, p int
	a, p = rectangle(10, 20)
	fmt.Println("Area = ", a)
	fmt.Println("Parameter = ", p)
}