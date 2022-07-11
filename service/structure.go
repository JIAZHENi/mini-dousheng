package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
)

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
			IsFavorite: db.IsLikeVideo(loginId, video.Id),
		}
		videoList = append(videoList, videoMassage)
	}
	return videoList
}

func NewCommentMassage(comments []model.Comment, loginId int64) []model.CommentMassage {
	var commentList []model.CommentMassage
	for _, comment := range comments {
		user, _ := db.FindUser(comment.UserId)
		userMessage, _ := NewUserMassage(user, loginId)
		commentMassage := model.CommentMassage{
			Comment:     comment,
			UserMessage: userMessage,
			CreateTime:  comment.CreateTime.Format("2006-01-02 15:04:05"),
		}
		commentList = append(commentList, commentMassage)
	}
	return commentList
}
