package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
	"runtime"
)

var (
	DB *gorm.DB
)

func Connect() {
	//引入环境变量
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	//使用环境变量   dns ="root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	dns_str := os.Getenv("DB_USERNAME") + ":"+ os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	var dns string
	if runtime.GOOS == "windows"{
		dns = dns_str
	}else{
		dns ="自己填上服务相关linux 的信息即可"
	}
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		panic("无法连接到数据库")
	}
	DB = db
}

func Close()  {
	defer DB.Close()
}

