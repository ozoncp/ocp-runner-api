package models

import (
	"fmt"
)

type Runner struct {
	Guid string
	Os   string
	Arch string
}

// GetGuid returns current runner guid
func (runner *Runner) GetGuid() string {
	return runner.Guid
}

// GetOs returns current runner os
func (runner *Runner) GetOs() string {
	return runner.Os
}

// GetArch returns current runner arch
func (runner *Runner) GetArch() string {
	return runner.Arch
}

// String returns formatted info
func (runner *Runner) String() string {
	return fmt.Sprintf("Runner (guid: %v, os: %v, arch: %v)", runner.Guid, runner.Os, runner.Arch)
}
