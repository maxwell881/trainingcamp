package biz


import (
	"work3/internal/data"
)


type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error){
	err = data.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error){
	if err = data.DB.Find(&todoList).Error; err != nil{
		return nil, err
	}
	return
}

func GetATodo(id string)(todo *Todo, err error){
	todo = new(Todo)
	if err = data.DB.Debug().Where("id=?", id).First(todo).Error; err!=nil{
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo)(err error){
	err = data.DB.Save(todo).Error
	return
}

func DeleteATodo(id string)(err error){
	err = data.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}




