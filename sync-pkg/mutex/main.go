package main

import (
	"fmt"
	"sync"
)

func main() {
	withoutMutex()
	withMutex()
}

// produce race condition
func withoutMutex() {
	var count int

	increment := func(wg *sync.WaitGroup) {
		defer wg.Done()
		count++
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("done - withoutMutex. Count is ", count)
}

// no race condition
func withMutex() {
	var count int
	var lock sync.Mutex

	increment := func(wg *sync.WaitGroup) {
		defer wg.Done()

		lock.Lock()
		count++
		lock.Unlock()
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("done - withMutex. Count is ", count)
}
