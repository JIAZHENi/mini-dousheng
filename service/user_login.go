package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
	"Git/mini-dousheng/utils/jwt"
	"log"
)

func Register(username string, password string) (userId int64, token string, err error) {
	userId, err = db.CreateUser(&model.User{UserName: username, Password: password})
	if err != nil {
		log.Println("创建用户失败")
		return
	}
	token, err = jwt.CreateToken(userId, username)
	if err != nil {
		log.Println("创建 token 失败")
		return
	}

	return

}

func Login(username string, password string) (userId int64, token string, err error) {
	userId, err = db.FindUser(&model.User{UserName: username, Password: password})
	if err != nil {
		log.Println("登录失败")
		return
	}
	token, err = jwt.CreateToken(userId, username)
	if err != nil {
		log.Println("登录失败")
		return
	}

	return

}
