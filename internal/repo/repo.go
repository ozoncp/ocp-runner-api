package repo

import (
	"context"
	"github.com/ozoncp/ocp-runner-api/internal/models"
	"github.com/rs/zerolog/log"
)

type Repo interface {
	AddRunner(ctx context.Context, runner *models.Runner) error
	AddRunners(ctx context.Context, runner []*models.Runner) error
	UpdateRunner(_ context.Context, runner *models.Runner) error
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
	log.Info().Str("added new runner", runner.String()).Send()
	return nil
}

// AddRunners adds bunch of new runners
func (*repo) AddRunners(_ context.Context, runners []*models.Runner) error {
	for _, runner := range runners {
		log.Info().Str("added new runner", runner.String()).Send()
	}
	return nil
}

// UpdateRunner updates runner data
func (*repo) UpdateRunner(_ context.Context, runner *models.Runner) error {
	log.Info().Str("updated runner", runner.String()).Send()
	return nil
}

// RemoveRunner removes runner by unique guid
func (*repo) RemoveRunner(_ context.Context, guid string) error {
	log.Info().Str("removed runner", guid).Send()
	return nil
}

// ListRunners returns list of runners
func (*repo) ListRunners(_ context.Context) ([]*models.Runner, error) {
	log.Info().Msg("listed all runners")
	return make([]*models.Runner, 10), nil
}
