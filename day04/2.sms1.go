package main

import (
	"fmt"
	"os"
)

//函数版学生管理系统

// student 这是一个学生的结构体
var allStudents map[int64]*student

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
func listAllStudent() {
	for k, v := range allStudents {
		fmt.Printf("学号：%d,姓名：%s\n", k, v.name)
	}
}
func addOneStudent() {

	var (
		id   int64
		name string
	)
	fmt.Printf("请输入新增学生的学号：")
	fmt.Scanln(&id)
	fmt.Printf("请输入新增学生的姓名：")
	fmt.Scanln(&name)
	allStudents[id] = newStudent(id, name)
}
func deleteOneStudent() {
	var id int64
	fmt.Printf("请输入删除学生的学号：")
	fmt.Scanln(&id)
	delete(allStudents, id)

}
func main2() {
	allStudents = make(map[int64]*student, 10)
	for {
		// 1.输出欢迎界面
		fmt.Print(`欢迎光临学生管理系统！您可以选择如下选项：
		1.查看所有学生;
		2.新增一名学生;
		3.删除一名学生;
		4.退出系统
	请选择：`)
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("您选择了%d操作\n", choice)
		switch choice {
		case 1:
			listAllStudent()
		case 2:
			addOneStudent()
		case 3:
			deleteOneStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("请重新输入！")
		}
	}

}
