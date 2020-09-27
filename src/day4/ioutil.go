package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./bufio.go")
	if err != nil {
		fmt.Println("读取失败")
		return
	}

	fmt.Println(string(content))
}
