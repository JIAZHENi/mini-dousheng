package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
)

func Feed(loginId int64, latestTime int64) (videoMessageList []model.VideoMessage, nextTime int64, err error) {
	var videoList []model.Video
	videoList, err = db.SelectVideos(latestTime)

	for _, video := range videoList {
		authorMassage, _ := NewMassageUser(video.Author, loginId)
		VideoMessage := model.VideoMessage{
			Id:            video.Id,
			Author:        authorMassage,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			IsFavorite:    false,
		}
		videoMessageList = append(videoMessageList, VideoMessage)
	}
	return
}
