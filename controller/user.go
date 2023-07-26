package controller

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	//"todolist/db"
	"todolist/model"
)

//用户接口
func UserIndex(c *gin.Context)  {
	//用户信息
	users := []string {"user1","user2","user3"}
	c.JSON(200,gin.H{"users":users})
}

func Users(c *gin.Context)  {
	// 处理获取用户列表的逻辑
	users:=model.Test()

	//users := []string{"user1", "user2", "user3"}
	c.JSON(200, gin.H{"users": users})
}
