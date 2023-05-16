package service

import "testing"

func TestCpuFreq(t *testing.T) {
	freq, err := CpuFreq()
	if err != nil {
		t.Log(err)
	}
	t.Log(freq)
}
