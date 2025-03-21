
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Function' />


```go
import (
	"time"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Function delay

Description: Delay function execution for a specific time.<br>Input params: (func() string, delay int)<br>

```go
func FunctionDelay() {
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

## 2. Function interval

Description: Execute function at a specific interval.<br>Input params: (func(), interval int)<br>

```go
func FunctionInterval() {
	gouse.IntervalFunc(func() {
		println("Interval")
	}, 1)
}
```

## 3. Function lock

Description: Lock function execution.<br>

```go
func FunctionLock() {
	oneInOneOutLockFunc := gouse.LockFunc(func(i any) any {
		return i
	}).(func(any) any)("one input - one output")
	gouse.Println(oneInOneOutLockFunc)

	twoInTwoOutLockFunc1, twoInTwoOutLockFunc2 := gouse.LockFunc(func(i1, i2 any) (any, any) {
		return i1, i2
	}).(func(any, any) (any, any))("two input - two output (a)", "two input - two output (b)")
	gouse.Println(twoInTwoOutLockFunc1, twoInTwoOutLockFunc2)

	gouse.LockFunc(func() {
		println("no input - no output")
	}).(func())()

	exampleRWLockFunc := func(a, b int) int {
		return a + b
	}

	lockedFunc := gouse.RWLockFunc(exampleRWLockFunc).(func(int, int) int)
	result := lockedFunc(5, 3)
	gouse.Println("RW Lock function result:", result)
}
```

## 4. Function parallel

Description: Parallelize multi functions execution.<br>Input params: (...func())<br>

```go
func FunctionParallel() {
	function1 := func() {
		for range 5 {
			time.Sleep(100 * time.Millisecond)
			gouse.Println("Function 1 is executing")
		}
	}

	function2 := func() {
		for range 5 {
			time.Sleep(200 * time.Millisecond)
			gouse.Println("Function 2 is executing")
		}
	}

	function3 := func() {
		for range 5 {
			time.Sleep(300 * time.Millisecond)
			gouse.Println("Function 3 is executing")
		}
	}

	gouse.ParallelizeFunc(function1, function2, function3)
}
```

## 5. Function remain

Description: Limit function execution.<br>Input params: (func(), times int)<br>

```go
func FunctionRemain() {
	gouse.RemainFunc(func() {
		println("Times")
	}, 3)
}
```

## 6. Function retry

Description: Retry function execution.<br>Input params: (func(), times int, delay int)<br>

```go
func FunctionRetry() {
	gouse.RetryFunc(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}
```

## 7. Function run time

Description: Measure function execution time.<br>Input params: (time.Time, func())<br>

```go
func FunctionRunTime() {
	exampleFunc := func() {
		time.Sleep(2 * time.Second)
		gouse.Println("Task completed.")
	}

	duration := gouse.RunTimeFunc(time.Now(), exampleFunc)
	gouse.Printf("Function run in: %v\n", duration)
}
```

## 8. Function defer wrapper

Description: Wrap defer function<br>Input params: (...func())<br>

```go
func FunctionDeferWrapper() {
	gouse.DeferWrapper(
		func() error {
			gouse.Println("Connecting to database...")
			return nil
		},
		func() {
			gouse.Println("Disconnecting from database...")
		},
	)
}
```
