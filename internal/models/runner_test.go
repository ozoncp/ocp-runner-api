package models

import (
	"runtime"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	guid, _ := uuid.NewUUID()
	runner := New(guid.String(), "windows", "x64")

	assert.Equal(t, guid.String(), runner.GetGuid())
	assert.Equal(t, "windows", runner.GetOs())
	assert.Equal(t, "x64", runner.GetArch())
}

func TestInitDefaults(t *testing.T) {
	runner := NewDefault()

	assert.NotEmpty(t, runner.GetGuid())
	assert.Equal(t, runtime.GOOS, runner.GetOs())
	assert.Equal(t, runtime.GOARCH, runner.GetArch())
}
