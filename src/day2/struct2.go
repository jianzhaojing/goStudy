package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//初始化一
	var p1 = new(Person)
	fmt.Println(p1)
	//初始化二
	var p2 = &Person{}
	fmt.Println(p2)
	//初始化三
	var p3 = Person{}
	fmt.Println(p3)
	//初始化四
	var p4 = Person{
		name: "小明",
		age:  18,
	}
	fmt.Println(p4)
	//初始化五
	var p5 Person
	p5.name = "小明"
	p5.age = 18

	fmt.Println(p5)
	//初始化
	var p6 = Person{
		"小明",
		18,
	}
	fmt.Println(p6)
}
