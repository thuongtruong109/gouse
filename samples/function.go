package samples

import (
	"fmt"
	"time"

	"github.com/thuongtruong109/gouse"
)

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

func SampleFuncInterval() {
	gouse.IntervalFunc(func() {
		println("Interval")
	}, 1)
}

func SampleFuncLock() {
	oneInOneOutFuc := gouse.LockFunc(func(i interface{}) interface{} {
		return i
	}).(func(interface{}) interface{})("one input - one output")
	fmt.Println(oneInOneOutFuc)

	twoInTwoOutFunc1, twoInTwoOutFunc2 := gouse.LockFunc(func(i1, i2 interface{}) (interface{}, interface{}) {
		return i1, i2
	}).(func(interface{}, interface{}) (interface{}, interface{}))("two input - two output (a)", "two input - two output (b)")
	fmt.Println(twoInTwoOutFunc1, twoInTwoOutFunc2)

	gouse.LockFunc(func() {
		println("no input - no output")
	}).(func())()
}

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

func SampleFuncRemain() {
	gouse.RemainFunc(func() {
		println("Times")
	}, 3)
}

func SampleFuncRetry() {
	gouse.RetryFunc(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}

func SampleFuncRunTime() {
	exampleFunc := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task completed.")
	}

	duration := gouse.RunTimeFunc(time.Now(), exampleFunc)
	fmt.Printf("Function run in: %v\n", duration)
}
