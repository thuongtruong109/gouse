
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.25rem 0.75rem 0.25rem 0;' type='info' text='🔖 Function' />


```go
import (
	"fmt"
	"time"
	"github.com/thuongtruong109/gouse"
)
```

### <Badge style='font-size: 1.1rem;' type='tip' text='1. sample func delay' />



```go
func SampleFuncDelay() {
	println("Delay start:")

	result := gouse.DelayF(func() string {
		return "Delayed return after 2s"
	}, 2)

	if result.HasReturn {
		println(result.Value)
	} else {
		println("No result")
	}

	gouse.DelayFunc(func() {
		println("Delayed not return")
	}, 3)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='2. sample func interval' />



```go
func SampleFuncInterval() {
	gouse.IntervalFunc(func() {
		println("Interval")
	}, 1)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='3. sample func lock' />



```go
func SampleFuncLock() {
	oneInOneOutLockFunc := gouse.LockFunc(func(i interface{}) interface{} {
		return i
	}).(func(interface{}) interface{})("one input - one output")
	fmt.Println(oneInOneOutLockFunc)

	twoInTwoOutLockFunc1, twoInTwoOutLockFunc2 := gouse.LockFunc(func(i1, i2 interface{}) (interface{}, interface{}) {
		return i1, i2
	}).(func(interface{}, interface{}) (interface{}, interface{}))("two input - two output (a)", "two input - two output (b)")
	fmt.Println(twoInTwoOutLockFunc1, twoInTwoOutLockFunc2)

	gouse.LockFunc(func() {
		println("no input - no output")
	}).(func())()

	exampleRWLockFunc := func(a, b int) int {
		return a + b
	}

	lockedFunc := gouse.RWLockFunc(exampleRWLockFunc).(func(int, int) int)
	result := lockedFunc(5, 3)
	fmt.Println("RW Lock function result:", result)

}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='4. sample func parallel' />



```go
func SampleFuncParallel() {
	function1 := func() {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Function 1 is executing")
		}
	}

	function2 := func() {
		for i := 0; i < 5; i++ {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Function 2 is executing")
		}
	}

	function3 := func() {
		for i := 0; i < 5; i++ {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("Function 3 is executing")
		}
	}

	gouse.ParallelizeFunc(function1, function2, function3)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='5. sample func remain' />



```go
func SampleFuncRemain() {
	gouse.RemainFunc(func() {
		println("Times")
	}, 3)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='6. sample func retry' />



```go
func SampleFuncRetry() {
	gouse.RetryFunc(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}
```

### <Badge style='font-size: 1.1rem;' type='tip' text='7. sample func run time' />



```go
func SampleFuncRunTime() {
	exampleFunc := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task completed.")
	}

	duration := gouse.RunTimeFunc(time.Now(), exampleFunc)
	fmt.Printf("Function run in: %v\n", duration)
}
```
