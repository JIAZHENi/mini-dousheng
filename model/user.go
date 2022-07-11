package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserMessage struct {
	User
	IsFollow bool `json:"is_follow"`
}

type UserResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	UserMessage `json:"user,omitempty"`
}

func ResponseUserSuccess(c *gin.Context, userId int64, token string) {
	c.JSON(http.StatusOK, &UserResponse{
		Response: Response{StatusCode: 0, StatusMsg: "操作成功"},
		UserId:   userId,
		Token:    token,
	})
}

func ResponseUserError(c *gin.Context) {
	c.JSON(http.StatusOK, &UserResponse{
		Response: Response{StatusCode: 1, StatusMsg: "操作失败"},
	})
}

func ResponseInfoSuccess(c *gin.Context, u UserMessage) {
	c.JSON(http.StatusOK, &UserInfoResponse{
		Response:    Response{StatusCode: 0, StatusMsg: "操作成功"},
		UserMessage: u,
	})
}
