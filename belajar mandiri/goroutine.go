package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 10)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(time.Second * 10)
	}
}

func main() {
	go numbers()
	go alphabets()
	time.Sleep(time.Second * 10)
}