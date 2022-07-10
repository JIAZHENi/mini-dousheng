package db

import (
	"Git/mini-dousheng/model"
)

func SelectVideos(latestTime int64) (videoList []model.Video, err error) {
	err = db.Where("create_time<?", latestTime).Order("create_time DESC").Limit(5).Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	return
}
