package db

import (
	"Git/mini-dousheng/model"
	"log"
)

func SelectUser(userId int64) (user model.User, err error) {
	err = db.Where("id = ?", userId).Find(&user).Error
	if err != nil {
		log.Println("SelectUser error:", err)
		return
	}
	return
}
