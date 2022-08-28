package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"

	// akan mengubah semua data di address2 tanpa mempengaruhi address1
	address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}

	// akan mengubah semua data yg mengarah ke address1
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}
	
	// maka addres1 city tidak akan berubah
	fmt.Println(address1)
	fmt.Println(address2)
}