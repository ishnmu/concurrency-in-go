package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepRoutine(m data, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Printf("good night: %s\n", m.message)
	time.Sleep(m.sleepTime)
	fmt.Printf("good morning: %s\n", m.message)
}

func sleepRoutine1(m data) {
	fmt.Printf("good night: %s\n", m.message)
	time.Sleep(m.sleepTime)
	fmt.Printf("good morning: %s\n", m.message)
}

type data struct {
	message   string
	sleepTime time.Duration
}

func main() {
	withoutWaitGroup()
	fmt.Println("++++++++++++++++++++++++++++++")
	withWaitGroup()

}

func withWaitGroup() {
	var wg sync.WaitGroup

	for _, d := range []data{
		{
			message:   "go-1",
			sleepTime: 5 * time.Second,
		},
		{
			message:   "go-2",
			sleepTime: 10 * time.Second,
		},
	} {
		wg.Add(1) //! wg.Add should be called outside/before invoking go routines
		go sleepRoutine(d, &wg)
	}

	wg.Wait()

	fmt.Println("all go routines completed")
}

func withoutWaitGroup() {

	for _, d := range []data{
		{
			message:   "go-1-without-wg",
			sleepTime: 5 * time.Second,
		},
		{
			message:   "go-2-without-wg",
			sleepTime: 10 * time.Second,
		},
	} {
		go sleepRoutine1(d)
	}

	fmt.Println("all go routines completed")
}
