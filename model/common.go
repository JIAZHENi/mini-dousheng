package model

import "time"

type User struct {
	Id            int64     `json:"id,omitempty"`
	UserName      string    `json:"-"`
	Password      string    `json:"-"`
	NickName      string    `gorm:"default:'小创作者'" json:"name,omitempty"`
	FollowCount   int64     `json:"follow_count,omitempty"`
	FollowerCount int64     `json:"follower_count,omitempty"`
	CreateTime    time.Time `gorm:"autoCreateTime" json:"-"`
	ModifyTime    time.Time `gorm:"autoUpdateTime" json:"-"`
}

type Video struct {
	Id            int64     `json:"id,omitempty"`
	Author        int64     `json:"-"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	Title         string    `json:"title,omitempty"`
	CreateTime    time.Time `gorm:"autoCreateTime" json:"-"`
}

type Enjoy struct {
	VideoId    int64
	UserId     int64
	IsFavorite bool      `gorm:"default:ture"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	ModifyTime time.Time `gorm:"autoUpdateTime"`
}

type Relation struct {
	UserId     int64
	FollowerId int64
	IsFollower bool      `gorm:"default:ture"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	ModifyTime time.Time `gorm:"autoUpdateTime"`
}

type Comment struct {
	CommentId  int64     `json:"id,omitempty"`
	VideoId    int64     `json:"-"`
	UserId     int64     `json:"-"`
	Remark     string    `json:"content,omitempty"`
	ActionType bool      `gorm:"default:ture" json:"-"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_date,omitempty"`
	ModifyTime time.Time `gorm:"autoUpdateTime" json:"-"`
}
