package main

import(
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCpu(t *testing.T) {
	cpu, err := os.ReadFile("/proc/stat")
	if err == nil {
		t.Error(string(cpu))
		str := strings.ReplaceAll(string(cpu), "  ", " ")
		//str := string(cpu)
		t.Error(str)
		i := strings.Split(str, " ")
		//i := strings.Split(strings.Join(strings.Split(str, " "), ""), " ")
		t.Error(i)
		fmt.Sprint(i)
	}
}
