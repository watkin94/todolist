package db

import (
	"context"
	"fmt"
	"time"
	"github.com/go-redis/redis/v8"
)

func init()  {
	fmt.Println("------------------------------redis-init------------------------------------")
}

func RedisInit()  {
	//这里使用redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 服务器密码，如果没有密码则为空字符串
		DB:       0,                // Redis 数据库索引
	})


	
	// 创建一个上下文（Context）
	ctx := context.Background()
	//
	client.Set(ctx, "key", "value", 1*time.Hour)

	// 示例：获取键值对
	val, err := client.Get(ctx, "key").Result()

	if err == redis.Nil {
		fmt.Println("键值对不存在")
	} else if err != nil {
		fmt.Println("获取键值对失败:", err)
	} else {
		fmt.Println("获取到的值为:", val)
	}
}