package main

import "fmt"

func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println("Total = ", total)
}

func main() {
	add(20, 30)
}