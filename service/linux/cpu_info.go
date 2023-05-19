package linux

import (
	"github.com/yangqi93/raspberry-dashboard/service"
	"os"
	"strconv"
	"strings"
)

func CpuInfo() (cpu *service.Cpu, err error) {
	cpu = &service.Cpu{}
	c, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return nil, err
	}
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
	cpu.Count = i
	if i == 1 {
		cpu.Model = model + " " + bogomips
	} else {
		cpu.Model = model + " " + bogomips + " x" + strconv.Itoa(i)
	}
	cpu.PiModel = pimodel
	return cpu, nil
}
