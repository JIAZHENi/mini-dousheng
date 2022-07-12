package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
)

func RelationAction(loginId int64, userId int64, actionType int32) error {
	if actionType == 1 {
		err := db.Attention(loginId, userId)
		if err != nil {
			return err
		}
	} else {
		err := db.UnAttention(loginId, userId)
		if err != nil {
			return err
		}
	}
	return nil
}

func FollowList(userId int64, loginId int64) ([]model.UserMessage, error) {
	var followMassage []model.UserMessage
	followsId, err := db.SelectFollow(userId)
	if err != nil {
		return nil, err
	}
	for _, v := range followsId {
		user, _ := db.FindUser(v)
		follows, _ := NewUserMassage(user, loginId)
		followMassage = append(followMassage, follows)
	}
	return followMassage, nil
}

func FollowerList(userId int64, loginId int64) ([]model.UserMessage, error) {
	var followerMassage []model.UserMessage
	followersId, err := db.SelectFollower(userId)
	if err != nil {
		return nil, err
	}
	for _, v := range followersId {
		user, _ := db.FindUser(v)
		followers, _ := NewUserMassage(user, loginId)
		followerMassage = append(followerMassage, followers)
	}
	return followerMassage, nil
}
