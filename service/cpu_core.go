package service

import (
	"errors"
	"os"
	"strings"
)

func CpuCore() (stat *Stat, err error) {
	cpu, err := os.ReadFile("/proc/stat")
	if err != nil {
		return nil, err
	}
	str := strings.ReplaceAll(string(cpu), "  ", " ")
	i := strings.Split(str, " ")
	if i == nil {
		return nil, errors.New("cpu core info is nil")
	}
	stat.User = i[1]
	stat.Nice = i[2]
	stat.Sys = i[3]
	stat.Idle = i[4]
	stat.Iowait = i[5]
	stat.Irq = i[6]
	stat.Softirq = i[7]

	return stat, nil
}
