package service

import "testing"

func TestCpuTemp(t *testing.T) {
	temp, err := CpuTemp()
	if err != nil {
		t.Error(err)
	}
	t.Log(temp)
}
