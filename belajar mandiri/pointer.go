package main

import "fmt"

func main() {
	var a string = "hello world"
	fmt.Println("a = ", a)

	var b *string

	b = &a

	fmt.Println("b = ", b)
	fmt.Println("value b = ", *b)
}