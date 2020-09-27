package main

import (
	"fmt"
	"io/ioutil"
)

func ioutilWrite() {
	str := "沙河小昂"
	err := ioutil.WriteFile("./write2.go", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	ioutilWrite()
}
