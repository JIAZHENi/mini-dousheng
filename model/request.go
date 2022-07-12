package model

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

type CommentActionRequest struct {
	VideoId    int64  `form:"video_id" json:"video_id" binding:"required"`
	CommentId  int64  `form:"comment_id" json:"comment_id" binding:"required"`
	ActionType string `form:"action_type" json:"action_type" binding:"required"`
}

type RelationActionRequest struct {
	ToUserId   int64 `form:"to_user_id" json:"to_user_id" binding:"required"`
	ActionType int32 `form:"action_type" json:"action_type" binding:"required"`
}
