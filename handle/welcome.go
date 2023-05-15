package handle

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/gin-gonic/gin"
	"github.com/yangqi93/raspberry-dashboard/config"
	"html/template"
	"net"
	"os"
	"os/exec"
	"os/user"
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
	userName, unameR := "N/A", "N/A"
	u, err := user.Current()
	if err != nil {
		userName = u.Username
	}
	o := exec.Command("uname")
	osR, _ := o.CombinedOutput()
	uname, err := os.ReadFile("/proc/version")
	if err == nil {
		unameR = string(uname)
	}
	err = t.ExecuteTemplate(c.Writer, "layout", gin.H{
		"title":    "Welcome",
		"piModel":  "Raspberry Pi",
		"ip":       GetLocalIP(),
		"user":     userName,
		"os":       string(osR),
		"hostName": config.Conf.Value.GetString("hostName"),
		"uname":    unameR,
		"net":      info.Net,
		"info":     info,
	})
	if err != nil {
		panic(err)
	}
}

func GetLocalIP() string {
	adders, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range adders {
			if ipNet, ok := addr.(*net.IPNet); ok &&
				!ipNet.IP.IsLinkLocalUnicast() && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
