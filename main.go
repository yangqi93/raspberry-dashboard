package main

import (
	"github.com/gin-gonic/gin"
	_ "zhu/config"
	_ "zhu/log"
	"zhu/router"
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
