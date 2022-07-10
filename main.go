package main

import (
	"Git/mini-dousheng/setup"
	"log"
)

func main() {
	if err := setup.Init(); err != nil {
		log.Println("setup err:", err)
	}
}
