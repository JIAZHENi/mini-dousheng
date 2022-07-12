package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var u model.UserRequest
	if err := c.ShouldBind(&u); err != nil {
		model.ResponseError(c)
		return
	}
	userId, token, err := service.UserRegister(u.UserName, u.Password)
	if err != nil {
		model.ResponseUserError(c)
		return
	}
	model.ResponseUserSuccess(c, userId, token)
}

func Login(c *gin.Context) {
	var u model.UserRequest
	if err := c.ShouldBind(&u); err != nil {
		model.ResponseError(c)
		return
	}
	userId, token, err := service.UserLogin(u.UserName, u.Password)
	if err != nil {
		model.ResponseUserError(c)
		return
	}
	model.ResponseUserSuccess(c, userId, token)
}

func UserInfo(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	var p model.UserIdRequest
	if err := c.ShouldBind(&p); err != nil {
		model.ResponseError(c)
		return
	}

	userMassage, err := service.UserInfo(loginId.(int64), p.UserId)
	if err != nil {
		model.ResponseError(c)
		return
	}

	model.ResponseInfoSuccess(c, userMassage)
}
