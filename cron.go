package gouse

import "time"

type ICron struct {
	duration  time.Duration
	stopAfter time.Duration
	callback  func()
	stopChan  chan bool
}

func NewCron(duration, stopAfter time.Duration, callback func()) *ICron {
	return &ICron{
		duration:  duration,
		stopAfter: stopAfter,
		callback:  callback,
		stopChan:  make(chan bool),
	}
}

func (cj *ICron) StartJob() {
	go func() {
		ticker := time.NewTicker(cj.duration)
		stopTimer := time.NewTimer(cj.stopAfter)

		defer ticker.Stop()
		defer stopTimer.Stop()

		for {
			select {
			case <-ticker.C:
				cj.callback()
			case <-stopTimer.C:
				cj.stopChan <- true
				return
			}
		}
	}()
}

func (cj *ICron) WaitJob() {
	<-cj.stopChan
}

func RunJob(duration, stopAfter uint64, callback func()) {
	cronJob := NewCron(time.Duration(duration)*time.Second, time.Duration(stopAfter)*time.Second, callback)
	cronJob.StartJob()
	cronJob.WaitJob()
}
