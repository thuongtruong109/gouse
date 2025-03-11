package gouse

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestParallelizeFunc(t *testing.T) {
	done := make(chan struct{})

	func1 := func() {
		time.Sleep(1 * time.Second)
	}

	func2 := func() {
		time.Sleep(1 * time.Second)
	}

	go func() {
		ParallelizeFunc(func1, func2)
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Error("Parallelize(): timeout waiting to complete, maximum is 2 seconds")
	}
}

func TestLockFunc(t *testing.T) {
	var result int

	callback := func() {
		result = 42
	}

	lockFunc := LockFunc(callback)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		lockFunc.(func())()
	}()

	wg.Wait()

	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestRWLockFunc(t *testing.T) {
	var result int

	callback := func() {
		result = 84
	}

	rwLockFunc := RWLockFunc(callback)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		rwLockFunc.(func())()
	}()

	go func() {
		defer wg.Done()
		rwLockFunc.(func())()
	}()

	wg.Wait()

	if result != 84 {
		t.Errorf("expected 84, got %d", result)
	}
}

func TestLockHandlerWithNonFunctionCallback(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, but code did not panic")
		}
	}()

	_lockHandler("not a function", func() {}, func() {})
}

func TestLockFuncConcurrency(t *testing.T) {
	var count int
	var mutex sync.Mutex

	callback := func() {
		mutex.Lock()
		defer mutex.Unlock()
		count++
	}

	lockFunc := LockFunc(callback)

	var wg sync.WaitGroup
	goroutines := 100
	wg.Add(goroutines)

	for range goroutines {
		go func() {
			defer wg.Done()
			lockFunc.(func())()
		}()
	}

	wg.Wait()

	if count != goroutines {
		t.Errorf("expected %d, got %d", goroutines, count)
	}
}

func TestDeferWrapperSuccess(t *testing.T) {
	cleanupCalled := false
	cleanupFunc := func() {
		cleanupCalled = true
	}

	mainFunc := func() error {
		return nil
	}

	err := DeferWrapper(mainFunc, cleanupFunc)

	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if !cleanupCalled {
		t.Errorf("expected cleanup to be called, but it wasn't")
	}
}

func TestDeferWrapperFailure(t *testing.T) {
	cleanupCalled := false
	cleanupFunc := func() {
		cleanupCalled = true
	}

	mainFunc := func() error {
		return errors.New("main function error")
	}

	err := DeferWrapper(mainFunc, cleanupFunc)

	if err == nil || err.Error() != "main function error" {
		t.Errorf("expected 'main function error', but got %v", err)
	}
	if !cleanupCalled {
		t.Errorf("expected cleanup to be called, but it wasn't")
	}
}

func TestDeferWrapperNilCleanup(t *testing.T) {
	mainFunc := func() error {
		return nil
	}

	err := DeferWrapper(mainFunc, nil)

	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
}
