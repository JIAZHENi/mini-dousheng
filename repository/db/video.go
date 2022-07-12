package db

import (
	"Git/mini-dousheng/model"
	"errors"
	"time"
)

func PushVideos(latestTime int64, count int) ([]model.Video, int64, error) {
	var videos []model.Video
	timeStr := time.UnixMilli(latestTime).Format("2006-01-02 15:04:05")

	// 取比当前时间小的,也就是以前的视频，然后视频按大到小排序，先呈现最新的视频
	err := db.Where("create_time <= ?", timeStr).Order("create_time DESC").Limit(count).Find(&videos).Error
	if err != nil {
		return nil, 0, err
	}
	if len(videos) == 0 {
		return nil, 0, errors.New("PushVideo err")
	}
	// 把最新视频的时间
	// return videos, videos[0].CreateTime.UnixMilli(), nil
	return videos, time.Now().UnixMilli(), nil
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

func FindVideo(videoId int64) (model.Video, error) {
	var videos model.Video
	err := db.Where("id = ?", videoId).Find(&videos).Error
	if err != nil {
		return model.Video{}, err
	}
	return videos, nil
}
