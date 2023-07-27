package model

import (
	"todolist/db"
)

type User struct {
	//gorm.Model
	Name  string
	Email string
}

//定义表名
func (User) TableName() string {
	return "user"
}

//demo
func Test() []User {
	//对数据进行增删改查
	//user := User{Name: "John Doe", Email: "john@example.com"}
	//db.Create(&user)
	//现在对其进行查询
	var users []User
	db.DB.Find(&users)

	return users
}