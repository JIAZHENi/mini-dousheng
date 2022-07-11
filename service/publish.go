package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

func PublishVideo(data *multipart.FileHeader, loginId int64, title string) error {
	nowTime := time.Now().Format("20060102150405")
	// 文件路径，包含文件名字：用户ID_发表时间_视频名字 防止文件名冲突
	videoPath := fmt.Sprintf("public/video/%d_%s_%s", loginId, nowTime, title)
	urlPath := fmt.Sprintf("http://10.0.2.2:8080/static/video/%d_%s_%s", loginId, nowTime, title)

	err := func(file *multipart.FileHeader, dst string) error {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		out, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, src)
		return err
	}(data, videoPath)
	if err != nil {
		return err
	}

	err = db.SaveVideo(loginId, urlPath, title)
	if err != nil {
		return err
	}
	return nil
}

func GetPublishList(loginId int64, userId int64) ([]model.VideoMessage, error) {
	var videoList []model.VideoMessage
	videos, err := db.SelectVideo(userId)
	if err != nil {
		return nil, err
	}
	videoList = NewVideoList(videos, loginId)
	return videoList, nil
}
