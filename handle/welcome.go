package handle

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/gin-gonic/gin"
	"github.com/yangqi93/raspberry-dashboard/config"
	"html/template"
)

type WelcomeRequest struct {
	Page     int32 `form:"page" validate:"required"`
	PageSize int32 `form:"pageSize" validate:"required"`
}

func Welcome(c *gin.Context) {
	if ajax, ok := c.GetQuery("ajax"); ok && ajax == "true" {
		c.JSON(200, GetInfo())
		return
	}

	TemplateFiles = append(TemplateFiles, "templates/welcome/welcome.tmpl")

	t, err := template.New("test").Funcs(sprig.FuncMap()).ParseFiles(
		TemplateFiles...,
	)
	if err != nil {
		panic(err)
		//c.HTML(500, "error.tmpl", gin.H{"error": err.Error()})
	}

	info := GetInfo()
	err = t.ExecuteTemplate(c.Writer, "layout", gin.H{
		"title":    "Welcome",
		"piModel":  info.Cpu.PiModel,
		"ip":       info.LocalIp,
		"user":     info.UserName,
		"os":       info.Os,
		"hostName": config.HostName,
		"uname":    info.Uname,
		"net":      info.Net,
		"info":     info,
	})
	if err != nil {
		panic(err)
	}
}
