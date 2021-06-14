package models

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
