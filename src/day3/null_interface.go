package main

import "fmt"

func inter(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println(v)
	default:
		fmt.Println(v)
	}

}
func main() {
	// var x interface{}
	// x := "Hello 沙河"
	// v, ok := x.(string)
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("类型断言失败")
	// }

	inter(100)
}
