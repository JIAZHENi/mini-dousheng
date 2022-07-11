package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {
	// 1.获取参数
	loginId, _ := c.Get("loginId")
	Time := c.Query("latest_time")

	latestTime, _ := strconv.ParseInt(Time, 10, 64)
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	// 2.业务处理
	videoList, nextTime, err := service.Feed(loginId.(int64), latestTime)
	if err != nil {
		model.ResponseFeedError(c)
		return
	}
	data, err := json.Marshal(videoList) //跨包使用，应该保证字段开头大写，不然会缺少字段
	if err != nil {
		log.Printf("序列化错误 err=%v\n", err)
	}
	log.Printf("videoList 序列化后=%v\n", string(data))

	model.ResponseFeedSuccess(c, videoList, nextTime)
}
