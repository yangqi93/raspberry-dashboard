package service

type Arch interface {
	UpTime() (string, error)
	CpuTemp() ([]string, error)
	CpuFreq() (int, error)
	CpuInfo() (cpu *Cpu, err error)
	CpuCore() (stat *Stat, err error)
	MemInfo() (memory *Mem, err error)
	LoadAvg() (load *[]string, err error)
	NetInfo() (net *Net, err error)
	DiskInfo() (disk *Disk, err error)
	User() (string, error)
	Uname() (string, error)
	Platform() (string, error)
	LocalIp() (string, error)
}
