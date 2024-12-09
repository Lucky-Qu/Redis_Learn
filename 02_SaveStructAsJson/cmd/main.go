package main

import (
	"Redis_Learn/02_SaveStructAsJson/db/redis/dao"
	"Redis_Learn/02_SaveStructAsJson/model"
	"errors"
	"fmt"
	"time"
)

func main() {
	var user = model.User{
		Username: "LuckyQu",
		Password: "123456",
		Age:      5,
	}
	if err := dao.SaveUser("user:1", &user); err != nil {
		panic(err)
	}

	if userGot, err := dao.GetUser("user:1"); err != nil {
		if errors.Is(err, errors.New("redis: nil")) {
			fmt.Println("未查询到用户")
		}
	} else {
		fmt.Println(*userGot)
	}

	time.Sleep(5 * time.Second)

	if userGot, err := dao.GetUser("user:1"); err != nil {
		fmt.Println("未查询到用户")
	} else {
		fmt.Println(*userGot)
	}
}
