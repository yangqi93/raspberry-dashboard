package linux

import "os/exec"

func Os() (string, error) {
	o := exec.Command("uname")
	osR, _ := o.CombinedOutput()
	return string(osR), nil
}
