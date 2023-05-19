package linux

import "testing"

func TestLoadAvg(t *testing.T) {
	avg, err := LoadAvg()
	if err != nil {
		t.Error(err)
	}
	t.Log(avg)
}
