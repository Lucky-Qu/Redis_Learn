package service

import (
	"Redis_Learn/03_UsingRedisAsCache/cache/redis"
	"Redis_Learn/03_UsingRedisAsCache/db/dao"
	"Redis_Learn/03_UsingRedisAsCache/db/model"
	"fmt"
)

func GetUser(index int) (*model.User, error) {
	user, err := redis.GetUser(index)
	if err == nil {
		fmt.Println("这次是从缓存中取出数据")
		return user, nil
	} else {
		user, err = dao.GetUser(index)
		if err != nil {
			return nil, err
		}
		err = redis.SaveUser(index, user)
		if err != nil {
			fmt.Println("存入缓存出错")
		}
		fmt.Println("这次是从数据库取出数据")
		return user, nil
	}
}

func CreateUser(user *model.User) error {
	if err := dao.CreateUser(user); err != nil {
		return err
	}
	return nil
}
