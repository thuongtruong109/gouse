# Cron

```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

#### 1. SampleCronRun

```go
func SampleCronRun() {
	callback := func() {
		fmt.Println("Cron job executed")
	}

	gouse.RunJob(1, 5, callback)

	fmt.Println("Cron job stopped.")
}
```
