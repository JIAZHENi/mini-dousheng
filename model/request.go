package model

type UserRequest struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type FavoriteRequest struct {
	VideoId    int64 `form:"video_id" json:"video_id" binding:"required"`
	ActionType int32 `form:"action_type" json:"action_type" binding:"required"`
}
