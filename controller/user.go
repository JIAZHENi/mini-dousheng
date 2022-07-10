package controller

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	// 1.获取参数
	var user model.UserLogin
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User Register err"},
		})
		return
	}

	// 2.业务逻辑
	userId, token, err := service.Register(user.UserName, user.Password)

	// 3.响应
	if err != nil {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "注册失败"},
		})
	}

	c.JSON(http.StatusOK, model.UserLoginResponse{
		UserId: userId,
		Token:  token,
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
	})
}

func Login(c *gin.Context) {
	// 1.获取参数
	var user model.UserLogin
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User login err"},
		})
		return
	}
	// 2.业务处理
	userId, token, err := service.Login(user.UserName, user.Password)

	// 3.响应
	if err != nil {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "登录失败"},
		})
	}

	c.JSON(http.StatusOK, model.UserLoginResponse{
		UserId: userId,
		Token:  token,
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
	})
}

func UserInfo(c *gin.Context) {
	loginId, _ := c.Get("loginId")
	id := c.Query("user_id")
	userId, _ := strconv.ParseInt(id, 10, 64)

	UMessage, err := service.NewMassageUser(loginId.(int64), userId)
	if err != nil {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "获取信息失败"},
		})

	}

	c.JSON(http.StatusOK, model.UserResponse{
		Response: model.Response{StatusCode: 0},
		User:     UMessage,
	})
}
