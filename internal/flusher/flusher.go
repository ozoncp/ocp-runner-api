package flusher

import (
	"context"
	"github.com/ozoncp/ocp-runner-api/internal/models"
	"github.com/ozoncp/ocp-runner-api/internal/repo"
	"github.com/pkg/math"
)

type Flusher interface {
	Flush(ctx context.Context, runners []models.Runner) []models.Runner
}

type flusher struct {
	repo      repo.Repo
	chunkSize int
}

// New returns new instance
func New(repo repo.Repo, chunkSize int) Flusher {
	return &flusher{
		repo:      repo,
		chunkSize: chunkSize,
	}
}

// Flush saves all runners to the database
func (f *flusher) Flush(ctx context.Context, runners []models.Runner) []models.Runner {
	for from := 0; from < len(runners); from += f.chunkSize {
		to := math.Min(from+f.chunkSize, len(runners))

		chunk := runners[from:to]
		if err := f.repo.AddRunners(ctx, chunk); err != nil {
			return runners[from:]
		}
	}

	return nil
}
