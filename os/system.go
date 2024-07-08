package os

import (
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

// const megabyte uint64 = 1024 * 1024
// const gigabyte uint64 = 1024 * megabyte

type SystemInfo struct {
	Name string
	Platform string
	Arch string
	Hostname string
	NumsProcs uint64
	TotalMemory uint64
	FreeMemory uint64
	UsedMemory uint64
	UsedMemoryPercent string
}

func GetSystem() (*SystemInfo, error) {
	runTimeOS := runtime.GOOS

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return &SystemInfo{}, err
	}

	hostStat, err := host.Info()
	if err != nil {
		return &SystemInfo{}, err
	}

	return &SystemInfo{
		Name: runTimeOS,
		Platform: hostStat.Platform,
		Arch: hostStat.KernelArch,
		Hostname: hostStat.Hostname,
		NumsProcs: hostStat.Procs,
		TotalMemory: vmStat.Total,
		FreeMemory: vmStat.Free,
		UsedMemory: vmStat.Used,
		UsedMemoryPercent: strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64),
	}, nil
}