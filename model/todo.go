package model
import (
	"todolist/db"
)

type Todo struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `gorm:"not null" json:"title"`
	Status bool `gorm:"not null" json:"status"`
}

//创建表TODO
func CreateTodo()  {
	// 自动迁移数据库表
	db.DB.AutoMigrate(&Todo{})
}