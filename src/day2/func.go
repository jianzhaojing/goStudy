package main

import "fmt"

func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}
func main() {
	b := a("爸爸 ")
	b()

}
