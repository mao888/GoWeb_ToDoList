package models

import (
	"bubble/dao"
)

// Todo Model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

/*
	Todo这个Model的增删改查操作都放在这里
 */
// CreateATodo 创建todo
// 添加
func CreateATodo(todo *Todo) (err error){
	err = dao.DB.Create(&todo).Error
	return
}

// 查看所有的待办事项
func GetAllTodo() (todoList []*Todo, err error){
	if err = dao.DB.Find(&todoList).Error; err != nil{
		return nil, err
	}
	return
}

// 查一条todo
func GetATodo(id string)(todo *Todo, err error){
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err!=nil{
		return nil, err
	}
	return
}

// 修改某一个待办事项
func UpdateATodo(todo *Todo)(err error){
	err = dao.DB.Save(todo).Error
	return
}

// 删除某一个待办事项
func DeleteATodo(id string)(err error){
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}

