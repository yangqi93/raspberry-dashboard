package linux

import "testing"

func TestDiskInfo(t *testing.T) {
	info, err := DiskInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}
