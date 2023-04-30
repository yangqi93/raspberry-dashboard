package handle

import (
	"gitlab.com/tingshuo/go-diskstate/diskstate"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetInfo() Status {
	//计算树莓派的运行信息
	info := &Status{
		Time: time.Now().Unix(),
	}
	info.Page.Time.Start = []string{
		//当前时间微秒数
		strconv.FormatInt(time.Now().UnixNano(), 10),
		//当前时间秒数
		strconv.FormatInt(time.Now().Unix(), 10),
	}
	//uptime
	uptime, err := os.ReadFile("/proc/uptime")
	if err == nil {
		str := strings.Split(string(uptime), " ")
		info.Uptime = str[0]
	}
	//cpu
	f, err := os.ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq")
	if err == nil {
		freq := strings.ReplaceAll(string(f), "\n", "")
		info.Cpu.Freq, _ = strconv.Atoi(freq)
	}
	c, err := os.ReadFile("/proc/cpuinfo")
	if err == nil {
		var model, bogomips, pimodel string
		i := 0
		str := strings.Split(string(c), "\n")
		for _, v := range str {
			if strings.Contains(v, "Hardware") {
				model = strings.Split(v, ":")[1]
			}
			if strings.Contains(v, "BogoMIPS") {
				bogomips = strings.Split(v, ":")[1]
				i++
			}
			if strings.Contains(v, "Model") {
				pimodel = strings.Split(v, ":")[1]
			}
		}
		info.Cpu.Count = i
		if i == 1 {
			info.Cpu.Model = model + " " + bogomips
		} else {
			info.Cpu.Model = model + " " + bogomips + " x" + strconv.Itoa(i)
		}
		info.Cpu.PiModel = pimodel
	}

	//cpu core
	cpu, err := os.ReadFile("/proc/stat")
	if err == nil {
		str := strings.ReplaceAll(string(cpu), "  ", " ")
		i := strings.Split(str, " ")
		info.Cpu.Stat.User = i[1]
		info.Cpu.Stat.Nice = i[2]
		info.Cpu.Stat.Sys = i[3]
		info.Cpu.Stat.Idle = i[4]
		info.Cpu.Stat.Iowait = i[5]
		info.Cpu.Stat.Irq = i[6]
		info.Cpu.Stat.Softirq = i[7]
	}
	//cpu temp
	temp, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err == nil {
		info.Cpu.Temp = []string{string(temp)}
	} else {
		info.Cpu.Temp = []string{""}
	}

	//meminfo
	mem, err := os.ReadFile("/proc/meminfo")
	if err == nil {
		str := string(mem)
		//正则匹配
		reg := regexp.MustCompile(`MemTotal\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
		var total, free, buffers, cached, cached_percent, used, percent, real_used, real_free, real_percent, swap_total, swap_free, swap_used, swap_percent float64
		if len(reg) >= 2 {
			i, _ := strconv.Atoi(reg[1])
			total = math.Round(float64(i / 1024))
		}
		reg = regexp.MustCompile(`MemFree\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
		if len(reg) >= 2 {
			i, _ := strconv.Atoi(reg[1])
			free = math.Round(float64(i / 1024))
		}
		reg = regexp.MustCompile(`Buffers\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
		if len(reg) >= 2 {
			i, _ := strconv.Atoi(reg[1])
			buffers = math.Round(float64(i / 1024))
		}
		reg = regexp.MustCompile(`Cached\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
		if len(reg) >= 2 {
			i, _ := strconv.Atoi(reg[1])
			cached = math.Round(float64(i / 1024))
		}
		reg = regexp.MustCompile(`SwapTotal\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
		if len(reg) >= 2 {
			i, _ := strconv.Atoi(reg[1])
			swap_total = math.Round(float64(i / 1024))
		}
		reg = regexp.MustCompile(`SwapFree\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
		if len(reg) >= 2 {
			i, _ := strconv.Atoi(reg[1])
			swap_free = math.Round(float64(i / 1024))
		}
		// cached percent
		if cached != 0 && total != 0 {
			cached_percent = math.Round(cached / total * 100)
		}
		// used
		if total != 0 && free != 0 {
			used = total - free
		}
		// percent
		if total != 0 && used != 0 {
			percent = math.Round(used / total * 100)
		}
		// real_used
		if total != 0 && free != 0 && buffers != 0 && cached != 0 {
			real_used = total - free - buffers - cached
		}
		// real_free
		if total != 0 && real_used != 0 {
			real_free = math.Round(total - real_used)
		}
		// real_percent
		if total != 0 && real_used != 0 {
			real_percent = math.Round(real_used / total * 100)
		}
		// swap_used
		if swap_total != 0 && swap_free != 0 {
			swap_used = math.Round(swap_total - swap_free)
		}
		// swap_percent
		if swap_total != 0 && swap_used != 0 {
			swap_percent = math.Round(swap_used / swap_total * 100)
		}

		info.Mem.Total = total
		info.Mem.Free = free
		info.Mem.Buffers = buffers
		info.Mem.Cached = cached
		info.Mem.CachedPercent = cached_percent
		info.Mem.Used = used
		info.Mem.Percent = percent
		info.Mem.Real.Used = real_used
		info.Mem.Real.Free = real_free
		info.Mem.Real.Percent = real_percent
		info.Mem.Swap.Total = int(swap_total)
		info.Mem.Swap.Free = int(swap_free)
		info.Mem.Swap.Used = int(swap_used)
		info.Mem.Swap.Percent = int(swap_percent)
	}

	info.LoadAvg = []string{"0.00", "0.00", "0.00", "0/0"}
	load, err := os.ReadFile("/proc/loadavg")
	if err == nil {
		str := strings.Split(string(load), " ")
		info.LoadAvg = str
	}

	//net
	net, err := os.ReadFile("/proc/net/dev")
	if err == nil {
		str := strings.Split(string(net), "\n")
		info.Net.Count = len(str) - 3
		for _, v := range str {
			if strings.Contains(v, ":") {
				info.Net.Interfaces = append(info.Net.Interfaces, Interface{
					Name:     strings.Split(v, ":")[0],
					TotalIn:  strings.Fields(v)[1],
					TotalOut: strings.Fields(v)[9],
				})
			}
		}
	}

	//disk
	state := diskstate.DiskUsage("/")
	info.Disk.Total = float64(state.All / diskstate.GB)
	info.Disk.Used = float64(state.Used / diskstate.GB)
	info.Disk.Free = float64(state.Free / diskstate.GB)
	info.Disk.Percent = float64(state.Used / state.All * 100)
	return *info
}
