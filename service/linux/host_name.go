package linux

import "github.com/yangqi93/raspberry-dashboard/config"

func HostName() (string, error) {
	return config.Conf.Value.GetString("hostName"), nil
}
