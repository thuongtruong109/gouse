package gouse

import (
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

/* CPU */

type ICPU struct {
	ModelName string
	Family    string
	Speed     string
	Cores     int
}

func CPU() (*ICPU, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return &ICPU{}, err
	}

	return &ICPU{
		ModelName: cpuStat[0].ModelName,
		Family:    cpuStat[0].Family,
		Speed:     strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64),
		Cores:     runtime.NumCPU(),
	}, nil
}

/* Disk */

type IDisk struct {
	TotalSpace       uint64
	FreeSpace        uint64
	UsedSpace        uint64
	UsedSpacePercent string
}

func Disk() (*IDisk, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return &IDisk{}, err
	}

	return &IDisk{
		TotalSpace:       diskStat.Total,
		FreeSpace:        diskStat.Free,
		UsedSpace:        diskStat.Used,
		UsedSpacePercent: strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64),
	}, nil
}

/* System */

// const megabyte uint64 = 1024 * 1024
// const gigabyte uint64 = 1024 * megabyte

type ISystem struct {
	Name              string
	Platform          string
	Arch              string
	Hostname          string
	NumsProcs         uint64
	TotalMemory       uint64
	FreeMemory        uint64
	UsedMemory        uint64
	UsedMemoryPercent string
}

func System() (*ISystem, error) {
	runTimeOS := runtime.GOOS

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return &ISystem{}, err
	}

	hostStat, err := host.Info()
	if err != nil {
		return &ISystem{}, err
	}

	return &ISystem{
		Name:              runTimeOS,
		Platform:          hostStat.Platform,
		Arch:              hostStat.KernelArch,
		Hostname:          hostStat.Hostname,
		NumsProcs:         hostStat.Procs,
		TotalMemory:       vmStat.Total,
		FreeMemory:        vmStat.Free,
		UsedMemory:        vmStat.Used,
		UsedMemoryPercent: strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64),
	}, nil
}
