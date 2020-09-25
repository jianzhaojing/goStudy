package main

import "fmt"

type Dog struct {
	name string
	age  int
}

func newDog(name string, age int) Dog {
	return Dog{
		name: name,
		age:  age,
	}
}

func (d *Dog) addAge() {
	d.age++
}

func (d Dog) wang() {
	fmt.Printf("%s:wangwangwang", d.name)
}

func main() {
	d := newDog("狗", 12)
	d.wang()
	fmt.Println(d)
	d.addAge()
	fmt.Println(d)
}
