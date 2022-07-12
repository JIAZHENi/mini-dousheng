package db

import (
	"Git/mini-dousheng/model"
	"gorm.io/gorm"
)

func VideoLikeAction(userId int64, videoId int64, action int32) error {
	var u model.Enjoy
	var v model.Video
	tx := db.Begin()
	// 点赞关系表的操作
	if IsLikeAction(userId, videoId) {
		err := tx.Model(&u).Where("user_id = ? and video_id = ?", userId, videoId).Update("is_favorite", action).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		u.UserId = userId
		u.VideoId = videoId
		err := tx.Create(&u).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// 视频点赞数量的更新
	if action == 1 {
		err := tx.Model(&v).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		err := tx.Model(&v).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func IsLikeAction(userId int64, videoId int64) bool {
	var u model.Enjoy
	err := db.Where("user_id = ? and video_id = ?", userId, videoId).First(&u).Error
	if err != nil {
		return false
	}
	return true
}

func IsLikeVideo(userId int64, videoId int64) bool {
	var u model.Enjoy
	err := db.Where("video_id = ? and user_id = ? and is_favorite = ?", videoId, userId, 1).First(&u).Error
	if err != nil {
		return false
	}
	return true
}

func VideoLikeList(userId int64) ([]int64, error) {
	var E []model.Enjoy
	var Videos []int64
	err := db.Select("video_id").Where("user_id = ? and is_favorite = ?", userId, 1).Find(&E).Error
	if err != nil {
		return nil, err
	}
	for _, v := range E {
		Videos = append(Videos, v.VideoId)
	}
	return Videos, nil
}
