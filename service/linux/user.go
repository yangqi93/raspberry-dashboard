package linux

import "os/user"

func User() (string, error) {
	userName := "N/A"
	u, err := user.Current()
	if err != nil {
		userName = u.Username
	}
	return userName, err
}
