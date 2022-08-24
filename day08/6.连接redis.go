package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisDb *redis.Client

func main() {

	redisDb := redis.NewClient(&redis.Options{
		Addr:     "11.2.2.128:6379",
		Password: "",
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := redisDb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis失败,err:%v", err)
		return
	}
	fmt.Println("连接redis成功")

	value, err := redisDb.Get(ctx, "dsb2").Result()
	if err != nil {
		fmt.Printf("Get key failed,err:%v", err)
		return
	}
	fmt.Println(value)

	valueCmd := redisDb.Get(ctx, "key")
	fmt.Println(valueCmd.Val(), valueCmd.Err())

}
