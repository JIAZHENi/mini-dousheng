package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type VideoMessage struct {
	Id            int64       `json:"id,omitempty"`
	Author        UserMessage `json:"author,omitempty"`
	PlayUrl       string      `json:"play_url,omitempty"`
	CoverUrl      string      `json:"cover_url,omitempty"`
	FavoriteCount int64       `json:"favorite_count,omitempty"`
	CommentCount  int64       `json:"comment_count,omitempty"`
	Title         string      `json:"title,omitempty"`
	IsFavorite    bool        `json:"is_favorite,omitempty"`
}

type FeedResponse struct {
	Response
	NextTime  int64          `json:"next_time"`
	VideoList []VideoMessage `json:"video_list"`
}

func VideosResponseError(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "视频拉取错误"},
		VideoList: nil,
		NextTime:  time.Now().Unix(),
	})
}

func VideosResponseSuccess(c *gin.Context, videoList []VideoMessage, nextTime int64) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "Success"},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})

}
