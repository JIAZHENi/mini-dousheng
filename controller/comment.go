package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CommentAction(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	video := c.Query("video_id")
	comment := c.Query("comment_id")
	actionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(video, 10, 64)
	commentId, _ := strconv.ParseInt(comment, 10, 64)

	if actionType == "1" {
		text := c.Query("comment_text")
		err := service.CommentAdd(loginId.(int64), videoId, commentId, text)
		if err != nil {
			model.ResponseError(c)
			return
		}
	} else {
		err := service.CommentDelete(loginId.(int64), videoId, commentId)
		if err != nil {
			model.ResponseError(c)
			return
		}
	}

	model.ResponseSuccess(c)
}

func CommentList(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	video := c.Query("video_id")
	videoId, _ := strconv.ParseInt(video, 10, 64)
	log.Println(videoId, loginId.(int64), "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	commentList, err := service.SelectComment(videoId, loginId.(int64))
	if err != nil {
		model.ResponseError(c)
		return
	}
	data, err := json.Marshal(commentList) //跨包使用，应该保证字段开头大写，不然会缺少字段
	if err != nil {
		log.Printf("序列化错误 err=%v\n", err)
	}
	log.Printf("videoList 序列化后=%v\n", string(data))

	model.ResponseCommentListSuccess(c, commentList)
}
