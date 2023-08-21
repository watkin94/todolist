package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"todolist/db"
	"context"
)

//测试
func UserIndex(c *gin.Context)  {
	//用户信息
	users := []string {"user1","user2","user3"}
	c.JSON(200,gin.H{"users":users})
}


//测试--在gin控制器中使用redis
func Users(c *gin.Context)  {
	// 处理获取用户列表的逻辑
	//users := []string {"user1","user2","user3"}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db.Client.Set(ctx, "key3", "usersssss", 1*time.Hour)
	redis_user ,err := db.Client.Get(ctx, "key3").Result()

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"users": redis_user})
}
