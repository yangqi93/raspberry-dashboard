package service

import "testing"

func TestNetInfo(t *testing.T) {
	info, err := NetInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}
