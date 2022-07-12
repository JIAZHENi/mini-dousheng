package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	var p model.CommentActionRequest
	if err := c.ShouldBind(&p); err != nil {
		model.ResponseParameterError(c)
		return
	}

	err := service.CommentAction(loginId.(int64), p)
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponseSuccess(c)
}

func CommentList(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	var p model.VideoRequest
	if err := c.ShouldBind(&p); err != nil {
		model.ResponseParameterError(c)
		return
	}

	commentList, err := service.SelectComment(p.Video, loginId.(int64))
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponseCommentListSuccess(c, commentList)
}
