package router

import (
	"github.com/gin-gonic/gin"
	"raspberry-dashboard/handle"
	"raspberry-dashboard/middleware"
)

func Init(r *gin.Engine) {

	r.Use(gin.Recovery(), middleware.RequestLog())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/v1/welcome")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/welcome", handle.Welcome)
	}
}
