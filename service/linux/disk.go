package linux

import (
	"github.com/yangqi93/raspberry-dashboard/service"
	"gitlab.com/tingshuo/go-diskstate/diskstate"
)

func DiskInfo() (disk *service.Disk, err error) {
	disk = &service.Disk{}
	state := diskstate.DiskUsage("/")
	disk.Total = float64(state.All / diskstate.GB)
	disk.Used = float64(state.Used / diskstate.GB)
	disk.Free = float64(state.Free / diskstate.GB)
	disk.Percent = float64(100 * state.Used / state.All)
	return disk, nil
}
