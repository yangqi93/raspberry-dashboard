package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/yangqi93/raspberry-dashboard/config"
	_ "github.com/yangqi93/raspberry-dashboard/log"
	"github.com/yangqi93/raspberry-dashboard/router"
)

func main() {
	engin := gin.Default()

	//路由初始化
	engin.Static("/assets", "assets")
	router.Init(engin)
	err := engin.Run(":9001")
	if err != nil {
		panic(err)
	}

}
