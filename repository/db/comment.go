package db

import (
	"Git/mini-dousheng/model"
	"gorm.io/gorm"
)

func CommentAdd(videoId int64, userId int64, commentId int64, text string) error {
	var c model.Comment
	var v model.Video
	c.VideoId = videoId
	c.UserId = userId
	c.CommentId = commentId
	c.Remark = text
	tx := db.Begin()
	err := tx.Create(&c).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&v).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func CommentDelete(videoId int64, userId int64, commentId int64) error {
	var c model.Comment
	var v model.Video
	tx := db.Begin()
	err := tx.Model(&c).Where("video_id = ? and user_id = ? and comment_id = ?", videoId, userId, commentId).Update("action_type", 0).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&v).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func SelectComment(videoId int64) ([]model.Comment, error) {
	var c []model.Comment
	err := db.Where("video_id", videoId).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
