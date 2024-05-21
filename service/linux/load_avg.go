package linux

import (
	"os"
	"strings"
)

func LoadAvg() (load []string, err error) {
	load = []string{"0.00", "0.00", "0.00", "0/0"}
	l, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return load, err
	}
	load = strings.Split(string(l), " ")
	return load, nil
}
