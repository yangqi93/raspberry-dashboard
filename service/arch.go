package service

type Arch interface {
	// UpTime 获取系统运行时间
	UpTime() (string, error)
	// CpuTemp 获取CPU温度
	CpuTemp() ([]string, error)
	// CpuFreq 获取CPU频率
	CpuFreq() (int, error)
	// CpuInfo 获取CPU信息
	CpuInfo() (cpu *Cpu, err error)
	// CpuCore 获取CPU核心信息
	CpuCore() (stat *Stat, err error)
	// MemInfo 获取内存信息
	MemInfo() (memory *Mem, err error)
	// LoadAvg 获取系统负载
	LoadAvg() (load []string, err error)
	// NetInfo 获取网络信息
	NetInfo() (net *Net, err error)
	// DiskInfo 获取磁盘信息
	DiskInfo() (disk *Disk, err error)
	// User 获取当前用户
	User() (string, error)
	// Uname 获取系统信息
	Uname() (string, error)
	// Platform 获取系统平台
	Platform() (string, error)
	// LocalIp 获取本地IP
	LocalIp() (string, error)
}
