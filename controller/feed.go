package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {
	// 1.获取参数
	loginId, _ := c.Get("loginId")
	TimeStr := c.Query("latest_time")
	log.Println(TimeStr, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	Time32, _ := strconv.Atoi(TimeStr)
	latestTime := int64(Time32)
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	// 2.业务处理
	videoList, nextTime, err := service.Feed(loginId.(int64), latestTime)
	if err != nil {
		model.ResponseFeedError(c)
		return
	}
	log.Println(videoList)

	model.ResponseFeedSuccess(c, videoList, nextTime)
}
