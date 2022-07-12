package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	var p model.VideoPublishRequest
	if err := c.ShouldBind(&p); err != nil {
		model.ResponseParameterError(c)
		return
	}
	err := service.PublishVideo(p.VideoFile, loginId.(int64), p.VideoTitle)
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponseSuccess(c)
}

func PublishList(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	var p model.UserIdRequest
	if err := c.ShouldBind(&p); err != nil {
		model.ResponseParameterError(c)
		return
	}

	videoList, err := service.GetPublishList(loginId.(int64), p.UserId)
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponsePublishListSuccess(c, videoList)

}
