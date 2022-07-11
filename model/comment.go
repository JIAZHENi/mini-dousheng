package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentMassage struct {
	Comment
	UserMessage `json:"user,omitempty"`
	CreateTime  string `json:"create_date,omitempty"`
}

type ResponseComment struct {
	Response
	CommentList []CommentMassage `json:"comment_list,omitempty"`
}

func ResponseCommentListSuccess(c *gin.Context, u []CommentMassage) {
	c.JSON(http.StatusOK, &ResponseComment{
		Response:    Response{StatusCode: 0, StatusMsg: "操作成功"},
		CommentList: u,
	})
}
