package linux

import (
	"os"
	"strings"
)

func UpTime() (string, error) {
	uptime, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "", err
	}
	str := strings.Split(string(uptime), " ")
	return str[0], nil
}
