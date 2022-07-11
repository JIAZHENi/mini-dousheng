package setup

import (
	"Git/mini-dousheng/repository/db"
	"Git/mini-dousheng/routers"
	"log"
)

func Init() error {
	// 1
	if err := db.Init(); err != nil {
		log.Println("Init db err:", err)
		return err
	}
	// 2
	r := routers.SetupRouter("release")
	if err := r.Run(); err != nil {
		log.Println("SetupRouter err:", err)
		return err
	}

	return nil
}
