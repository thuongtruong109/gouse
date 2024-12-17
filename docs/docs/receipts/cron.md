
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Cron' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Cron run

Description: Run a cron job<br>Input params: (duration uint64, stopAfter uint64, callback func())<br>

```go
func CronRun() {
	callback := func() {
		gouse.Println("Cron job executed")
	}

	gouse.RunJob(1, 5, callback)

	gouse.Println("Cron job stopped.")
}
```
