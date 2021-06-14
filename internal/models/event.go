package models

import "fmt"

type RunnerEventType int

const (
	Created RunnerEventType = iota
	Updated
	Removed
)

type RunnerEvent struct {
	Type RunnerEventType
	Body map[string]interface{}
}

func (re *RunnerEvent) String() string {
	return fmt.Sprintf("%v", re.getDisplay())
}

func (re *RunnerEvent) getDisplay() string {
	switch re.Type {
	case Created:
		return "created"
	case Updated:
		return "updated"
	case Removed:
		return "removed"
	default:
		return "unknown"
	}
}
