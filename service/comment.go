package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
)

func CommentAdd(userId int64, videoId int64, commentId int64, text string) error {
	return db.CommentAdd(userId, videoId, commentId, text)
}

func CommentDelete(userId int64, videoId int64, commentId int64) error {
	return db.CommentDelete(videoId, userId, commentId)
}

func SelectComment(video int64, loginId int64) ([]model.CommentMassage, error) {
	var commentMassages []model.CommentMassage
	commentList, err := db.SelectComment(video)
	if err != nil {
		return nil, err
	}
	commentMassages = NewCommentMassage(commentList, loginId)
	return commentMassages, nil
}
