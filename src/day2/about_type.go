package main

import "fmt"

//自定义类型
type myInt int

//类型别名
type ints = int

func main() {
	var i myInt
	var i2 ints
	fmt.Printf("type:%T value:%v", i, i)
	fmt.Println()

	fmt.Printf("type:%T value:%v", i2, i2)
	fmt.Println()
}
