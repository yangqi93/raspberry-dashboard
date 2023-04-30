package handle

import (
	"github.com/gin-gonic/gin"
	"html/template"
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
		Freq int32    `json:"freq"`
		Temp []string `json:"temp"`
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

	t, err := template.New("test").Funcs(template.FuncMap{
		"abc": func(x int) bool {
			return x == 0 || (x+1)%4 == 0
		},
		"cde": func(x int) bool {
			return x != 0 && (x+1)%4 == 0
		},
	}).ParseFiles(
		TemplateFiles...,
	)
	if err != nil {
		panic(err)
		//c.HTML(500, "error.tmpl", gin.H{"error": err.Error()})
	}

	info := GetInfo()
	err = t.ExecuteTemplate(c.Writer, "layout", gin.H{
		"title":    "Welcome",
		"piModel":  "aa",
		"ip":       "hostip",
		"user":     "root",
		"os":       "os",
		"hostName": "host name",
		"uname":    "uname",
		"net":      info.Net,
	})
	if err != nil {
		panic(err)
	}
}
