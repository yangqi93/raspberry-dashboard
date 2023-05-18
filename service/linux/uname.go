package linux

import "os"

func Uname() (string, error) {
	unameR := "N/A"
	uname, err := os.ReadFile("/proc/version")
	if err == nil {
		unameR = string(uname)
	}
	return unameR, nil
}
