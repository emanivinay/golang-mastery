package main

import (
	"fmt"
	"sync"
)

func createChan[T any]() chan T {
	channel := make(chan T)
	return channel
}

func main() {
	fmt.Println("Hello, World!")

	intChan := createChan[int]()
	var wg sync.WaitGroup
	wg.Add(2)
	sender := func(v int) {
		defer wg.Done()
		intChan <- v
		close(intChan)
	}

	receiver := func() {
		defer wg.Done()
		v := <-intChan
		fmt.Println(v)
	}

	go sender(100)
	go receiver()
	wg.Wait()
}
