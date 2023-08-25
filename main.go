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
	secretkey := []byte("111") //密钥，不能泄漏
	//生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"foo":"bar",
		"nbf":time.Date(2022,10,10,12,0,0,0,time.UTC).Unix(),
	})

	//利用token对象的签名方法生成签名字符串
	tokenString,err := token.SignedString(secretkey)
	fmt.Println(tokenString,err)


	fmt.Println("-------------------------------测试检验token-------------------------------------")
	//tokeTmp := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE2OTY5MzkyMDB9.xQTqqa8sylUhy7hoKJPB6GK_54hhfuSm3V3p-8soqOU"

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
	//fmt.Println(token_re,err)

	/*****************************************写个中间件**********************************************/
	//return;
	db.Connect()
	model.CreateTodo()	//创建表迁移
	defer db.Close()
	db.RedisInit()	//测试redis
	route.ApiInit()
}