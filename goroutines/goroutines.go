package goroutines

import (
	"fmt"
	"runtime"
	"sync"
)

var wait_group = sync.WaitGroup{}
var counter int = 0
var mutex = sync.RWMutex{}

func increment() {
	counter++
	mutex.Unlock()
	wait_group.Done()
}

func say_hello(number, thread int) {
	fmt.Printf("thread: %v counter: %v\n", thread, counter)
	mutex.RUnlock()
	wait_group.Done()
}

func Go() {
	fmt.Println("  Goroutines")

	fmt.Printf("Number of threads: %v\n", runtime.GOMAXPROCS(-1))

	// Gortoutines (green threads) and mutexes
	message := "hello"

	wait_group.Add(2)

	mutex.RLock()
	go say_hello(45, -1)
	go func(msg string) {
		fmt.Println(msg)
		wait_group.Done()
	}(message)

	message = "Goodbye"

	wait_group.Wait()

	for i := 0; i < 10; i++ {
		wait_group.Add(2)
		mutex.Lock()
		go increment()
		mutex.RLock()
		go say_hello(counter, i)
	}

	wait_group.Wait()
}
