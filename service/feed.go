package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
	"time"
)

func Feed(loginId int64, latest int64) ([]model.VideoMessage, int64, error) {
	var videoList []model.VideoMessage
	videos, err := db.PushVideos(latest, 4)
	if err != nil {
		return nil, 0, err
	}
	videoList = NewVideoList(videos, loginId)
	return videoList, time.Now().Unix(), nil
}
