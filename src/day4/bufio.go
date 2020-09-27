package main

import (
	"bufio"
	"fmt"
	"os"
)

func bufioRead() {
	file, err := os.Open("./bufio.go")
	if err != nil {
		fmt.Printf("读取失败，err:%v", err)
		return
	}

	defer file.Close()

	readFile := bufio.NewReader(file)
	// fmt.Println(readFile)
	for {
		line, err := readFile.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

func main() {
	bufioRead()
}
