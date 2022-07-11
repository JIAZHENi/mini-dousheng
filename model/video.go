package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoMessage struct {
	Video
	Author     UserMessage `json:"author,omitempty"`
	IsFavorite bool        `json:"is_favorite,omitempty"`
}

type FeedResponse struct {
	Response
	NextTime int64          `json:"next_time"`
	Videos   []VideoMessage `json:"video_list"`
}

type PublishListResponse struct {
	Response
	Videos []VideoMessage `json:"video_list"`
}

func ResponseFeedSuccess(c *gin.Context, videoList []VideoMessage, nextTime int64) {
	c.JSON(http.StatusOK, &FeedResponse{
		Response: Response{StatusCode: 0, StatusMsg: "操作成功"},
		NextTime: nextTime,
		Videos:   videoList,
	})
}

func ResponseFeedError(c *gin.Context) {
	c.JSON(http.StatusOK, &FeedResponse{
		Response: Response{StatusCode: 1, StatusMsg: "操作失败"},
		NextTime: 0,
		Videos:   nil,
	})
}

func ResponsePublishListSuccess(c *gin.Context, videoList []VideoMessage) {
	c.JSON(http.StatusOK, &PublishListResponse{
		Response: Response{StatusCode: 0, StatusMsg: "操作成功"},
		Videos:   videoList,
	})
}

func ResponseFavoriteListSuccess(c *gin.Context, videoList []VideoMessage) {
	c.JSON(http.StatusOK, &PublishListResponse{
		Response: Response{StatusCode: 0, StatusMsg: "操作成功"},
		Videos:   videoList,
	})
}
