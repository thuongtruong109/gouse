package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Get CPU information
Params: None -> *gouse.ICPU
*/
func OsCpu() {
	cpuInfo, err := gouse.CPU()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(":: CPU Information")
	fmt.Println("Model Name: ", cpuInfo.ModelName)
	fmt.Println("Family: ", cpuInfo.Family)
	fmt.Println("Speed: ", cpuInfo.Speed, "MHz")
	fmt.Println("Cores: ", cpuInfo.Cores)
	fmt.Println("Vendor ID: ", cpuInfo.VendorID)
	fmt.Println("Model: ", cpuInfo.Model)
	fmt.Println("Stepping: ", cpuInfo.Stepping)
	fmt.Println("Physical ID: ", cpuInfo.PhysicalID)
	fmt.Println("Core ID: ", cpuInfo.CoreID)
	fmt.Println("Mhz: ", cpuInfo.Mhz)
	fmt.Println("Cache Size: ", cpuInfo.CacheSize)
	fmt.Println("Flags: ", cpuInfo.Flags)
	fmt.Println("Microcode: ", cpuInfo.Microcode)
}

/*
Description: Get Disk information
Params: None -> *gouse.IDisk
*/
func OsDisk() {
	diskInfo, err := gouse.Disk()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(":: Disk Information")
	fmt.Println("Total Space: ", diskInfo.Total, "MB")
	fmt.Println("Free Space: ", diskInfo.Free, "MB")
	fmt.Println("Used Space: ", diskInfo.Used, "MB")
	fmt.Println("Used Space Percent: ", diskInfo.UsedPercent, "%")
	fmt.Println("File System Type: ", diskInfo.Fstype)
	fmt.Println("Path: ", diskInfo.Path)
	fmt.Println("Inodes Total: ", diskInfo.InodesTotal)
	fmt.Println("Inodes Used: ", diskInfo.InodesUsed)
	fmt.Println("Inodes Free: ", diskInfo.InodesFree)
	fmt.Println("Inodes Used Percent: ", diskInfo.InodesUsedPercent, "%")
}

/*
Description: Get Memory information
Params: None -> *gouse.IHost
*/
func OsHost() {
	hostInfo, err := gouse.Host()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(":: Host Information")
	fmt.Println("Name: ", hostInfo.Name)
	fmt.Println("Platform: ", hostInfo.Platform)
	fmt.Println("Arch: ", hostInfo.Arch)
	fmt.Println("Hostname: ", hostInfo.Hostname)
	fmt.Println("Number of Processes: ", hostInfo.NumsProcs)
	fmt.Println("Uptime: ", hostInfo.Uptime)
	fmt.Println("Boot Time: ", hostInfo.BootTime)
	fmt.Println("OS: ", hostInfo.OS)
	fmt.Println("Kernel Version: ", hostInfo.KernelVersion)
	fmt.Println("Host ID: ", hostInfo.HostID)
	fmt.Println("Virtualization System: ", hostInfo.VirtualizationSystem)
	fmt.Println("Virtualization Role: ", hostInfo.VirtualizationRole)
	fmt.Println("Platform Family: ", hostInfo.PlatformFamily)
	fmt.Println("Platform Version: ", hostInfo.PlatformVersion)
}

/*
Description: Get OS Partition information
Params: None -> *gouse.IPartition
*/
func OsPartion() {
	partitions, err := gouse.Partition()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(":: Partitions Information")
	for _, v := range partitions {
		fmt.Println("Device: ", v.Device)
		fmt.Println("Mountpoint: ", v.Mountpoint)
		fmt.Println("Fstype: ", v.Fstype)
		fmt.Println("Opts: ", v.Opts)
	}
}

/*
Description: Get OS Process information
Params: None -> *gouse.IProcess
*/
func OsUser() {
	users, err := gouse.User()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(":: Users Information")
	for _, v := range users {
		fmt.Println("User: ", v.User)
		fmt.Println("Terminal: ", v.Terminal)
		fmt.Println("Host: ", v.Host)
		fmt.Println("Started: ", v.Started)
	}
}

/*
Description: Get OS I/O information
Params: None -> *gouse.IIo
*/
func OsIO() {
	ioStat, err := gouse.Io()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(":: IO Information")
	for _, v := range ioStat {
		fmt.Println("Read Count: ", v.ReadCount)
		fmt.Println("Merge Read Count: ", v.MergedReadCount)
		fmt.Println("Write Count: ", v.WriteCount)
		fmt.Println("Merge Write Count: ", v.MergedWriteCount)
		fmt.Println("Read Bytes: ", v.ReadBytes)
		fmt.Println("Write Bytes: ", v.WriteBytes)
		fmt.Println("Read Time: ", v.ReadTime)
		fmt.Println("Write Time: ", v.WriteTime)
		fmt.Println("Iops In Progress: ", v.IopsInProgress)
		fmt.Println("I/O Time: ", v.IoTime)
		fmt.Println("Weighted IO: ", v.WeightedIO)
		fmt.Println("Name: ", v.Name)
		fmt.Println("Serial Number: ", v.SerialNumber)
		fmt.Println("Label: ", v.Label)
	}
}

