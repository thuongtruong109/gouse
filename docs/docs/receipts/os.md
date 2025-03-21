
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Os' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Cpu

Description: Get CPU information<br>

```go
func Cpu() {
	cpuInfo, err := gouse.CPU()
	if err != nil {
		gouse.Println("Error: ", err)
		return
	}

	gouse.Println("CPU Information")
	gouse.Println("Model Name: ", cpuInfo.ModelName)
	gouse.Println("Family: ", cpuInfo.Family)
	gouse.Println("Speed: ", cpuInfo.Speed, "MHz")
	gouse.Println("Cores: ", cpuInfo.Cores)
}
```

## 2. Disk

Description: Get Disk information<br>

```go
func Disk() {
	diskInfo, err := gouse.Disk()
	if err != nil {
		gouse.Println("Error: ", err)
		return
	}

	gouse.Println("Disk Information")
	gouse.Println("Total Space: ", diskInfo.Total, "MB")
	gouse.Println("Free Space: ", diskInfo.Free, "MB")
	gouse.Println("Used Space: ", diskInfo.Used, "MB")
	gouse.Println("Used Space Percent: ", diskInfo.UsedPercent, "%")
}
```

## 3. System

Description: Get Memory information<br>

```go
func System() {
	systemInfo, err := gouse.System()
	if err != nil {
		gouse.Println("Error: ", err)
		return
	}

	gouse.Println("System Information")
	gouse.Println("OS: ", systemInfo.Name)
	gouse.Println("Platform: ", systemInfo.Platform)
	gouse.Println("Arch: ", systemInfo.Arch)
	gouse.Println("Hostname: ", systemInfo.Hostname)
	gouse.Println("Number of Processes: ", systemInfo.NumsProcs)
	// gouse.Println("Total Memory: ", systemInfo.TotalMemory, "MB")
	// gouse.Println("Free Memory: ", systemInfo.FreeMemory, "MB")
	// gouse.Println("Used Memory: ", systemInfo.UsedMemory, "MB")
	// gouse.Println("Used Memory Percent: ", systemInfo.UsedMemoryPercent, "%")
}
```

## 4. Profile

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
