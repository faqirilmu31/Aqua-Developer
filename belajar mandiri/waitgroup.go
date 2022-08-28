package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("hallo")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("world")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("reni")
	}()

	wg.Wait()
}
