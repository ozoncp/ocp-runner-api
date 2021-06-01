package flusher

import (
	"context"
	"github.com/ozoncp/ocp-runner-api/internal/models"
	"github.com/ozoncp/ocp-runner-api/internal/repo"
	"github.com/ozoncp/ocp-runner-api/internal/utils"
)

type Flusher interface {
	Flush(ctx context.Context, runners []*models.Runner) []*models.Runner
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
func (f *flusher) Flush(ctx context.Context, runners []*models.Runner) []*models.Runner {
	bulks := utils.SplitToBulks(runners, f.chunkSize)
	for i, bulk := range bulks {
		if err := f.repo.AddRunners(ctx, bulk); err != nil {
			return runners[i*len(bulk):]
		}
	}

	return nil
}
