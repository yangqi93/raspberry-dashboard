package linux

import (
	"net"
)

func LocalIp() (string, error) {
	localIp := "N/A"
	adders, err := net.InterfaceAddrs()
	if err != nil {
		return localIp, err
	}
	for _, addr := range adders {
		if ipNet, ok := addr.(*net.IPNet); ok &&
			!ipNet.IP.IsLinkLocalUnicast() && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			localIp = ipNet.IP.String()
		}
	}
	return localIp, nil
}
