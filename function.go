package gouse

import (
	"reflect"
	"sync"
	"time"
)

/* Lock */

type MutexWrapper struct {
	mutex sync.Mutex
}

func (mw *MutexWrapper) lock() {
	mw.mutex.Lock()
}

func (mw *MutexWrapper) unLock() {
	mw.mutex.Unlock()
}

func LockFunc(callback interface{}) interface{} {
	m := MutexWrapper{}
	m.lock()
	defer m.unLock()

	callbackType := reflect.TypeOf(callback)
	if callbackType.Kind() != reflect.Func {
		panic("callback must be a function")
	}

	return reflect.MakeFunc(callbackType, func(params []reflect.Value) []reflect.Value {
		// Convert params to interface{} slice
		var resultInterfaces []interface{}
		for _, result := range reflect.ValueOf(callback).Call(params) {
			resultInterfaces = append(resultInterfaces, result.Interface())
		}

		// Return the results as reflect.Value
		var resultReflectValues []reflect.Value
		for _, result := range resultInterfaces {
			resultReflectValues = append(resultReflectValues, reflect.ValueOf(result))
		}

		return resultReflectValues
	}).Interface()
}

// RWLock

// type RWMutexWrapper struct {
// 	rwMutex sync.RWMutex
// }

// func (mw *RWMutexWrapper) rLock() {
// 	mw.rwMutex.RLock()
// }

// func (mw *RWMutexWrapper) rUnLock() {
// 	mw.rwMutex.RUnlock()
// }

/* Delay */

type DelayedResult[T any] struct {
	Value     T
	HasReturn bool
}

func delay[T any](f func() T, timeout int, hasReturn bool) DelayedResult[T] {
	resultChan := make(chan T, 1)

	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		result := f()
		if hasReturn {
			resultChan <- result
		}
	}()

	if hasReturn {
		return DelayedResult[T]{Value: <-resultChan, HasReturn: true}
	}

	return DelayedResult[T]{HasReturn: false}
}

func DelayF[T any](f func() T, timeout int) DelayedResult[T] {
	return delay(f, timeout, true)
}

func DelayFunc(f func(), timeout int) DelayedResult[struct{}] {
	return delay(func() struct{} {
		f()
		return struct{}{}
	}, timeout, false)
}

/* Utilities */

func ParallelizeFunc(functions ...func()) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(functions))

	ch := make(chan struct{}, len(functions))
	done := make(chan struct{}, len(functions))

	for _, f := range functions {
		ch <- struct{}{}
		go func(copyFunc func()) {
			defer func() {
				<-ch
				waitGroup.Done()
			}()
			copyFunc()
		}(f)
	}

	go func() {
		waitGroup.Wait()
		close(done)
	}()

	<-done
}

func RetryFunc(fn func() error, attempts int, sleep int) (err error) {
	retry := func() {
		if err = fn(); err != nil {
			return
		}
	}

	for i := attempts; i > 0; i-- {
		if i > 1 {
			retry()
			SleepSecond(sleep)
		} else if i == 1 {
			retry()
		} else {
			return nil
		}
	}
	return
}

func RemainFunc(fn func(), attempts int) {
	for i := 0; i < attempts; i++ {
		fn()
	}
}

func IntervalFunc(fn func(), timeout int) {
	for {
		fn()
		SleepSecond(timeout)
	}
}

func RunTimeFunc(startTime time.Time, task func()) time.Duration {
	task()
	elapsedTime := float64(time.Since(startTime).Seconds() * 1000)
	return time.Duration(elapsedTime)
}
