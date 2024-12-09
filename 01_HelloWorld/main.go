package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := rdb.SetEx(context.Background(), "sentence", "HelloWorld", 5*time.Second).Err(); err != nil {
		panic(err)
	}
	fmt.Println(rdb.Get(context.Background(), "sentence").String())
	time.Sleep(6 * time.Second)
	fmt.Println(rdb.Get(context.Background(), "sentence").String())
}
