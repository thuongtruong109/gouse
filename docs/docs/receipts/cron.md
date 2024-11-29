
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.25rem 0.75rem 0.25rem 0;' type='info' text='🔖 Cron' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. sample cron run' />



```go
func SampleCronRun() {
	callback := func() {
		fmt.Println("Cron job executed")
	}

	gouse.RunJob(1, 5, callback)

	fmt.Println("Cron job stopped.")
}
```
