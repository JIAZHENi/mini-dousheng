package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
	"strconv"
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
	user := c.Query("user_id")
	userId, _ := strconv.ParseInt(user, 10, 64)

	videoList, err := service.GetPublishList(loginId.(int64), userId)
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponsePublishListSuccess(c, videoList)

}
