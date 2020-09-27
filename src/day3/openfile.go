package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./test.go")
	if err != nil {
		fmt.Printf("打开文件失败，err:%v", err)
		return
	}
	//关闭文件
	defer file.Close()

	//读文件
	var content []byte
	var tmp = make([]byte, 128)

	for {
		n, err := file.Read(tmp)
		fmt.Println(n)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}

		if err != nil {
			fmt.Printf("读取失败，err:%s", err)
		}

		content = append(content, tmp[:n]...)
	}

	fmt.Println(string(content))
}
