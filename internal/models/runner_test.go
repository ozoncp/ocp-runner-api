package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetters(t *testing.T) {
	guid, _ := uuid.NewUUID()
	runner := Runner{guid.String(), "windows", "x64"}

	assert.Equal(t, guid.String(), runner.GetGuid())
	assert.Equal(t, "windows", runner.GetOs())
	assert.Equal(t, "x64", runner.GetArch())
}
