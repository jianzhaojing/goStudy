package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func Add() {
	for i:=0;i<5000;i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}

	wg.Done()
}

func main() {
	wg.Add(2)

	go Add()
	go Add()

	wg.Wait()

	fmt.Println(x)
}
