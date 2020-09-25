package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Gender string `json:"gender" db:"gender"`
}

type Class struct {
	Title    string
	Students []*Student
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i < 200; i++ {
		stu := &Student{
			Id:     i,
			Name:   fmt.Sprintf("stu%03d", i),
			Gender: "男",
		}

		c.Students = append(c.Students, stu)
	}

	//json序列化
	// for _, v := range c.Students {
	// 	fmt.Println(v)
	// }
	data, err := json.Marshal(c) 
	//反json json.Unmarshal([]byte(str), &p2)
	if err != nil {
		fmt.Println("json failed")
		return
	}
	// fmt.Println(data)
	fmt.Printf("json:%s\n", data)
	// fmt.Println(c)
}
