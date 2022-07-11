package db

import (
	"Git/mini-dousheng/model"
)

func FindUser(userId int64) (model.User, error) {
	var u model.User
	err := db.Where("id = ?", userId).Find(&u).Error
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func CreateUser(userName string, password string) (int64, error) {
	var u model.User
	u.UserName = userName
	u.Password = password
	err := db.Create(&u).Error
	if err != nil {
		return 0, err
	}
	return u.Id, nil
}

func UserIsExist(userName string, password string) (int64, error) {
	var u model.User
	err := db.Where("user_name = ? and password = ?", userName, password).First(&u).Error
	if err != nil {
		return 0, err
	}
	return u.Id, nil
}
