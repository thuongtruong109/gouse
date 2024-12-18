package gouse

import "time"

type ICronJob struct {
	duration  time.Duration
	stopAfter time.Duration
	callback  func()
	stopChan  chan bool
}

func NewCronJob(duration, stopAfter time.Duration, callback func()) *ICronJob {
	return &ICronJob{
		duration:  duration,
		stopAfter: stopAfter,
		callback:  callback,
		stopChan:  make(chan bool),
	}
}

func (cj *ICronJob) StartJob() {
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

func (cj *ICronJob) WaitJob() {
	<-cj.stopChan
}

func RunJob(duration, stopAfter uint64, callback func()) {
	cronJob := NewCronJob(time.Duration(duration)*time.Second, time.Duration(stopAfter)*time.Second, callback)
	cronJob.StartJob()
	cronJob.WaitJob()
}
