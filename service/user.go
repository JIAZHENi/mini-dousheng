package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
	"Git/mini-dousheng/utils/jwt"
)

func UserRegister(userName string, password string) (int64, string, error) {
	userId, err := db.CreateUser(userName, password)
	if err != nil {
		return 0, "", err
	}

	token, err := jwt.CreateToken(userId, userName)
	if err != nil {
		return 0, "", err
	}

	return userId, token, nil
}

func UserLogin(userName string, password string) (int64, string, error) {
	userId, err := db.UserIsExist(userName, password)
	if err != nil {
		return 0, "", err
	}

	token, err := jwt.CreateToken(userId, userName)
	if err != nil {
		return 0, "", err
	}

	return userId, token, nil
}

func UserInfo(loginId int64, userId int64) (model.UserMessage, error) {
	var u model.UserMessage
	user, err := db.FindUser(userId)
	if err != nil {
		return model.UserMessage{}, err
	}
	u, err = NewUserMassage(user, loginId)
	if err != nil {
		return model.UserMessage{}, err
	}
	return u, nil

}
