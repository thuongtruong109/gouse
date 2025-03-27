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
	oneInOneOutLockFunc := gouse.LockFunc(func(i any) any {
		return i
	}).(func(any) any)("one input - one output")
	fmt.Println(oneInOneOutLockFunc)

	twoInTwoOutLockFunc1, twoInTwoOutLockFunc2 := gouse.LockFunc(func(i1, i2 any) (any, any) {
		return i1, i2
	}).(func(any, any) (any, any))("two input - two output (a)", "two input - two output (b)")
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
		for range 5 {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Function 1 is executing")
		}
	}

	function2 := func() {
		for range 5 {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Function 2 is executing")
		}
	}

	function3 := func() {
		for range 5 {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("Function 3 is executing")
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
		fmt.Println("Task completed.")
	}

	duration := gouse.RunTimeFunc(time.Now(), exampleFunc)
	fmt.Printf("Function run in: %v\n", duration)
}

/*
Description: Wrap defer function
Input params: (...func())
*/
func FunctionDeferWrapper() {
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
