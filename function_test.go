package gouse

import (
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
