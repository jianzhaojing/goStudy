package main

import (
	"bufio"
	"fmt"
	"os"
)

func UseBufio() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容")

	str, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", str)
}

func main() {
	UseBufio()
}
