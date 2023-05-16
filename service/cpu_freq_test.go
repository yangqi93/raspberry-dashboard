package service

import "testing"

func TestCpuFreq(t *testing.T) {
	freq, err := CpuFreq()
	if err != nil {
		t.Error(err)
	}
	t.Log(freq)
}
