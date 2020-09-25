package main

import (
	"fmt"
	"unicode"
)

var str = " hello 中国"

func main() {
	var count int
	for _, value := range str {
		// fmt.Println(value)
		if unicode.Is(unicode.Han, value) {
			count++
		}
	}

	fmt.Println(count)
}
