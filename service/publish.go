package service

import (
	"Git/mini-dousheng/model"
	"Git/mini-dousheng/repository/db"
	"fmt"
	"mime/multipart"
	"time"
)

func Publish(data *multipart.FileHeader, userId int64, title string) (err error) {

	//pwd, _ := os.Getwd() // 项目根路径
	nowTime := time.Now().Format("20060102150405")
	// 文件路径，包含文件名字：用户ID_发表时间_视频名字 防止文件名冲突
	// videoPath := fmt.Sprintf("%s/public/video/%s_%s_%s", pwd, loginIDStr, nowTime, video.Filename)
	videoPath := fmt.Sprintf("public/video/%d_%s_%s", userId, nowTime, title)
	urlPath := fmt.Sprintf("http://10.0.2.2:8080/static/video/%d_%s_%s", userId, nowTime, title)

	err = SaveVideo(data, videoPath)
	if err != nil {
		return err
	}

	video := model.Video{
		Author:   userId,
		Title:    title,
		PlayUrl:  urlPath,
		CoverUrl: "http://10.0.2.2:8080/static/cover/0.jpg",
	}
	err = db.SaveVideo(&video)

	return

}
