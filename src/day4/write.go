package main

import (
	"fmt"
	"os"
)

func write() {
	fileObj, err := os.OpenFile("./write2.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开失败")
		return
	}

	defer fileObj.Close()

	str := "沙河小王子"
	fileObj.Write([]byte(str)) //byte
	fileObj.WriteString("hello 沙河")
}
func main() {
	write()
}
