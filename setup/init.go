package setup

import (
	"Git/mini-dousheng/repository/db"
	"Git/mini-dousheng/routers"
	"fmt"
)

func Init() error {
	// 1
	if err := db.Init(); err != nil {
		fmt.Println("init db err:", err)
		return err
	}
	// 2
	r := routers.SetupRouter("release")
	if err := r.Run(); err != nil {
		fmt.Println("SetupRouter err:", err)
		return err
	}

	return nil
}
