package model

import "time"

type User struct {
	Id            int64
	UserName      string
	Password      string
	NickName      string `gorm:"default:'小创作者'"`
	FollowCount   int64
	FollowerCount int64
	CreateTime    time.Time
	ModifyTime    time.Time
}

type Video struct {
	Id            int64
	Author        int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
	CreateTime    time.Time
}

type Enjoy struct {
	VideoId    int64
	UserId     int64
	IsFavorite bool `gorm:"default:ture"`
	CreateTime time.Time
	ModifyTime time.Time
}

type Relation struct {
	UserId     int64
	FollowerId int64
	IsFollower bool `gorm:"default:ture"`
	CreateTime time.Time
	ModifyTime time.Time
}

type Comment struct {
	CommentId  int64
	VideoId    int64
	UserId     int64
	Remark     string
	ActionType bool `gorm:"default:ture"`
	CreateTime time.Time
	ModifyTime time.Time
}
