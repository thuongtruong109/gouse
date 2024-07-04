package cron

import (
	"github.com/thuongtruong109/gouse/cron"
	"fmt"
)

func SampleCronRun() {
	callback := func() {
		fmt.Println("Cron job executed")
	}

	cron.Run(1, 5, callback)

	fmt.Println("Cron job stopped.")
}