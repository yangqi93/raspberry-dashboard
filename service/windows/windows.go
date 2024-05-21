package windows

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/yangqi93/raspberry-dashboard/service"
	"gitlab.com/tingshuo/go-diskstate/diskstate"
	"math"
	"os/user"
	"strconv"
	"time"
)

type archWindows struct{}

func (w archWindows) User() (string, error) {
	userName := "N/A"
	u, err := user.Current()
	if err != nil && u != nil {
		userName = u.Username
	}
	return userName, err
}

func (w archWindows) Uname() (string, error) {
	h, _ := host.Info()
	return h.Hostname, nil
}

func (w archWindows) Platform() (string, error) {
	h, _ := host.Info()
	return h.Platform, nil
}

func (w archWindows) LocalIp() (string, error) {
	h, _ := host.Info()
	return h.HostID, nil
}

func (w archWindows) UpTime() (string, error) {
	//获取系统运行时间
	c, _ := cpu.Times(false)
	return fmt.Sprintf("%d", int(c[0].Total())), nil
}

func (w archWindows) CpuFreq() (int, error) {
	c, _ := cpu.Info()
	return int(c[0].Mhz), nil
}

func (w archWindows) CpuInfo() (cpuInfo *service.Cpu, err error) {
	c, _ := cpu.Info()
	return &service.Cpu{
		Count:   int(c[0].Cores),
		Model:   c[0].ModelName,
		PiModel: c[0].Model,
	}, nil
}

func (w archWindows) CpuCore() (stat *service.Stat, err error) {
	c, _ := cpu.Times(false)
	return &service.Stat{
		User:    fmt.Sprintf("%.2f", c[0].User),
		Nice:    fmt.Sprintf("%.2f", c[0].Nice),
		Sys:     fmt.Sprintf("%.2f", c[0].System),
		Idle:    fmt.Sprintf("%.2f", c[0].Idle),
		Iowait:  fmt.Sprintf("%.2f", c[0].Iowait),
		Irq:     fmt.Sprintf("%.2f", c[0].Irq),
		Softirq: fmt.Sprintf("%.2f", c[0].Softirq),
	}, nil
}

func (w archWindows) MemInfo() (memory *service.Mem, err error) {
	m1, _ := mem.VirtualMemory()
	m2, _ := mem.SwapMemory()
	memory = &service.Mem{
		Total:         math.Round(float64(m1.Total / 8 / 1024 / 1024)),
		Free:          math.Round(float64(m1.Free / 8 / 1024 / 1024)),
		Buffers:       math.Round(float64(m1.Buffers / 8 / 1024 / 1024)),
		Cached:        math.Round(float64(m1.Cached / 8 / 1024 / 1024)),
		CachedPercent: math.Round(float64(m1.Cached / m1.Total)),
		Used:          math.Round(float64(m1.Used / 8 / 1024 / 1024)),
		Percent:       math.Round(m1.UsedPercent),
		Real: service.Real{
			Used:    math.Round(float64(m1.Used / 8 / 1024 / 1024)),
			Free:    math.Round(float64(m1.Free / 8 / 1024 / 1024)),
			Percent: math.Round(m1.UsedPercent),
		},
		Swap: service.Swap{
			Total:   int(math.Round(float64(m2.Total / 8 / 1024 / 1024))),
			Free:    int(math.Round(float64(m2.Free / 8 / 1024 / 1024))),
			Used:    int(math.Round(float64(m2.Used / 8 / 1024 / 1024))),
			Percent: int(math.Round(m2.UsedPercent)),
		},
	}
	return memory, nil
}

func (w archWindows) LoadAvg() (load []string, err error) {
	c, _ := cpu.Percent(time.Second*1, true)
	//c5, _ := cpu.Percent(time.Minute*5, false)
	//c10, _ := cpu.Percent(time.Minute*10, false)
	//load = []string{
	//	fmt.Sprintf("%.2f", c[0]),
	//	fmt.Sprintf("%.2f", c5[0]),
	//	fmt.Sprintf("%.2f", c10[0]),
	//}
	for _, v := range c {
		load = append(load, fmt.Sprintf("%.2f", v))
	}
	return load, nil
}

func (w archWindows) NetInfo() (netInfo *service.Net, err error) {
	n, _ := net.IOCounters(false)
	info := &service.Net{}
	for _, v := range n {
		info.Interfaces = append(info.Interfaces, service.Interface{
			Name:     v.Name,
			TotalIn:  strconv.FormatUint(v.BytesRecv, 10),
			TotalOut: strconv.FormatUint(v.BytesSent, 10),
		})
	}
	info.Count = len(info.Interfaces)
	return info, nil

}

func (w archWindows) DiskInfo() (di *service.Disk, err error) {
	diskInfo, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	var drivers []string
	for _, d := range diskInfo {
		drivers = append(drivers, d.Device)
	}
	info, err := disk.IOCounters(drivers...)
	if err != nil {
		return nil, err
	}
	var total, used, free float64
	for n := range info {
		state := diskstate.DiskUsage(n)
		total += float64(state.All / diskstate.GB)
		used += float64(state.Used / diskstate.GB)
		free += float64(state.Free / diskstate.GB)
	}
	return &service.Disk{
		Total:   total,
		Free:    free,
		Used:    used,
		Percent: math.Round(100 * used / total),
	}, nil

}

func (w archWindows) CpuTemp() ([]string, error) {
	return nil, errors.New("not support")
}

func NewWindowsArch() service.Arch {
	return &archWindows{}
}
