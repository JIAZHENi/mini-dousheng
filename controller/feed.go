package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Feed(c *gin.Context) {
	// 1.获取参数
	loginId, _ := c.Get("loginId")
	var p model.LatestTimeRequest
	if err := c.ShouldBind(&p); err != nil {
		model.ResponseParameterError(c)
		return
	}
	if p.LatestTime == 0 {
		p.LatestTime = time.Now().Unix()
	}
	// 2.业务处理
	videoList, nextTime, err := service.Feed(loginId.(int64), p.LatestTime)
	if err != nil {
		model.ResponseFeedError(c)
		return
	}
	log.Println(videoList)

	model.ResponseFeedSuccess(c, videoList, nextTime)
}
