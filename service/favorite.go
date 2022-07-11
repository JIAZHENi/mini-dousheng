package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
)

func Favorite(userId int64, videoId int64, action int32) error {
	return db.VideoLikeAction(userId, videoId, action)
}

func FavoriteList(loginId int64, userId int64) ([]model.VideoMessage, error) {
	var videoList []model.Video
	videosID, err := db.VideoLikeList(userId)
	if err != nil {
		return nil, err
	}
	for _, id := range videosID {
		video, _ := db.FindVideo(id)
		videoList = append(videoList, video)
	}
	videosMessage := NewVideoList(videoList, loginId)
	return videosMessage, nil
}
