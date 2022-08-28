package main

import (
	"fmt"
	"sync"
)

var saldo int = 0

var isLock bool = false

func addSaldo() {
	for i := 0; i < 10000000; i++ {
		saldo++
	}
}

func Lock(lockedFunc func()) {
	for isLock {
		if !isLock {
			break
		}
	}

	isLock = true

	lockedFunc()

	isLock = false
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		Lock(addSaldo)
	}()

	go func() {
		defer wg.Done()
		Lock(addSaldo)
	}()

	wg.Wait()
	fmt.Println(saldo)
}