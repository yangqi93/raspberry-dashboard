package handle

import (
	"os"
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
	_, err = os.ReadFile("/proc/meminfo")
	if err == nil {
		//$str = implode("", $str);
		//
		//        preg_match_all("/MemTotal\s{0,}\:+\s{0,}([\d\.]+).+?MemFree\s{0,}\:+\s{0,}([\d\.]+).+?Cached\s{0,}\:+\s{0,}([\d\.]+).+?SwapTotal\s{0,}\:+\s{0,}([\d\.]+).+?SwapFree\s{0,}\:+\s{0,}([\d\.]+)/s", $str, $buf);
		//        preg_match_all("/Buffers\s{0,}\:+\s{0,}([\d\.]+)/s", $str, $buffers);
		//
		//        $D['mem']['total'] = round($buf[1][0]/1024, 2);
		//        $D['mem']['free'] = round($buf[2][0]/1024, 2);
		//        $D['mem']['buffers'] = round($buffers[1][0]/1024, 2);
		//        $D['mem']['cached'] = round($buf[3][0]/1024, 2);
		//        $D['mem']['cached_percent'] = (floatval($D['mem']['cached'])!=0)?round($D['mem']['cached']/$D['mem']['total']*100,2):0;
		//        $D['mem']['used'] = $D['mem']['total']-$D['mem']['free'];
		//        $D['mem']['percent'] = (floatval($D['mem']['total'])!=0)?round($D['mem']['used']/$D['mem']['total']*100,2):0;
		//        $D['mem']['real']['used'] = $D['mem']['total'] - $D['mem']['free'] - $D['mem']['cached'] - $D['mem']['buffers'];
		//        $D['mem']['real']['free'] = round($D['mem']['total'] - $D['mem']['real']['used'],2);
		//        $D['mem']['real']['percent'] = (floatval($D['mem']['total'])!=0)?round($D['mem']['real']['used']/$D['mem']['total']*100,2):0;
		//        $D['mem']['swap']['total'] = round($buf[4][0]/1024, 2);
		//        $D['mem']['swap']['free'] = round($buf[5][0]/1024, 2);
		//        $D['mem']['swap']['used'] = round($D['mem']['swap']['total']-$D['mem']['swap']['free'], 2);
		//        $D['mem']['swap']['percent'] = (floatval($D['mem']['swap']['total'])!=0)?round($D['mem']['swap']['used']/$D['mem']['swap']['total']*100,2):0;

	}

	info.LoadAvg = []string{"0.00", "0.00", "0.00", "2\\/275"}

	return *info

}
