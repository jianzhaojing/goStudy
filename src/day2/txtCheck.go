package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	newFunc := func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}

	return newFunc
}

func main() {
	func2 := makeSuffixFunc(".txt")

	res := func2("古力娜扎")
	fmt.Println(res)
}
