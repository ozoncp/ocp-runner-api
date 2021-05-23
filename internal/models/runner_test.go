package models

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestInit(t *testing.T) {
	runner := NewRunner()

	guid, _ := uuid.NewUUID()
	runner.Init(guid.String(), "windows", "x64")

	assert.Equal(t, guid.String(), runner.Guid)
	assert.Equal(t, "windows", runner.Os)
	assert.Equal(t, "x64", runner.Arch)
}

func TestInitDefaults(t *testing.T) {
	runner := NewRunner()

	runner.InitDefaults()

	assert.NotEmpty(t, runner.Guid)
	assert.Equal(t, runtime.GOOS, runner.Os)
	assert.Equal(t, runtime.GOARCH, runner.Arch)
}
