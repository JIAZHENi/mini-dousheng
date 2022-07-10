package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Feed(c *gin.Context) {
	// 1.获取参数
	loginId, _ := c.Get("loginId")
	Time := c.Query("latest_time")
	latestTime, _ := strconv.ParseInt(Time, 10, 64)
	// 2.业务处理
	videoList, nextTime, err := service.Feed(loginId.(int64), latestTime)
	if err != nil {
		log.Println("Feed error")
		model.VideosResponseError(c)
		return
	}

	// 3.返回响应
	model.VideosResponseSuccess(c, videoList, nextTime)

}
