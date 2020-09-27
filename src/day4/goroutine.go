package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello", i)
	wg.Done()
}
func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go hello(i)
	}

	fmt.Println("hello world")

	wg.Wait()
}
