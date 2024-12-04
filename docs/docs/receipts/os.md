
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Os' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. os cpu' />



```go
func OsCpu() {
	cpuInfo, err := gouse.CPU()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("CPU Information")
	fmt.Println("Model Name: ", cpuInfo.ModelName)
	fmt.Println("Family: ", cpuInfo.Family)
	fmt.Println("Speed: ", cpuInfo.Speed, "MHz")
	fmt.Println("Cores: ", cpuInfo.Cores)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. os disk' />



```go
func OsDisk() {
	diskInfo, err := gouse.Disk()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Disk Information")
	fmt.Println("Total Space: ", diskInfo.TotalSpace, "MB")
	fmt.Println("Free Space: ", diskInfo.FreeSpace, "MB")
	fmt.Println("Used Space: ", diskInfo.UsedSpace, "MB")
	fmt.Println("Used Space Percent: ", diskInfo.UsedSpacePercent, "%")
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. os system' />



```go
func OsSystem() {
	systemInfo, err := gouse.System()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("System Information")
	fmt.Println("OS: ", systemInfo.Name)
	fmt.Println("Platform: ", systemInfo.Platform)
	fmt.Println("Arch: ", systemInfo.Arch)
	fmt.Println("Hostname: ", systemInfo.Hostname)
	fmt.Println("Number of Processes: ", systemInfo.NumsProcs)
	fmt.Println("Total Memory: ", systemInfo.TotalMemory, "MB")
	fmt.Println("Free Memory: ", systemInfo.FreeMemory, "MB")
	fmt.Println("Used Memory: ", systemInfo.UsedMemory, "MB")
	fmt.Println("Used Memory Percent: ", systemInfo.UsedMemoryPercent, "%")
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='4. os profile' />



```go
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
```
