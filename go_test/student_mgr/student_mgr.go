package main

import "fmt"

type student struct {
	id   int64
	name string
}

//学生管理者
type studentMgr struct {
	allStudent map[int64]student
}

//查看学生
func (s studentMgr) showStudents() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}

//增加学生
func (s studentMgr) addStudent() {
	var (
		stuID   int64
		stuName string
	)
	//获取用户输入
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := student{
		id:   stuID,
		name: stuName,
	}

	s.allStudent[newStu.id] = newStu
}

//修改学生
func (s studentMgr) editStudent() {
	var stuID int64
	fmt.Print("请输入修改的学号：")
	fmt.Scanln(&stuID)

	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("需要修改的原信息： 学号：%d 姓名：%s\n", stuObj.id, stuObj.name)

	fmt.Print("请输入新的名字：")
	var newName string
	fmt.Scanln(&newName)
	stuObj.name = newName
	s.allStudent[stuID] = stuObj //更新map中的结构体需要整个结构体一起更新
}

//删除学生
func (s studentMgr) deleteStudent() {
	var stuID int64
	fmt.Print("请输入需要删除的学号：")
	fmt.Scanln(&stuID)

	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	delete(s.allStudent, stuID)
	fmt.Println("删除成功")
}
