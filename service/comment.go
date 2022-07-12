package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
)

func CommentAction(userId int64, p model.CommentActionRequest) error {
	if p.ActionType == "1" {
		return db.CommentAdd(p.VideoId, userId, p.CommentId, p.CommentText)
	}

	return db.CommentDelete(p.VideoId, userId, p.CommentId)

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
