package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/model"
	"todolist/db"
)

func GetTodo(c *gin.Context){
	//查询todo表的所有数据
	var todoList []model.Todo
	if err:=db.DB.Find(&todoList).Error;err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,todoList)
	}
}


func AddTodo(c *gin.Context) {
	var todo model.Todo
	c.BindJSON(&todo)
	if err := db.DB.Create(&todo).Error;err != nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,todo)
	}
}

func EditTodo(c *gin.Context) {
	id,_ := c.Params.Get("id")
	var todo model.Todo
	if err := db.DB.Where("id=?",id).First(&todo).Error;err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}

	c.BindJSON(&todo)
	if err := db.DB.Save(&todo).Error;err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,todo)
	}
}


func DeleteTodo(c *gin.Context) {
	id,ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"error":"无效id"})
	}
	if err:=db.DB.Where("id=?",id).Delete(model.Todo{}).Error;err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,gin.H{id:"deleted"})
	}
}

