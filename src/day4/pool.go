package main

import (
	"fmt"
	"time"
)

func work(j int, jobs <-chan int, result chan<- int) {
	for item := range jobs {
		fmt.Printf("work%d开始工作，job:%d", j, item)
		fmt.Println()
		time.Sleep(time.Second)
		fmt.Printf("work%d开始结束，job:%d", j, item)
		fmt.Println()
		result <- item * 2
	}
	fmt.Println()
}

func main() {
	var jobs = make(chan int, 100)
	var result = make(chan int, 100)

	//生产5个job
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	//生成3个gorountine
	for j := 0; j < 3; j++ {
		go work(j, jobs, result)
	}

	for i := 0; i < 5; i++ {
		<-result
	}
}
