package linux

import "github.com/yangqi93/raspberry-dashboard/service"

type archLinux struct {
}

func (a archLinux) User() (string, error) {
	return User()
}

func (a archLinux) Uname() (string, error) {
	return Uname()
}

func (a archLinux) Platform() (string, error) {
	return Os()
}

func (a archLinux) LocalIp() (string, error) {
	return LocalIp()
}

func (a archLinux) CpuInfo() (cpu *service.Cpu, err error) {
	return CpuInfo()
}

func (a archLinux) CpuCore() (stat *service.Stat, err error) {
	return CpuCore()
}

func (a archLinux) MemInfo() (memory *service.Mem, err error) {
	return MemInfo()
}

func (a archLinux) LoadAvg() (load *[]string, err error) {
	return LoadAvg()
}

func (a archLinux) NetInfo() (net *service.Net, err error) {
	return NetInfo()
}

func (a archLinux) DiskInfo() (disk *service.Disk, err error) {
	return DiskInfo()
}

func (a archLinux) CpuFreq() (int, error) {
	return CpuFreq()
}

func (a archLinux) UpTime() (string, error) {
	return UpTime()
}

func (a archLinux) CpuTemp() ([]string, error) {
	return CpuTemp()
}

func NewLinuxArch() service.Arch {
	return &archLinux{}
}
