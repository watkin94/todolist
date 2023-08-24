package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"todolist/db"
	"context"
)


//用于签名的字符串
var mySigningKey = []byte("watkin")


//直接使用默认声明创建jwt
//func GenRegisteredClaims() (string,error) {
	//创建Claims
	//claims:=&jwt.RegisteredClaims{
	//	ExpiresAt : jwt.NewNumericDate(),
	//}
//}


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


//带协程的gouser
func GoUser(c *gin.Context)  {
	//模拟用协程处理异步任务
	go go1()
	go go2()
	res := []string {"dddd"}
	c.JSON(200, gin.H{"users": res})
}