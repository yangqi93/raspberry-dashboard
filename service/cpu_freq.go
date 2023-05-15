package service

import (
	"os"
	"strconv"
	"strings"
)

func CpuFreq() (int, error) {
	f, err := os.ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq")
	if err != nil {
		return 0, err
	}
	freq := strings.ReplaceAll(string(f), "\n", "")
	return strconv.Atoi(freq)
}
