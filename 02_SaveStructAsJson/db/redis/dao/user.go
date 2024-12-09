package dao

import (
	"Redis_Learn/02_SaveStructAsJson/model"
	"context"
	"encoding/json"
	"time"
)

func SaveUser(index string, user *model.User) error {
	jsonUser, err := json.Marshal(user)
	rdb.SetEx(context.Background(), index, jsonUser, 3*time.Second)
	return err
}

func GetUser(index string) (*model.User, error) {
	userJson, err := rdb.Get(context.Background(), index).Result()
	if err != nil {
		return nil, err
	}
	user := model.User{}
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
