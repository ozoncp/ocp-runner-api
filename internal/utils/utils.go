package utils

import (
	"errors"
	"fmt"
	"github.com/ozoncp/ocp-runner-api/internal/models"
	"github.com/pkg/math"
)

// SplitToBulks split the source slice into the slice of slices
func SplitToBulks(source []*models.Runner, batchSize int) [][]*models.Runner {
	var result [][]*models.Runner
	var from, to int

	for {
		if from == len(source) {
			break
		}
		to = math.Min(from+batchSize, len(source))
		result = append(result, source[from:to])
		from = to
	}

	return result
}

// GroupByGuid returns map of runners grouped by unique guid
func GroupByGuid(runners []*models.Runner) (result map[string]*models.Runner, err error) {
	if runners == nil || len(runners) == 0 {
		return nil, errors.New("received nil or empty slice")
	}

	result = make(map[string]*models.Runner, len(runners))

	defer func() {
		if ex := recover(); ex != nil {
			err = errors.New(ex.(string))
			result = nil
		}
	}()

	for _, runner := range runners {
		if _, found := result[runner.Guid]; found {
			panic(fmt.Sprintf("Guid %v already exist in result map!", runner.Guid))
		}
		result[runner.Guid] = runner
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
