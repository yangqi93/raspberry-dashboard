package linux

import "testing"

func TestLocalIp(t *testing.T) {
	ip, err := LocalIp()
	if err != nil {
		t.Error(err)
	}
	t.Log(ip)
}
