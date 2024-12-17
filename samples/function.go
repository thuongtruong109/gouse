package samples

import (
	"fmt"
	"time"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Delay function execution for a specific time.
Input params: (func() string, delay int)
*/
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

/*
Description: Execute function at a specific interval.
Input params: (func(), interval int)
*/
func FunctionInterval() {
	gouse.IntervalFunc(func() {
		println("Interval")
	}, 1)
}

/*
Description: Lock function execution.
*/
func FunctionLock() {
	oneInOneOutLockFunc := gouse.LockFunc(func(i interface{}) interface{} {
		return i
	}).(func(interface{}) interface{})("one input - one output")
	gouse.Println(oneInOneOutLockFunc)

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

/*
Description: Parallelize multi functions execution.
Input params: (...func())
*/
func FunctionParallel() {
	function1 := func() {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			gouse.Println("Function 1 is executing")
		}
	}

	function2 := func() {
		for i := 0; i < 5; i++ {
			time.Sleep(200 * time.Millisecond)
			gouse.Println("Function 2 is executing")
		}
	}

	function3 := func() {
		for i := 0; i < 5; i++ {
			time.Sleep(300 * time.Millisecond)
			gouse.Println("Function 3 is executing")
		}
	}

	gouse.ParallelizeFunc(function1, function2, function3)
}

/*
Description: Limit function execution.
Input params: (func(), times int)
*/
func FunctionRemain() {
	gouse.RemainFunc(func() {
		println("Times")
	}, 3)
}

/*
Description: Retry function execution.
Input params: (func(), times int, delay int)
*/
func FunctionRetry() {
	gouse.RetryFunc(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}

/*
Description: Measure function execution time.
Input params: (time.Time, func())
*/
func FunctionRunTime() {
	exampleFunc := func() {
		time.Sleep(2 * time.Second)
		gouse.Println("Task completed.")
	}

	duration := gouse.RunTimeFunc(time.Now(), exampleFunc)
	gouse.Printf("Function run in: %v\n", duration)
}

/*
Description: Wrap defer function
Input params: (...func())
*/
func FuncDeferWrapper() {
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
