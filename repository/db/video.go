package db

import (
	"Git/mini-dousheng/model"
	"time"
)

func PushVideos(latestTime int64, count int) ([]model.Video, error) {
	var videos []model.Video
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	err := db.Where("create_time<?", timeStr).Order("create_time DESC").Limit(count).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func SaveVideo(userId int64, playURL string, title string) error {
	video := model.Video{
		Author:   userId,
		PlayUrl:  playURL,
		CoverUrl: "http://10.0.2.2:8080/static/cover/0.jpg",
		Title:    title,
	}
	err := db.Save(&video).Error
	if err != nil {
		return err
	}
	return nil
}

func SelectVideo(userId int64) ([]model.Video, error) {
	var videos []model.Video
	err := db.Where("author = ?", userId).Find(&videos).Error
	if err != nil {
		return []model.Video{}, err
	}
	return videos, nil
}
