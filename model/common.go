package model

import "time"

type User struct {
	Id            int64     `json:"id"`
	UserName      string    `json:"-"`
	Password      string    `json:"-"`
	NickName      string    `gorm:"default:'小创作者'" json:"name,omitempty"`
	FollowCount   int64     `json:"follow_count"`
	FollowerCount int64     `json:"follower_count"`
	CreateTime    time.Time `gorm:"autoCreateTime" json:"-"`
	ModifyTime    time.Time `gorm:"autoUpdateTime" json:"-"`
}

type Video struct {
	Id            int64     `json:"id"`
	Author        int64     `json:"-"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	Title         string    `json:"title,omitempty"`
	CreateTime    time.Time `gorm:"autoCreateTime" json:"-"`
}

type Enjoy struct {
	VideoId    int64
	UserId     int64
	IsFavorite bool      `gorm:"default:1"` // 1 代表 ture, 0 代表 false
	CreateTime time.Time `gorm:"autoCreateTime"`
	ModifyTime time.Time `gorm:"autoUpdateTime"`
}

type Relation struct {
	UserId     int64
	FollowerId int64
	IsFollower bool      `gorm:"default:1"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	ModifyTime time.Time `gorm:"autoUpdateTime"`
}

type Comment struct {
	CommentId  int64     `json:"id"`
	VideoId    int64     `json:"-"`
	UserId     int64     `json:"-"`
	Remark     string    `json:"content,omitempty"`
	ActionType bool      `gorm:"default:1" json:"-"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"-"`
	ModifyTime time.Time `gorm:"autoUpdateTime" json:"-"`
}
