package gouse

import (
	"reflect"
	"sync"
	"time"
)

/* Lock */

func lockHandler(callback any, lock, unlock func()) any {
	callbackType := reflect.TypeOf(callback)
	if callbackType.Kind() != reflect.Func {
		panic("callback must be a function")
	}

	return reflect.MakeFunc(callbackType, func(params []reflect.Value) []reflect.Value {
		lock()
		defer unlock()
		return reflect.ValueOf(callback).Call(params)
	}).Interface()
}

func LockFunc(callback any) any {
	var mutex sync.Mutex
	return lockHandler(callback, mutex.Lock, mutex.Unlock)
}

func RWLockFunc(callback any) any {
	var rwMutex sync.RWMutex
	return lockHandler(callback, rwMutex.RLock, rwMutex.RUnlock)
}

func DeferWrapper(mainFunc func() error, cleanupFunc func()) error {
	defer func() {
		if cleanupFunc != nil {
			cleanupFunc()
		}
	}()

	if err := mainFunc(); err != nil {
		return err
	}

	return nil
}

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
			SleepS(sleep)
		} else if i == 1 {
			retry()
		} else {
			return nil
		}
	}
	return
}

func RemainFunc(fn func(), attempts int) {
	for range attempts {
		fn()
	}
}

func IntervalFunc(fn func(), timeout int) {
	for {
		fn()
		SleepS(timeout)
	}
}

func RunTimeFunc(startTime time.Time, task func()) time.Duration {
	task()
	return DiffTimeNow(startTime)
}
