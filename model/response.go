package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"` // omitempty 表示如果这个字段为空,则返回的 json 数据中没有这个字段
}

func ResponseError(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		StatusCode: 1,
		StatusMsg:  "操作失败",
	})
}

func ResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "操作成功",
	})
}

func ResponseParameterError(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		StatusCode: 1,
		StatusMsg:  "参数错误",
	})
}
