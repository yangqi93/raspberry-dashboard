package handle

import (
	"github.com/yangqi93/raspberry-dashboard/service"
	"strconv"
	"time"
)

func GetInfo() service.Status {
	//计算树莓派的运行信息
	info := &service.Status{
		Time:    time.Now().Unix(),
		Version: "1.0",
	}
	info.Page.Time.Start = []string{
		//当前时间微秒数
		strconv.FormatInt(time.Now().UnixNano(), 10),
		//当前时间秒数
		strconv.FormatInt(time.Now().Unix(), 10),
	}

	//uptime
	uptime, err := service.UpTime()
	if err == nil {
		info.Uptime = uptime
	}

	//cpu
	f, err := service.CpuFreq()
	if err == nil {
		info.Cpu.Freq = f
	}
	c, err := service.CpuInfo()
	if err == nil {
		info.Cpu.Count = c.Count
		info.Cpu.Model = c.Model
		info.Cpu.PiModel = c.PiModel
	}

	//cpu core
	stat, err := service.CpuCore()
	if err == nil {
		info.Cpu.Stat = *stat
	}

	//cpu temp
	temp, err := service.CpuTemp()
	if err == nil {
		info.Cpu.Temp = temp
	}

	//meminfo
	mem, err := service.MemInfo()
	if err == nil {
		info.Mem = *mem
	}

	load, err := service.LoadAvg()
	if err == nil {
		info.LoadAvg = *load
	}

	//net
	net, err := service.NetInfo()
	if err == nil {
		info.Net = *net
	}

	//disk
	disk, err := service.DiskInfo()
	if err == nil {
		info.Disk = *disk
	}
	return *info
}
