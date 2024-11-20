package gouse

import (
	"testing"
	"time"
)

func TestNewCronJob(t *testing.T) {
	duration := time.Second * 1
	stopAfter := time.Second * 5
	callback := func() {}

	cronJob := NewCronJob(duration, stopAfter, callback)

	if cronJob.duration != duration {
		t.Errorf("Expected duration %v, but got %v", duration, cronJob.duration)
	}
	if cronJob.stopAfter != stopAfter {
		t.Errorf("Expected stopAfter %v, but got %v", stopAfter, cronJob.stopAfter)
	}
	if cronJob.callback == nil {
		t.Error("Expected callback to be set, but it was nil")
	}
	if cronJob.stopChan == nil {
		t.Error("Expected stopChan to be set, but it was nil")
	}
}

func TestCronJobStartAndWait(t *testing.T) {
	duration := time.Millisecond * 10
	stopAfter := time.Millisecond * 50
	callbackCalled := false

	callback := func() {
		callbackCalled = true
	}

	cronJob := NewCronJob(duration, stopAfter, callback)
	cronJob.StartJob()
	cronJob.WaitJob()

	if !callbackCalled {
		t.Error("Expected callback to be called, but it wasn't")
	}
}

func TestRunJob(t *testing.T) {
	duration := uint64(1)
	stopAfter := uint64(2)
	callbackCalled := false

	callback := func() {
		callbackCalled = true
	}

	RunJob(duration, stopAfter, callback)

	if !callbackCalled {
		t.Error("Expected callback to be called, but it wasn't")
	}
}
