package models

import "web/bubble/dao"

// TodoModel
type Todo struct {
	ID int	`json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

// CreateATodo 创建todo
func CreateATodo(todo *Todo)(err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo()(todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	return
}

func GetATodo(id string)(todo Todo, err error) {
	err = dao.DB.Where("id=?", id).First(&todo).Error
	return
}

func UpdateATodo(todo *Todo)(err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteATodo(id string)(err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return
}