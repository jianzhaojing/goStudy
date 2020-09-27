package main

import "fmt"

func func1(ch1 chan int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}

	close(ch1)
}

func func2(ch1 chan int, ch2 chan int) {
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}

		ch2 <- tmp * tmp
	}

	close(ch2)
}

func main() {
	var ch1 = make(chan int, 100)
	var ch2 = make(chan int, 100)

	go func1(ch1)
	go func2(ch1, ch2)

	for item := range ch2 {
		fmt.Println(item)
	}

}
