package utils

import (
	"github.com/ozoncp/ocp-runner-api/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitToBulks(t *testing.T) {
	source := make([]*models.Runner, 11)
	for i := 0; i < len(source)-1; i++ {
		source[i] = models.NewRunner()
	}

	result := SplitToBulks(source, 2)

	assert.Len(t, result, 6)
	assert.Len(t, result[0], 2)
	assert.Len(t, result[5], 1)
}

func TestGroupByGuid(t *testing.T) {
	// Good case
	source := []*models.Runner{
		models.NewRunner(),
		models.NewRunner(),
		models.NewRunner(),
	}
	for _, runner := range source {
		runner.InitDefaults()
	}

	result, err := GroupByGuid(source)

	assert.Nil(t, err)
	assert.Len(t, result, 3)

	// Bad case
	badSource := []*models.Runner{
		models.NewRunner(),
		models.NewRunner(),
		models.NewRunner(),
	}
	for _, runner := range source {
		runner.Init("same-guid", "windows", "x64")
	}

	badResult, err := GroupByGuid(badSource)

	assert.Error(t, err)
	assert.Nil(t, badResult)
}

func TestSliceExcept(t *testing.T) {
	source := []string{"a", "b", "c", "d", "e", "g"}
	except := []string{"a", "q", "w", "e"}

	result := SliceExcept(source, except)

	assert.Len(t, result, 4)
	assert.Equal(t, []string{"b", "c", "d", "g"}, result)
}
