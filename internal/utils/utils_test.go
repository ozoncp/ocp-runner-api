package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	source := []interface{} {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	result := SplitSlice(source, 2)

	assert.Len(t, result, 6)
	assert.Len(t, result[0], 2)
	assert.Len(t, result[5], 1)
}

func TestSwapKeyValue(t *testing.T) {
	source := map[int]string {
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
	}

	result, err := SwapKeyValue(source)

	assert.Nil(t, err)
	for key, value := range result {
		assert.Equal(t, key, source[value])
		assert.Equal(t, value, getKey(source, key))
	}

	badSource := map[int]string {
		0: "banana",
		1: "lime",
		2: "melon",
		3: "banana",
	}

	badResult, err := SwapKeyValue(badSource)
	assert.Error(t, err)
	assert.Nil(t, badResult)
}

// getKey returns key by value
func getKey(source map[int]string, value string) int {
	for key, val := range source {
		if val == value {
			return key
		}
	}
	return -1
}

func TestSliceExcept(t *testing.T) {
	source := []string { "a", "b", "c", "d", "e", "g" }
	except := []string { "a", "q", "w", "e" }

	result := SliceExcept(source, except)

	assert.Len(t, result, 4)
	assert.Equal(t, []string { "b", "c", "d", "g" }, result)
}