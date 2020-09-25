package main

import (
	"fmt"
	"os"
)

type Student struct {
	id    int
	name  string
	class int
}

var classMap map[int]*Student

func newStudent(id int, name string, class int) *Student {
	return &Student{
		id:    id,
		name:  name,
		class: class,
	}
}

func listStudents() {
	for _, v := range classMap {
		fmt.Printf("学生姓名：%s，学号：%d，班级：%d班", v.name, v.id, v.class)
		fmt.Println()
	}
}

func addStudents() {
	var (
		id    int
		name  string
		class int
	)

	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入学生班级：")
	fmt.Scanln(&class)

	classMap[id] = newStudent(id, name, class)
}

func updateStudents() {
	var (
		id    int
		name  string
		class int
	)

	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入学生班级：")
	fmt.Scanln(&class)

	classMap[id].name = name
	classMap[id].class = class
}

func deleteStudents() {
	var id int
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)

	delete(classMap, id)
}

func exits() {
	os.Exit(1)
}

func main() {
	//初始化
	classMap = make(map[int]*Student, 100)
	for {
		//欢迎界面
		fmt.Println("欢迎登陆学生管理系统")
		fmt.Println(`
	1 查询
	2 新增
	3 修改
	4 删除
	5 退出   
`)

		//检查输入的值
		var num int
		fmt.Scanln(&num)
		fmt.Printf("您已选择了%d这个选项", num)
		fmt.Println()

		//匹配
		switch num {
		case 1:
			listStudents()
		case 2:
			addStudents()
		case 3:
			updateStudents()
		case 4:
			deleteStudents()
		default:
			os.Exit(1)
		}
	}

}
