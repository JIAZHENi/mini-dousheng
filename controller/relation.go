package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	var p model.RelationActionRequest
	if err := c.ShouldBind(&p); err != nil {
		model.ResponseParameterError(c)
		return
	}
	err := service.RelationAction(loginId.(int64), p.ToUserId, p.ActionType)
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponseSuccess(c)
}

func FollowList(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	var p model.UserIdRequest
	if err := c.ShouldBind(&p); err != nil {
		model.ResponseParameterError(c)
		return
	}

	followList, err := service.FollowList(p.UserId, loginId.(int64))
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponseFollowSuccess(c, followList)

}

func FollowerList(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	var p model.UserIdRequest
	if err := c.ShouldBind(&p); err != nil {
		model.ResponseParameterError(c)
		return
	}

	followerList, err := service.FollowerList(p.UserId, loginId.(int64))
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponseFollowSuccess(c, followerList)

}
