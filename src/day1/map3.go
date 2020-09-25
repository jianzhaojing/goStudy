package main

import "fmt"

func main() {
	var sliceMap = make(map[int][]string, 5)

	v, k := sliceMap[0]
	if k {
		fmt.Println(v)
	} else {
		sliceMap[0] = make([]string, 10, 10)
		sliceMap[0][0] = "中国"
		fmt.Println(sliceMap[0])
	}

}
