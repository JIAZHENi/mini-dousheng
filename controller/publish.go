package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	videoTitle := c.PostForm("title")
	video, err := c.FormFile("data")
	if err != nil {
		model.ResponseError(c)
		return
	}
	err = service.PublishVideo(video, loginId.(int64), videoTitle)
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
		model.ResponseError(c)
		return
	}

	videoList, err := service.GetPublishList(loginId.(int64), p.UserId)
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponsePublishListSuccess(c, videoList)

}
