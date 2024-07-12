package os

import (
	"fmt"
	"github.com/thuongtruong109/gouse/os"
)

func SampleOsCpu() {
	cpuInfo, err := os.GetCPU()
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