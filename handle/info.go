package handle

import (
	"errors"
	"github.com/yangqi93/raspberry-dashboard/service"
	"github.com/yangqi93/raspberry-dashboard/service/linux"
	"github.com/yangqi93/raspberry-dashboard/service/windows"
	"runtime"
	"strconv"
	"time"
)

func GetInfo() (service.Status, error) {
	//根据平台选择不同的实例
	arch := runtime.GOOS
	var bios service.Arch
	switch arch {
	case "linux":
		bios = linux.NewLinuxArch()
	case "windows":
		bios = windows.NewWindowsArch()
	default:
		return service.Status{}, errors.New("not support platform")
	}

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
	uptime, err := bios.UpTime()
	if err != nil {
		return *info, err
	}
	info.Uptime = uptime

	//cpu
	f, err := bios.CpuFreq()
	if err == nil {
		info.Cpu.Freq = f
	}
	c, err := bios.CpuInfo()
	if err == nil {
		info.Cpu.Count = c.Count
		info.Cpu.Model = c.Model
		info.Cpu.PiModel = c.PiModel
	}

	//cpu core
	stat, err := bios.CpuCore()
	if err == nil {
		info.Cpu.Stat = *stat
	}

	//cpu temp
	temp, err := bios.CpuTemp()
	if err == nil {
		info.Cpu.Temp = temp
	}

	//meminfo
	mem, err := bios.MemInfo()
	if err == nil {
		info.Mem = *mem
	}

	load, err := bios.LoadAvg()
	if err == nil {
		info.LoadAvg = load
	}

	//net
	net, err := bios.NetInfo()
	if err == nil {
		info.Net = *net
	}

	//disk
	disk, err := bios.DiskInfo()
	if err == nil {
		info.Disk = *disk
	}

	ip, err := bios.LocalIp()
	if err == nil {
		info.LocalIp = ip
	}

	os, err := bios.Platform()
	if err == nil {
		info.Os = os
	}

	uname, err := bios.Uname()
	if err == nil {
		info.Uname = uname
	}
	return *info, nil
}
