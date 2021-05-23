package models

import (
	"fmt"
	"github.com/google/uuid"
	"runtime"
)

type Runner struct {
	Guid string
	Os   string
	Arch string
}

// NewRunner creates new Runner instance
func NewRunner() *Runner {
	return &Runner{}
}

// Init process instance fields initialization
func (runner *Runner) Init(guid string, os string, arch string) {
	runner.Guid = guid
	runner.Os = os
	runner.Arch = arch
}

// InitDefaults process instance fields initialization by default values
func (runner *Runner) InitDefaults() {
	guid, _ := uuid.NewUUID()
	runner.Init(guid.String(), runtime.GOOS, runtime.GOARCH)
}

// String returns formatted info
func (runner *Runner) String() string {
	return fmt.Sprintf("Runner (guid: %v, os: %v, arch: %v)", runner.Guid, runner.Os, runner.Arch)
}
