package db

import (
	"Git/mini-dousheng/model"
	"errors"
	"gorm.io/gorm"
)

func Attention(loginId int64, toUserId int64) error {
	if loginId == toUserId {
		return errors.New("自己不能关注自己")
	}
	var r model.Relation
	var u model.User
	tx := db.Begin()
	if IsExistRelation(loginId, toUserId) {
		err := tx.Model(&r).Where("user_id = ? and follower_id = ?", toUserId, loginId).Update("is_favorite", 1).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		r.UserId = toUserId
		r.FollowerId = loginId
		err := tx.Create(&r).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// 更新用户表
	// 登录用户的关注数加一
	err := tx.Model(&u).Where("id = ?", loginId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 被登录用户的关注的用户粉丝数加一
	err = tx.Model(&u).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func UnAttention(loginId int64, toUserId int64) error {
	var r model.Relation
	var u model.User
	tx := db.Begin()
	if IsExistRelation(loginId, toUserId) {
		err := tx.Model(&r).Where("user_id = ? and follower_id = ?", toUserId, loginId).Update("is_favorite", 0).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		return errors.New("UnAttention error")
	}
	// 更新用户表
	// 登录用户的关注数减一
	err := tx.Model(&u).Where("id = ?", loginId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 被登录用户的关注的用户粉丝数减一
	err = tx.Model(&u).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func IsExistRelation(loginId int64, toUserId int64) bool {
	var r model.Relation
	err := db.Where("user_id = ? and follower_id = ?", toUserId, loginId).First(&r).Error
	if err != nil {
		return false
	}
	return true
}

func SelectFollow(userId int64) ([]int64, error) {
	var r []model.Relation
	var followsId []int64
	err := db.Select("user_id").Where("follower_id = ? and is_favorite = ?", userId, 1).Find(&r).Error
	if err != nil {
		return nil, err
	}
	for _, v := range r {
		followsId = append(followsId, v.UserId)
	}
	return followsId, nil
}

func SelectFollower(userId int64) ([]int64, error) {
	var r []model.Relation
	var followerId []int64
	err := db.Select("follower_id").Where("user_id = ? and is_favorite = ?", userId, 1).Find(&r).Error
	if err != nil {
		return nil, err
	}
	for _, v := range r {
		followerId = append(followerId, v.FollowerId)
	}
	return followerId, nil
}

func IsAttention(userId int64, followerId int64) bool {
	if !IsExistRelation(followerId, userId) {
		return false
	}
	var r model.Relation
	err := db.Where("user_id = ? and follower_id = ? and is_favorite = ?", userId, followerId, 1).First(&r).Error
	if err != nil {
		return false
	}
	return true
}
