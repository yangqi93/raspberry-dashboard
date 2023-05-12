package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yangqi93/raspberry-dashboard/log"
)

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		request := c.Request.URL.Query()
		c.Next()
		// after request body
		response := "aaaa"

		info := map[string]interface{}{"url": c.Request.URL, "method": c.Request.Method, "request": request, "response": response}
		infoJson, _ := json.Marshal(info)
		log.Log.Info("request info", string(infoJson))
	}
}
