package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "pisang", "jeruk", "duku"}

	for i, b := range buah {
		fmt.Printf("Elemen %d : %s\n", i, b)
	}
}