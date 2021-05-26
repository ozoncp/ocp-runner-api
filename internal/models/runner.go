package models

import (
	"fmt"
	"github.com/google/uuid"
	"runtime"
)

type Runner struct {
	guid string
	os   string
	arch string
}

// New creates new Runner instance
func New(guid string, os string, arch string) *Runner {
	return &Runner{guid, os, arch}
}

// NewDefault creates new Runner instance with system values
func NewDefault() *Runner {
	guid, _ := uuid.NewUUID()
	return &Runner{guid.String(), runtime.GOOS, runtime.GOARCH}
}

// GetGuid returns current runner guid
func (runner *Runner) GetGuid() string {
	return runner.guid
}

// GetOs returns current runner os
func (runner *Runner) GetOs() string {
	return runner.os
}

// GetArch returns current runner arch
func (runner *Runner) GetArch() string {
	return runner.arch
}

// String returns formatted info
func (runner *Runner) String() string {
	return fmt.Sprintf("Runner (guid: %v, os: %v, arch: %v)", runner.guid, runner.os, runner.arch)
}
