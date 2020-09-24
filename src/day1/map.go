package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//map（映射）
func main() {
	//按照某个固定的顺序遍历map
	var scoreMap = make(map[string]int, 100)

	//添加50个键值对
	for i := 1; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		// fmt.Println(key)
		value := rand.Intn(100)
		scoreMap[key] = value

	}

	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }

	//按照key从小到大的排序
	//1. 先去除key放到切片中
	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}

	//2  对key做排序
	sort.Strings(keys)
	// fmt.Println(keys)
	//3  按照排序后的key对scoreMap排序
	for _, val := range keys {
		fmt.Println(val, scoreMap[val])
	}
}