/*
Description: Get Memory information
Params: None -> *gouse.IMemory
*/
func OsMemory() {
	memoryInfo, err := gouse.Memory()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(":: Memory Information")
	fmt.Println("Virtual Total: ", memoryInfo.VirtualTotal, "MB")
	fmt.Println("Virtual Used: ", memoryInfo.VirtualUsed, "MB")
	fmt.Println("Virtual Free: ", memoryInfo.VirtualFree, "MB")
	fmt.Println("Virtual Used Percent: ", memoryInfo.VirtualUsedPercent, "%")
	fmt.Println("Virtual Available: ", memoryInfo.VirtualAvailable, "MB")
	fmt.Println("Virtual Active: ", memoryInfo.VirtualActive, "MB")
	fmt.Println("Virtual Inactive: ", memoryInfo.VirtualInactive, "MB")
	fmt.Println("Vrtual Wired: ", memoryInfo.VirtualWired, "MB")
	fmt.Println("Virtual Laundry: ", memoryInfo.VirtualLaundry, "MB")
	fmt.Println("Virtual Buffers: ", memoryInfo.VirtualBuffers, "MB")
	fmt.Println("Virtual Cached: ", memoryInfo.VirtualCached, "MB")
	fmt.Println("Virtual Write Back: ", memoryInfo.VirtualWriteBack, "MB")
	fmt.Println("Virtual Dirty: ", memoryInfo.VirtualDirty, "MB")
	fmt.Println("Virtual Write Back Tmp: ", memoryInfo.VirtualWriteBackTmp, "MB")
	fmt.Println("Virtual Shared: ", memoryInfo.VirtualShared, "MB")
	fmt.Println("Virtual Slab: ", memoryInfo.VirtualSlab, "MB")
	fmt.Println("Virtual Sreclaimable: ", memoryInfo.VirtualSreclaimable, "MB")
	fmt.Println("Virtual Sunreclaim: ", memoryInfo.VirtualSunreclaim, "MB")
	fmt.Println("Virtual Page Tables: ", memoryInfo.VirtualPageTables, "MB")
	fmt.Println("Virtual Swap Cached: ", memoryInfo.VirtualSwapCached, "MB")
	fmt.Println("Virtual Commit Limit: ", memoryInfo.VirtualCommitLimit, "MB")
	fmt.Println("Virtual Committed AS: ", memoryInfo.VirtualCommittedAS, "MB")
	fmt.Println("Virtual High Total: ", memoryInfo.VirtualHighTotal, "MB")
	fmt.Println("Virtual High Free: ", memoryInfo.VirtualHighFree, "MB")
	fmt.Println("Virtual Low Total: ", memoryInfo.VirtualLowTotal, "MB")
	fmt.Println("Virtual Low Free: ", memoryInfo.VirtualLowFree, "MB")
	fmt.Println("Virtual Mapped: ", memoryInfo.VirtualMapped, "MB")
	fmt.Println("Virtual Malloc Total: ", memoryInfo.VirtualMallocTotal, "MB")
	fmt.Println("Virtual Malloc Used: ", memoryInfo.VirtualMallocUsed, "MB")
	fmt.Println("Virtual Malloc Chunk: ", memoryInfo.VirtualMallocChunk, "MB")
	fmt.Println("Virtual Huge Pages Total: ", memoryInfo.VirtualHugePagesTotal, "MB")
	fmt.Println("Virtual Huge Pages Free: ", memoryInfo.VirtualHugePagesFree, "MB")
	fmt.Println("Virtual Huge Pages Rsvd: ", memoryInfo.VirtualHugePagesRsvd, "MB")
	fmt.Println("Virtual Huge Pages Surp: ", memoryInfo.VirtualHugePagesSurp, "MB")
	fmt.Println("Virtual Huge Page Size: ", memoryInfo.VirtualHugePageSize, "MB")
	fmt.Println("Virtual Anon Huge Pages: ", memoryInfo.VirtualAnonHugePages, "MB")

	fmt.Println("Swap Total: ", memoryInfo.SwapTotal, "MB")
	fmt.Println("Swap Used: ", memoryInfo.SwapUsed, "MB")
	fmt.Println("Swap Free: ", memoryInfo.SwapFree, "MB")
	fmt.Println("Swap Used Percent: ", memoryInfo.SwapUsedPercent, "%")
	fmt.Println("Swap Sin: ", memoryInfo.SwapSin, "MB")
	fmt.Println("Swap Sout: ", memoryInfo.SwapSout, "MB")
	fmt.Println("Swap Pg In: ", memoryInfo.SwapPgIn, "MB")
	fmt.Println("Swap Pg Out: ", memoryInfo.SwapPgOut, "MB")
	fmt.Println("Swap Pg Fault: ", memoryInfo.SwapPgFault, "MB")
	fmt.Println("Swap Pg Maj Fault: ", memoryInfo.SwapPgMajFault, "MB")
}

/*
Description: Run Profile
Input params: (cpu.pprof, mem.pprof)
*/
func OsProfile() {
	var cpuprofile, memprofile = "cpu.pprof", "mem.pprof"
	gouse.Profile(cpuprofile, memprofile)

	// run this command to test
	// go tool pprof <cpu.pprof_or_mem.pprof>

	// How to generate visual analysis images
	//	Open https://graphviz.gitlab.io/download/ and follow the prompts to download and install.
	// After the installation is complete, for Windows, add the bin folder of the Graphviz installation path after setting the environment variable path.

	// Update main.go
	// 	package main
	// import _ "net/http/pprof"
	// import "net/http"
	// func main() {
	//     go func() {
	//         _ = http.ListenAndServe("0.0.0.0:8081", nil)
	//     }()
	//     // your code
	// }

	// How to check the memory size of each module
	// go tool pprof -alloc_space -cum http://localhost:8081/debug/pprof/heap

	// View specific data list in web browser
	// http://localhost:8081/debug/pprof/heap?debug=1

	// 	Command line method to generate visual analysis images
	// 	go tool pprof http://localhost:8081/debug/pprof/goroutine

	//	View specific data list in web browser
	//
	// http://localhost:8081/debug/pprof/goroutine?debug=1
}
