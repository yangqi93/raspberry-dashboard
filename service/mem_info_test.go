package service

import "testing"

func TestMemInfo(t *testing.T) {
	info, err := MemInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}
