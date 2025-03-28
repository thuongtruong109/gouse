
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Context' />


```go
import (
	"fmt"
	"time"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Context

Description: Management of context and cancellation in Go.<br>

```go
func Context() {
	ctx, cancel := gouse.CtxCancel()
	gouse.SetCtx(ctx)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Process cancelled.")
				return
			default:
				// Simulate some work
				fmt.Println("Working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Simulating program running and waiting for termination signal
	time.Sleep(5 * time.Second)

	// After 5 seconds, trigger cancellation
	cancel()

	// Wait for the program to exit
	time.Sleep(2 * time.Second)

	// Retrieve the context and print if it's cancelled
	fmt.Println("Is context cancelled?", gouse.GetCtx().Done() != nil)
}
```
