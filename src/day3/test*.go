package main

import "fmt"

func zhizhen(x *int) {
	// fmt.Println(*x)
	*x = 100
}

func main() {
	a := 10
	zhizhen(&a)
	fmt.Println(a)
}
