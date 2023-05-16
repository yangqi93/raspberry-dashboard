package service

import "testing"

func TestUpTime(t *testing.T) {
	uptime, err := UpTime()
	if err != nil {
		t.Error(err)
	}
	t.Log(uptime)
}
