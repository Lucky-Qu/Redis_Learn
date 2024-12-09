package dao

import "Redis_Learn/03_UsingRedisAsCache/db/model"

func GetUser(id int) (*model.User, error) {
	var user model.User
	if err := db.Model(&model.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func CreateUser(user *model.User) error {
	if err := db.Model(&model.User{}).Create(user).Error; err != nil {
		return err
	}
	return nil
}
