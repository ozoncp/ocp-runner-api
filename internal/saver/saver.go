package saver

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-runner-api/internal/flusher"
	"github.com/ozoncp/ocp-runner-api/internal/models"
	"github.com/ozoncp/ocp-runner-api/internal/time"
)

type Saver interface {
	Init(ctx context.Context)
	Save(ctx context.Context, runner *models.Runner) error
	Close()
}

// NewSaver constructor
func NewSaver(capacity int, alarm time.Alarm, flusher flusher.Flusher) Saver {
	return &saver{
		capacity: capacity,
		started:  make(chan struct{}),
		runners:  make(chan *models.Runner),
		done:     make(chan struct{}),
		alarm:    alarm,
		flusher:  flusher,
	}
}

type saver struct {
	capacity int
	started  chan struct{}
	runners  chan *models.Runner
	done     chan struct{}
	alarm    time.Alarm
	flusher  flusher.Flusher
}

// Init initializes saver instance
func (s *saver) Init(ctx context.Context) {
	go s.flushing(ctx)
	<-s.started
}

// Save saves the runner
func (s *saver) Save(_ context.Context, runner *models.Runner) error {
	select {
	case s.runners <- runner:
		return nil
	default:
		return errors.New("save after Close() method called")
	}
}

// flushing flushes runners by alarm signals
func (s *saver) flushing(ctx context.Context) {
	var runners []*models.Runner

	readyToFlush := make(chan struct{}, 1)
	readyToFlush <- struct{}{}

	alarms := s.alarm.Alarm()

	for {
		select {
		case runner := <-s.runners:
			if len(runners) == s.capacity {
				runners = runners[1:]
			}
			runners = append(runners, runner)
		case <-alarms:
			runners = s.flusher.Flush(ctx, runners)
		case <-ctx.Done():
			_ = s.flusher.Flush(ctx, runners)
			close(s.done)
			return
		case <-readyToFlush:
			close(s.started)
		}
	}
}

// Close quits the saver
func (s *saver) Close() {
	<-s.done
	close(s.runners)
}
