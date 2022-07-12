package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	var v model.FavoriteRequest
	if err := c.ShouldBind(&v); err != nil {
		model.ResponseParameterError(c)
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
	var p model.UserIdRequest
	if err := c.ShouldBind(&p); err != nil {
		model.ResponseParameterError(c)
		return
	}

	videoList, err := service.FavoriteList(loginId.(int64), p.UserId)
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponseFavoriteListSuccess(c, videoList)

}
