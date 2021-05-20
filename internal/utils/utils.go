package utils

import (
	"errors"
	"fmt"
	"github.com/pkg/math"
)

// SplitSlice split the source slice into the slice of slices
func SplitSlice(source []interface{}, batchSize int) [][]interface{} {
	var result [][]interface{}
	var from, to int

	for {
		if from == len(source) {
			break
		}
		to = math.Min(from + batchSize, len(source))
		result = append(result, source[from: to])
		from = to
	}

	return result
}

// SwapKeyValue returns new map where keys are swapped with values
func SwapKeyValue(source map[int]string) (result map[string]int, err error) {
	result = make(map[string]int, len(source))

	defer func() {
		if ex := recover(); ex != nil {
			err = errors.New(ex.(string))
			result = nil
		}
	}()

	for key, value := range source {
		if _, found := result[value]; found {
			panic(fmt.Sprintf("Key %v already exist in result map!", value))
		}
		result[value] = key
	}

	return result, err
}

// SliceExcept returns new slice without specified values
func SliceExcept(source []string, except []string) []string {
	var result []string

	for _, value := range source {
		if !contains(value, except) {
			result = append(result, value)
		}
	}

	return result
}

func contains(v string, s []string) bool {
	for _, item := range s {
		if v == item {
			return true
		}
	}
	return false
}