package os

import (
	"strconv"

	"github.com/shirou/gopsutil/v4/disk"
)

type DiskInfo struct {
	TotalSpace uint64
	FreeSpace uint64
	UsedSpace uint64
	UsedSpacePercent string
}

func GetDisk() (*DiskInfo, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return &DiskInfo{}, err
	}

	return &DiskInfo{
		TotalSpace: diskStat.Total,
		FreeSpace: diskStat.Free,
		UsedSpace: diskStat.Used,
		UsedSpacePercent: strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64),
	}, nil

	return &DiskInfo{}, nil
}