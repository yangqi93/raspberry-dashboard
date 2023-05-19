package windows

import "github.com/yangqi93/raspberry-dashboard/service"

type archWindows struct{}

func (w archWindows) User() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) Uname() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) Platform() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) LocalIp() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) UpTime() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) CpuFreq() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) CpuInfo() (cpu *service.Cpu, err error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) CpuCore() (stat *service.Stat, err error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) MemInfo() (memory *service.Mem, err error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) LoadAvg() (load *[]string, err error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) NetInfo() (net *service.Net, err error) {
	//TODO implement me
	panic("implement me")
}

func (w archWindows) DiskInfo() (disk *service.Disk, err error) {
	//TODO implement me
	panic("implement me")
}

func NewWindowsArch() service.Arch {
	return &archWindows{}
}

func (w archWindows) CpuTemp() ([]string, error) {
	return []string{"windows 38.5"}, nil
}
