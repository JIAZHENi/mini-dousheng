package db

import (
	"Git/mini-dousheng/model"
	"log"
)

func SelectUser(userId int64) (user model.User, err error) {
	err = db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		log.Println("SelectUser error:", err)
		return
	}
	return
}

func CreateUser(u *model.User) (userId int64, err error) {
	err = db.Create(&u).Error
	if err != nil {
		return 0, err
	}
	return u.Id, nil
}

func FindUser(u *model.User) (userId int64, err error) {
	err = db.First(&u).Error
	if err != nil {
		return -1, err
	}
	return u.Id, nil
}
