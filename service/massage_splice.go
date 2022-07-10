package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
)

func NewMassageUser(userId int64, loginId int64) (model.UserMessage, error) {
	user, err := db.SelectUser(userId)

	return model.UserMessage{
		Id:            user.Id,
		Name:          user.NickName,
		FollowerCount: user.FollowerCount,
		FollowCount:   user.FollowCount,
		IsFollow:      false,
	}, err

}
