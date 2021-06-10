package utils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ozoncp/ocp-runner-api/internal/models"
)

func TestSplitToBulks(t *testing.T) {
	source := make([]*models.Runner, 11)
	for i := 0; i < len(source)-1; i++ {
		source[i] = &models.Runner{}
	}

	result := SplitToBulks(source, 2)

	assert.Len(t, result, 6)
	assert.Len(t, result[0], 2)
	assert.Len(t, result[5], 1)
}

func TestGroupByGuid(t *testing.T) {
	opSystems := []string{"windows", "linux", "macOS"}

	// Good case
	var source []*models.Runner
	for _, system := range opSystems {
		guid, _ := uuid.NewUUID()
		runner := &models.Runner{
			Guid: guid.String(),
			Os:   system,
			Arch: "x64",
		}
		source = append(source, runner)
	}

	result, err := GroupByGuid(source)

	assert.Nil(t, err)
	assert.Len(t, result, 3)

	// Bad case
	var badSource []*models.Runner
	for _, system := range opSystems {
		runner := &models.Runner{
			Guid: "same-guid",
			Os:   system,
			Arch: "x64",
		}
		badSource = append(badSource, runner)
	}

	badResult, err := GroupByGuid(badSource)

	assert.Error(t, err)
	assert.Nil(t, badResult)

	// Empty case
	emptyResult, err := GroupByGuid(nil)

	assert.Nil(t, emptyResult)
	assert.Error(t, err)
}

func TestSliceExcept(t *testing.T) {
	source := []string{"a", "b", "c", "d", "e", "g"}
	except := []string{"a", "q", "w", "e"}

	result := SliceExcept(source, except)

	assert.Len(t, result, 4)
	assert.Equal(t, []string{"b", "c", "d", "g"}, result)
}
