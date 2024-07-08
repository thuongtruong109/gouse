package os

import (
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUInfo struct {
	ModelName string
	Family string
	Speed string
	Cores int
}

func GetCPU() (*CPUInfo, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return &CPUInfo{}, err
	}

	return &CPUInfo{
		ModelName: cpuStat[0].ModelName,
		Family: cpuStat[0].Family,
		Speed: strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64),
		Cores: runtime.NumCPU(),
	}, nil
}