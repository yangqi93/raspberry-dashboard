package linux

import (
	"github.com/yangqi93/raspberry-dashboard/service"
	"os"
	"strings"
)

func NetInfo() (net *service.Net, err error) {
	net = &service.Net{}
	netInfo, err := os.ReadFile("/proc/net/dev")
	if err != nil {
		return nil, err
	}
	str := strings.Split(string(netInfo), "\n")
	net.Count = len(str) - 3
	for _, v := range str {
		if strings.Contains(v, ":") {
			net.Interfaces = append(net.Interfaces, service.Interface{
				Name:     strings.Split(v, ":")[0],
				TotalIn:  strings.Fields(v)[1],
				TotalOut: strings.Fields(v)[9],
			})
		}
	}
	return net, nil
}
