package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
)

func Feed(loginId int64, latest int64) ([]model.VideoMessage, int64, error) {
	var videoList []model.VideoMessage
	videos, nextTime, err := db.PushVideos(latest, 10)
	if err != nil {
		return nil, 0, err
	}
	videoList = NewVideoList(videos, loginId)
	return videoList, nextTime, nil
}
