package repo

import (
	"context"
	"fmt"
	"github.com/ozoncp/ocp-runner-api/internal/models"
)

type Repo interface {
	AddRunner(ctx context.Context, runner *models.Runner) error
	AddRunners(ctx context.Context, runner []*models.Runner) error
	RemoveRunner(ctx context.Context, guid string) error
	ListRunners(ctx context.Context) ([]*models.Runner, error)
}

type repo struct{}

// New creates new instance
func New() Repo {
	return &repo{}
}

// AddRunner adds new runner
func (*repo) AddRunner(_ context.Context, runner *models.Runner) error {
	fmt.Printf("added new runner (%v)\n", runner)
	return nil
}

// AddRunners adds bunch of new runners
func (*repo) AddRunners(_ context.Context, runners []*models.Runner) error {
	for _, runner := range runners {
		fmt.Printf("added new runner (%v)\n", runner)
	}
	return nil
}

// RemoveRunner removes runner by unique guid
func (*repo) RemoveRunner(_ context.Context, guid string) error {
	fmt.Printf("removed runner by guid %v\n", guid)
	return nil
}

// ListRunners returns list of runners
func (*repo) ListRunners(_ context.Context) ([]*models.Runner, error) {
	fmt.Println("select all runners")
	return make([]*models.Runner, 10), nil
}
