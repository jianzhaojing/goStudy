package main

import "fmt"

func main() {
	var mapSlice = make([]map[int]string, 3)
	if mapSlice[0] == nil {
		fmt.Println("相等")
	} else {
		fmt.Println("不想等")
	}
	mapSlice[0] = make(map[int]string, 10)
	mapSlice[0][100] = "沙河"
	fmt.Println(mapSlice)
}
