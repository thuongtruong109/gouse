
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Os' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Os cpu

Description: Get CPU information<br>Params: None -> *gouse.ICPU<br>

```go
func OsCpu() {
	cpuInfo, err := gouse.CPU()
	if err != nil {
		gouse.Println("Error: ", err)
		return
	}

	gouse.Println(":: CPU Information")
	gouse.Println("Model Name: ", cpuInfo.ModelName)
	gouse.Println("Family: ", cpuInfo.Family)
	gouse.Println("Speed: ", cpuInfo.Speed, "MHz")
	gouse.Println("Cores: ", cpuInfo.Cores)
	gouse.Println("Vendor ID: ", cpuInfo.VendorID)
	gouse.Println("Model: ", cpuInfo.Model)
	gouse.Println("Stepping: ", cpuInfo.Stepping)
	gouse.Println("Physical ID: ", cpuInfo.PhysicalID)
	gouse.Println("Core ID: ", cpuInfo.CoreID)
	gouse.Println("Mhz: ", cpuInfo.Mhz)
	gouse.Println("Cache Size: ", cpuInfo.CacheSize)
	gouse.Println("Flags: ", cpuInfo.Flags)
	gouse.Println("Microcode: ", cpuInfo.Microcode)
}
```

## 2. Os disk

Description: Get Disk information<br>Params: None -> *gouse.IDisk<br>

```go
func OsDisk() {
	diskInfo, err := gouse.Disk()
	if err != nil {
		gouse.Println("Error: ", err)
		return
	}

	gouse.Println(":: Disk Information")
	gouse.Println("Total Space: ", diskInfo.Total, "MB")
	gouse.Println("Free Space: ", diskInfo.Free, "MB")
	gouse.Println("Used Space: ", diskInfo.Used, "MB")
	gouse.Println("Used Space Percent: ", diskInfo.UsedPercent, "%")
	gouse.Println("File System Type: ", diskInfo.Fstype)
	gouse.Println("Path: ", diskInfo.Path)
	gouse.Println("Inodes Total: ", diskInfo.InodesTotal)
	gouse.Println("Inodes Used: ", diskInfo.InodesUsed)
	gouse.Println("Inodes Free: ", diskInfo.InodesFree)
	gouse.Println("Inodes Used Percent: ", diskInfo.InodesUsedPercent, "%")
}
```

## 3. Os host

Description: Get Memory information<br>Params: None -> *gouse.IHost<br>

```go
func OsHost() {
	hostInfo, err := gouse.Host()
	if err != nil {
		gouse.Println("Error: ", err)
		return
	}

	gouse.Println(":: Host Information")
	gouse.Println("Name: ", hostInfo.Name)
	gouse.Println("Platform: ", hostInfo.Platform)
	gouse.Println("Arch: ", hostInfo.Arch)
	gouse.Println("Hostname: ", hostInfo.Hostname)
	gouse.Println("Number of Processes: ", hostInfo.NumsProcs)
	gouse.Println("Uptime: ", hostInfo.Uptime)
	gouse.Println("Boot Time: ", hostInfo.BootTime)
	gouse.Println("OS: ", hostInfo.OS)
	gouse.Println("Kernel Version: ", hostInfo.KernelVersion)
	gouse.Println("Host ID: ", hostInfo.HostID)
	gouse.Println("Virtualization System: ", hostInfo.VirtualizationSystem)
	gouse.Println("Virtualization Role: ", hostInfo.VirtualizationRole)
	gouse.Println("Platform Family: ", hostInfo.PlatformFamily)
	gouse.Println("Platform Version: ", hostInfo.PlatformVersion)
}
```

## 4. Os partion

Description: Get OS Partition information<br>Params: None -> *gouse.IPartition<br>

```go
func OsPartion() {
	partitions, err := gouse.Partition()
	if err != nil {
		gouse.Println("Error: ", err)
		return
	}

	gouse.Println(":: Partitions Information")
	for _, v := range partitions {
		gouse.Println("Device: ", v.Device)
		gouse.Println("Mountpoint: ", v.Mountpoint)
		gouse.Println("Fstype: ", v.Fstype)
		gouse.Println("Opts: ", v.Opts)
	}
}
```

## 5. Os user

Description: Get OS Process information<br>Params: None -> *gouse.IProcess<br>

```go
func OsUser() {
	users, err := gouse.User()
	if err != nil {
		gouse.Println("Error: ", err)
		return
	}

	gouse.Println(":: Users Information")
	for _, v := range users {
		gouse.Println("User: ", v.User)
		gouse.Println("Terminal: ", v.Terminal)
		gouse.Println("Host: ", v.Host)
		gouse.Println("Started: ", v.Started)
	}
}
```

## 6. Os i o

Description: Get OS I/O information<br>Params: None -> *gouse.IIo<br>

```go
func OsIO() {
	ioStat, err := gouse.Io()
	if err != nil {
		gouse.Println("Error: ", err)
		return
	}

	gouse.Println(":: IO Information")
	for _, v := range ioStat {
		gouse.Println("Read Count: ", v.ReadCount)
		gouse.Println("Merge Read Count: ", v.MergedReadCount)
		gouse.Println("Write Count: ", v.WriteCount)
		gouse.Println("Merge Write Count: ", v.MergedWriteCount)
		gouse.Println("Read Bytes: ", v.ReadBytes)
		gouse.Println("Write Bytes: ", v.WriteBytes)
		gouse.Println("Read Time: ", v.ReadTime)
		gouse.Println("Write Time: ", v.WriteTime)
		gouse.Println("Iops In Progress: ", v.IopsInProgress)
		gouse.Println("I/O Time: ", v.IoTime)
		gouse.Println("Weighted IO: ", v.WeightedIO)
		gouse.Println("Name: ", v.Name)
		gouse.Println("Serial Number: ", v.SerialNumber)
		gouse.Println("Label: ", v.Label)
	}
}
```

