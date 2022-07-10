package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	// 1.获取参数
	loginId, _ := c.Get("loginId")
	videoTitle, _ := c.GetPostForm("title")
	video, err := c.FormFile("data")
	if err != nil {
		model.ResponseSuccess(c)
		return
	}

	// 2.业务逻辑
	err = service.Publish(video, loginId.(int64), videoTitle)

	// 3.响应请求
	if err != nil {
		model.ResponseError(c)
		return
	}
	model.ResponseSuccess(c)

}
