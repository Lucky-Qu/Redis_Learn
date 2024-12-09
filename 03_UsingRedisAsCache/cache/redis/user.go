package redis

import (
	"Redis_Learn/03_UsingRedisAsCache/db/model"
	"context"
	"encoding/json"
	"strconv"
	"time"
)

func GetUser(index int) (*model.User, error) {
	userJson, err := rdb.Get(context.Background(), strconv.Itoa(index)).Result()
	if err != nil {
		return nil, err
	}
	var user model.User
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func SaveUser(index int, user *model.User) error {
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}
	rdb.SetEx(context.Background(), strconv.Itoa(index), userJson, 10*time.Second)
	return nil
}
