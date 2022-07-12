package model

import "mime/multipart"

type UserRequest struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type FavoriteRequest struct {
	VideoId    int64 `form:"video_id" json:"video_id" binding:"required"`
	ActionType int32 `form:"action_type" json:"action_type" binding:"required"`
}

type UserIdRequest struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required"`
}

type VideoRequest struct {
	Video int64 `form:"video_id" json:"video_id" binding:"required"`
}

type LatestTimeRequest struct {
	LatestTime int64 `form:"latest_time" json:"latest_time"`
}

type VideoPublishRequest struct {
	VideoFile  *multipart.FileHeader `form:"data" json:"data" binding:"required"`
	VideoTitle string                `form:"title" json:"title" binding:"required"`
}

type CommentActionRequest struct {
	VideoId     int64  `form:"video_id" json:"video_id" binding:"required"`
	CommentId   int64  `form:"comment_id" json:"comment_id"`
	ActionType  string `form:"action_type" json:"action_type" binding:"required"`
	CommentText string `form:"comment_text" json:"comment_text"`
}

type RelationActionRequest struct {
	ToUserId   int64 `form:"to_user_id" json:"to_user_id" binding:"required"`
	ActionType int32 `form:"action_type" json:"action_type" binding:"required"`
}
