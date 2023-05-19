package service

type Status struct {
	Page    Page     `json:"page"`
	Time    int64    `json:"time"`
	Uptime  string   `json:"uptime"`
	Cpu     Cpu      `json:"cpu"`
	Mem     Mem      `json:"mem"`
	LoadAvg []string `json:"load_avg"`
	Disk    Disk     `json:"disk"`
	Net     Net      `json:"net"`
	Version string   `json:"version"`

	LocalIp  string `json:"local_ip"`
	UserName string `json:"user_name"`
	Uname    string `json:"uname"`
	Os       string `json:"os"`
	HostName string `json:"host_name"`
}

type Page struct {
	Time Time `json:"time"`
}
type Time struct {
	Start []string `json:"start"`
}

type Mem struct {
	Total         float64 `json:"total"`
	Free          float64 `json:"free"`
	Buffers       float64 `json:"buffers"`
	Cached        float64 `json:"cached"`
	CachedPercent float64 `json:"cached_percent"`
	Used          float64 `json:"used"`
	Percent       float64 `json:"percent"`
	Real          Real    `json:"real"`
	Swap          Swap    `json:"swap"`
}
type Real struct {
	Used    float64 `json:"used"`
	Free    float64 `json:"free"`
	Percent float64 `json:"percent"`
}
type Swap struct {
	Total   int `json:"total"`
	Free    int `json:"free"`
	Used    int `json:"used"`
	Percent int `json:"percent"`
}

type Disk struct {
	Total   float64 `json:"total"`
	Free    float64 `json:"free"`
	Used    float64 `json:"used"`
	Percent float64 `json:"percent"`
}

type Net struct {
	Count      int         `json:"count"`
	Interfaces []Interface `json:"interfaces"`
}

type Cpu struct {
	Stat    Stat     `json:"stat"`
	Freq    int      `json:"freq"`
	Temp    []string `json:"temp"`
	Count   int      `json:"count"`
	Model   string   `json:"model"`
	PiModel string   `json:"pi_model"`
}
type Stat struct {
	User    string `json:"user"`
	Nice    string `json:"nice"`
	Sys     string `json:"sys"`
	Idle    string `json:"idle"`
	Iowait  string `json:"iowait"`
	Irq     string `json:"irq"`
	Softirq string `json:"softirq"`
}

type Interface struct {
	Name     string `json:"name"`
	TotalIn  string `json:"total_in"`
	TotalOut string `json:"total_out"`
}
