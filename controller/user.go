package controller

import (
	"fmt"
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


//测试--在gin控制器中使用redis===>	// 处理获取用户列表的逻辑
func Users(c *gin.Context)  {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db.Client.Set(ctx, "key3", "usersssss", 1*time.Hour)	//调用redis
	redis_user ,err := db.Client.Get(ctx, "key3").Result()

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"users": redis_user})
}


//带协程的gouser===>模拟用协程处理异步任务,  但是运行顺序是先返回接口。然后再循环数组
func GoUser(c *gin.Context)  {
	go go1()
	go go2()
	res := []string {"dddd"}
	c.JSON(200, gin.H{"users": res})
}


func go1()  {
	//异步任务1
	for i:=0;i < 100000;i++{
		fmt.Println(i)
	}
}

func go2 ()  {
	//异步任务2
	for i:=0;i < 100000;i++{
		fmt.Println(i)
	}
}


