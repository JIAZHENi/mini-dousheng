package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	var v model.FavoriteRequest
	if err := c.ShouldBind(&v); err != nil {
		model.ResponseError(c)
		return
	}

	err := service.Favorite(loginId.(int64), v.VideoId, v.ActionType)
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponseSuccess(c)

}

func FavoriteList(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	user := c.Query("user_id")
	userId, _ := strconv.ParseInt(user, 10, 64)

	videoList, err := service.FavoriteList(loginId.(int64), userId)
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponseFavoriteListSuccess(c, videoList)

}
