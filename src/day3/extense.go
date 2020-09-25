package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Printf("%s会跑", a.name)
}

type Dog struct {
	feet uint8
	Animal
}

func (d Dog) waff() {
	fmt.Printf("%s会叫：汪汪", d.name)
}

func main() {
	dogs := Dog{
		feet:   4,
		Animal: Animal{name: "lele"},
	}
	dogs.move()
	dogs.waff()
}
