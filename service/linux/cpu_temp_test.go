package linux

import "testing"

func TestCpuTemp(t *testing.T) {
	temp, err := CpuTemp()
	if err != nil {
		t.Log(err)
	}
	t.Log(temp)
}
