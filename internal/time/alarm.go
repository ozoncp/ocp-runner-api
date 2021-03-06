package time

import (
	"context"
	"time"
)

type Alarm interface {
	Init()
	Alarm() <-chan struct{}
	Close()
}

// NewAlarm constructor
func NewAlarm(ctx context.Context, timeout time.Duration, clock Clock) Alarm {
	return &alarm{
		alarms:  make(chan struct{}),
		done:    make(chan struct{}),
		ctx:     ctx,
		timeout: timeout,
		clock:   clock,
	}
}

type alarm struct {
	alarms  chan struct{}
	done    chan struct{}
	ctx     context.Context
	timeout time.Duration
	clock   Clock
}

// Init runs alarm go-routine
func (a *alarm) Init() {
	go func() {
		timer := time.After(a.timeout)
		for {
			select {
			case <-timer:

				a.alarms <- struct{}{}
				timer = time.After(a.timeout)
			case <-a.ctx.Done():
				close(a.done)
				return
			}
		}
	}()
}

// Alarm returns notifications channel
func (a *alarm) Alarm() <-chan struct{} {
	return a.alarms
}

// Close quits the alarms
func (a *alarm) Close() {
	<-a.done
	close(a.alarms)
}
