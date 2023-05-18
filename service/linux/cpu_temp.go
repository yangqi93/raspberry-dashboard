package linux

import "os"

func CpuTemp() ([]string, error) {
	temp, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return []string{""}, err
	}
	return []string{string(temp)}, nil
}
