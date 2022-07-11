package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
)

//func NewUserMassage(user model.User, loginId int64) (model.UserMessage, error) {
//	return model.UserMessage{
//		Id:            user.Id,
//		NickName:      user.NickName,
//		FollowCount:   user.FollowCount,
//		FollowerCount: user.FollowerCount,
//		IsFollow:      false,
//	}, nil
//}

func NewUserMassage(user model.User, loginId int64) (model.UserMessage, error) {
	return model.UserMessage{
		User:     user,
		IsFollow: false,
	}, nil
}

func NewVideoList(videos []model.Video, loginId int64) []model.VideoMessage {
	var videoList []model.VideoMessage
	for _, video := range videos {
		author, _ := db.FindUser(video.Author)
		authorMassage, _ := NewUserMassage(author, loginId)
		videoMassage := model.VideoMessage{
			Video:      video,
			Author:     authorMassage,
			IsFavorite: false,
		}
		videoList = append(videoList, videoMassage)
	}
	return videoList
}
