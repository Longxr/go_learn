package main

import (
	"fmt"
	"os"
)

//学生管理系统

var smr studentMgr

//菜单
func showMenu() {
	fmt.Println("-----------欢迎使用学生管理系统-----------")
	fmt.Println(`
	1. 查看所有学生
	2. 添加学生
	3. 修改学生
	4. 删除学生
	5. 退出

	`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}

	for {
		showMenu()
		//等待输入
		fmt.Print("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("选择的菜单项为：", choice)

		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入错误")
		}
	}
}
