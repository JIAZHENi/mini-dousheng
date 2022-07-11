package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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
	user := c.Query("user_id")
	userId, _ := strconv.ParseInt(user, 10, 64)

	userMassage, err := service.UserInfo(loginId.(int64), userId)
	if err != nil {
		model.ResponseError(c)
		return
	}
	log.Println(userMassage)

	model.ResponseInfoSuccess(c, userMassage)
}
