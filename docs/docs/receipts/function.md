
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Function' />


```go
import (
	"fmt"
	"time"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Func delay



```go
func FuncDelay() {
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

## 2. Func interval



```go
func FuncInterval() {
	gouse.IntervalFunc(func() {
		println("Interval")
	}, 1)
}
```

## 3. Func lock



```go
func FuncLock() {
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

## 4. Func parallel



```go
func FuncParallel() {
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

## 5. Func remain



```go
func FuncRemain() {
	gouse.RemainFunc(func() {
		println("Times")
	}, 3)
}
```

## 6. Func retry



```go
func FuncRetry() {
	gouse.RetryFunc(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}
```

## 7. Func run time



```go
func FuncRunTime() {
	exampleFunc := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task completed.")
	}

	duration := gouse.RunTimeFunc(time.Now(), exampleFunc)
	fmt.Printf("Function run in: %v\n", duration)
}
```

## 8. Func defer wrapper



```go
func FuncDeferWrapper() {
	gouse.DeferWrapper(
		func() error {
			fmt.Println("Opening file...")
			return fmt.Errorf("failed to read file")
		},
		func() {
			fmt.Println("Closing file...")
		},
	)

	gouse.DeferWrapper(
		func() error {
			fmt.Println("Connecting to database...")
			return nil
		},
		func() {
			fmt.Println("Disconnecting from database...")
		},
	)
}
```
