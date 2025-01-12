package socket

import (
	"fmt"
	"sync"
	"time"
)

type Percentable interface {
	Current() int64
	Add()
	To(percent int64)
}

type processPercent struct {
	ProcessPercentMsger

	s           Sleeper
	percentLock sync.RWMutex
	percent     int64
}

type Sleeper interface {
	Sleep(time.Duration)
}

type realSleeper struct{}

func (r *realSleeper) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func newProcessPercent(sender ProcessPercentMsger, s Sleeper) Percentable {
	return &processPercent{
		s:                   s,
		percent:             0,
		ProcessPercentMsger: sender,
	}
}

func (pp *processPercent) Current() int64 {
	pp.percentLock.RLock()
	defer pp.percentLock.RUnlock()

	return pp.percent
}

func (pp *processPercent) Add() {
	pp.percentLock.Lock()
	defer pp.percentLock.Unlock()

	if pp.percent < 100 {
		pp.percent++
		pp.SendProcessPercent(fmt.Sprintf("%d", pp.percent))
	}
}

func (pp *processPercent) To(percent int64) {
	pp.percentLock.Lock()
	defer pp.percentLock.Unlock()

	sleepTime := 100 * time.Millisecond
	for pp.percent < percent {
		pp.s.Sleep(sleepTime)
		pp.percent++
		if sleepTime > 50*time.Millisecond {
			sleepTime = sleepTime / 2
		}
		pp.SendProcessPercent(fmt.Sprintf("%d", pp.percent))
	}
}
