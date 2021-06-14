package models

import (
	"fmt"
)

type Runner struct {
	Guid string
	Os   string
	Arch string
}

// String returns formatted info
func (runner *Runner) String() string {
	return fmt.Sprintf("Runner (guid: %v, os: %v, arch: %v)", runner.Guid, runner.Os, runner.Arch)
}
