package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var saldo int64 = 0

var isLock bool = false

var lock sync.Mutex


func addSaldo() {
	for i := 0; i < 10000000; i++ {
		atomic.AddInt64(&saldo, 1)
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
		addSaldo()
	}()

	go func() {
		defer wg.Done()
		addSaldo()
	}()

	wg.Wait()
	fmt.Println(saldo)
}
