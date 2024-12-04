package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Run a cron job
Input params: (duration uint64, stopAfter uint64, callback func())
*/
func CronRun() {
	callback := func() {
		fmt.Println("Cron job executed")
	}

	gouse.RunJob(1, 5, callback)

	fmt.Println("Cron job stopped.")
}
