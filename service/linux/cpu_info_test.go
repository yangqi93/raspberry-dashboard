package linux

import "testing"

func TestCpuInfo(t *testing.T) {
	info, err := CpuInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}
