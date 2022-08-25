package main

import "fmt"

type student struct {
	name string
	age  int

	sekolah school
}

type school struct {
	name string
	address string
}

func main() {
	var s1 student
	s1.name = "Reni"
	s1.age = 20
	s1.sekolah.name = "Universitas Sebelas Maret"
	s1.sekolah.address = "Klaten"

	students := []student{
		{name: "Reni", age:20, sekolah: school{name: "UNS", address: "Solo"}},
		{name: "Setya", age:21, sekolah: school{name: "UGM", address: "Jogja"}},
	}

	for _, student := range students {
		fmt.Println("name : ", student.name)
		fmt.Println("age : ", student.age)
		fmt.Println("----------------------")
	}

	fmt.Println("Name : ", s1.name)
	fmt.Println("Age : ", s1.age)
	fmt.Println("School : ", s1.sekolah.name)
	fmt.Println("Address : ", s1.sekolah.address)
}

