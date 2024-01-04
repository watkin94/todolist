package main

import (
	"fmt"
	"time"
	"todolist/db"
	"todolist/model"
	"todolist/route"
	"github.com/golang-jwt/jwt/v4"
)

type UsersClaims struct {
	Username string `json:"username"`
	Gender string `json:"gender"`
	Avatar string `json:"avatar"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func main() {
	fmt.Println("测试生成jwt")
	secretkey := []byte("watkin1994") //密钥，不能泄漏
	//生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"foo":"bar",
		"nbf":time.Date(2022,10,10,12,0,0,0,time.UTC).Unix(),
	})

	//利用token对象的签名方法生成签名字符串
	tokenString,err := token.SignedString(secretkey)
	fmt.Println(tokenString,err)


	fmt.Println("-------------------------------测试检验token-------------------------------------")

	token_re,err := jwt.ParseWithClaims(tokenString,&UsersClaims{}, func(t *jwt.Token)(interface{},error) {
		return secretkey,nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("检验成功!!!")
	claims,ok := token_re.Claims.(*UsersClaims)
	valid := token_re.Valid
	fmt.Println(claims,ok,valid)



	/*****************************************调用中间件**********************************************/
	db.Connect()		//连接mysql
	model.CreateTodo()	//创建表迁移
	defer db.Close()	//延迟关闭mysql
	db.RedisInit()		//redis初始化
	route.ApiInit()		//路由初始化
}