## 7. Os memory

Description: Get Memory information<br>Params: None -> *gouse.IMemory<br>

```go
func OsMemory() {
	memoryInfo, err := gouse.Memory()
	if err != nil {
		gouse.Println("Error: ", err)
		return
	}

	gouse.Println(":: Memory Information")
	gouse.Println("Virtual Total: ", memoryInfo.VirtualTotal, "MB")
	gouse.Println("Virtual Used: ", memoryInfo.VirtualUsed, "MB")
	gouse.Println("Virtual Free: ", memoryInfo.VirtualFree, "MB")
	gouse.Println("Virtual Used Percent: ", memoryInfo.VirtualUsedPercent, "%")
	gouse.Println("Virtual Available: ", memoryInfo.VirtualAvailable, "MB")
	gouse.Println("Virtual Active: ", memoryInfo.VirtualActive, "MB")
	gouse.Println("Virtual Inactive: ", memoryInfo.VirtualInactive, "MB")
	gouse.Println("Vrtual Wired: ", memoryInfo.VirtualWired, "MB")
	gouse.Println("Virtual Laundry: ", memoryInfo.VirtualLaundry, "MB")
	gouse.Println("Virtual Buffers: ", memoryInfo.VirtualBuffers, "MB")
	gouse.Println("Virtual Cached: ", memoryInfo.VirtualCached, "MB")
	gouse.Println("Virtual Write Back: ", memoryInfo.VirtualWriteBack, "MB")
	gouse.Println("Virtual Dirty: ", memoryInfo.VirtualDirty, "MB")
	gouse.Println("Virtual Write Back Tmp: ", memoryInfo.VirtualWriteBackTmp, "MB")
	gouse.Println("Virtual Shared: ", memoryInfo.VirtualShared, "MB")
	gouse.Println("Virtual Slab: ", memoryInfo.VirtualSlab, "MB")
	gouse.Println("Virtual Sreclaimable: ", memoryInfo.VirtualSreclaimable, "MB")
	gouse.Println("Virtual Sunreclaim: ", memoryInfo.VirtualSunreclaim, "MB")
	gouse.Println("Virtual Page Tables: ", memoryInfo.VirtualPageTables, "MB")
	gouse.Println("Virtual Swap Cached: ", memoryInfo.VirtualSwapCached, "MB")
	gouse.Println("Virtual Commit Limit: ", memoryInfo.VirtualCommitLimit, "MB")
	gouse.Println("Virtual Committed AS: ", memoryInfo.VirtualCommittedAS, "MB")
	gouse.Println("Virtual High Total: ", memoryInfo.VirtualHighTotal, "MB")
	gouse.Println("Virtual High Free: ", memoryInfo.VirtualHighFree, "MB")
	gouse.Println("Virtual Low Total: ", memoryInfo.VirtualLowTotal, "MB")
	gouse.Println("Virtual Low Free: ", memoryInfo.VirtualLowFree, "MB")
	gouse.Println("Virtual Mapped: ", memoryInfo.VirtualMapped, "MB")
	gouse.Println("Virtual Malloc Total: ", memoryInfo.VirtualMallocTotal, "MB")
	gouse.Println("Virtual Malloc Used: ", memoryInfo.VirtualMallocUsed, "MB")
	gouse.Println("Virtual Malloc Chunk: ", memoryInfo.VirtualMallocChunk, "MB")
	gouse.Println("Virtual Huge Pages Total: ", memoryInfo.VirtualHugePagesTotal, "MB")
	gouse.Println("Virtual Huge Pages Free: ", memoryInfo.VirtualHugePagesFree, "MB")
	gouse.Println("Virtual Huge Pages Rsvd: ", memoryInfo.VirtualHugePagesRsvd, "MB")
	gouse.Println("Virtual Huge Pages Surp: ", memoryInfo.VirtualHugePagesSurp, "MB")
	gouse.Println("Virtual Huge Page Size: ", memoryInfo.VirtualHugePageSize, "MB")
	gouse.Println("Virtual Anon Huge Pages: ", memoryInfo.VirtualAnonHugePages, "MB")

	gouse.Println("Swap Total: ", memoryInfo.SwapTotal, "MB")
	gouse.Println("Swap Used: ", memoryInfo.SwapUsed, "MB")
	gouse.Println("Swap Free: ", memoryInfo.SwapFree, "MB")
	gouse.Println("Swap Used Percent: ", memoryInfo.SwapUsedPercent, "%")
	gouse.Println("Swap Sin: ", memoryInfo.SwapSin, "MB")
	gouse.Println("Swap Sout: ", memoryInfo.SwapSout, "MB")
	gouse.Println("Swap Pg In: ", memoryInfo.SwapPgIn, "MB")
	gouse.Println("Swap Pg Out: ", memoryInfo.SwapPgOut, "MB")
	gouse.Println("Swap Pg Fault: ", memoryInfo.SwapPgFault, "MB")
	gouse.Println("Swap Pg Maj Fault: ", memoryInfo.SwapPgMajFault, "MB")
}
```

## 8. Profile

Description: Run Profile<br>Input params: (cpu.pprof, mem.pprof)<br>

```go
func Profile() {
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
```
