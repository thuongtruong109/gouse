package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

func SampleCronRun() {
	callback := func() {
		fmt.Println("Cron job executed")
	}

	gouse.RunJob(1, 5, callback)

	fmt.Println("Cron job stopped.")
}
