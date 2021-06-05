package time

import "time"

type Clock interface {
	Now() time.Time
}

// NewClock constructor
func NewClock() Clock {
	return &clock{}
}

type clock struct{}

// Now returns current time
func (clock) Now() time.Time {
	return time.Now()
}
