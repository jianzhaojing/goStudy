package main

import (
	"fmt"
	"strings"
)

var str = "how do you do"

func main() {
	newStr := strings.Split(str, " ")
	wordCount := make(map[string]int, 10)
	for _, value := range newStr {
		_, k := wordCount[value]
		if k {
			wordCount[value]++
		} else {
			wordCount[value] = 1
		}
	}
	// fmt.Printf("%T", newStr)
	fmt.Println(wordCount)
}
