package linux

import "testing"

func TestCpuCore(t *testing.T) {
	stat, err := CpuCore()
	if err != nil {
		t.Error(err)
	}
	t.Log(stat)
}
