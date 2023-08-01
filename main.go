package main

import (
	"todolist/db"
	"todolist/model"
	"todolist/route"
)

func main() {
	db.Connect()
	model.CreateTodo()	//创建表迁移
	defer db.Close()
	db.RedisInit()	//测试redis
	route.ApiInit()
}