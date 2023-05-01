package handle

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/gin-gonic/gin"
	"html/template"
	"net"
	"os"
	"os/exec"
	"os/user"
	"raspberry-dashboard/config"
)

type WelcomeRequest struct {
	Page     int32 `form:"page" validate:"required"`
	PageSize int32 `form:"pageSize" validate:"required"`
}

type Status struct {
	Page struct {
		Time struct {
			Start []string `json:"start"`
		} `json:"time"`
	} `json:"page"`
	Time   int64  `json:"time"`
	Uptime string `json:"uptime"`
	Cpu    struct {
		Stat struct {
			User    string `json:"user"`
			Nice    string `json:"nice"`
			Sys     string `json:"sys"`
			Idle    string `json:"idle"`
			Iowait  string `json:"iowait"`
			Irq     string `json:"irq"`
			Softirq string `json:"softirq"`
		} `json:"stat"`
		Freq    int      `json:"freq"`
		Temp    []string `json:"temp"`
		Count   int      `json:"count"`
		Model   string   `json:"model"`
		PiModel string   `json:"pi_model"`
	} `json:"cpu"`
	Mem struct {
		Total         float64 `json:"total"`
		Free          float64 `json:"free"`
		Buffers       float64 `json:"buffers"`
		Cached        float64 `json:"cached"`
		CachedPercent float64 `json:"cached_percent"`
		Used          float64 `json:"used"`
		Percent       float64 `json:"percent"`
		Real          struct {
			Used    float64 `json:"used"`
			Free    float64 `json:"free"`
			Percent float64 `json:"percent"`
		} `json:"real"`
		Swap struct {
			Total   int `json:"total"`
			Free    int `json:"free"`
			Used    int `json:"used"`
			Percent int `json:"percent"`
		} `json:"swap"`
	} `json:"mem"`
	LoadAvg []string `json:"load_avg"`
	Disk    struct {
		Total   float64 `json:"total"`
		Free    float64 `json:"free"`
		Used    float64 `json:"used"`
		Percent float64 `json:"percent"`
	} `json:"disk"`
	Net struct {
		Count      int         `json:"count"`
		Interfaces []Interface `json:"interfaces"`
	} `json:"net"`
	Version string `json:"version"`
}

type Interface struct {
	Name     string `json:"name"`
	TotalIn  string `json:"total_in"`
	TotalOut string `json:"total_out"`
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
