package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var xiaoming Person
	xiaoming.name = "小明"
	xiaoming.age = 18
	fmt.Printf("type:%T value:%#v", xiaoming, xiaoming)
	fmt.Println()

	var xiaohong = new(Person)
	fmt.Printf("type:%T value:%#v", xiaohong, xiaohong)
	fmt.Println()

	var a = 100
	b := &a
	fmt.Printf("a:%d,b:%v", a, b)
	fmt.Println()

	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b)

}
