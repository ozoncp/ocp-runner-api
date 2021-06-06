package api

import (
	"context"
	"github.com/ozoncp/ocp-runner-api/internal/models"
	"github.com/rs/zerolog/log"

	server "github.com/ozoncp/ocp-runner-api/pkg/ocp-runner-api"
)

// CRUD runner API
type api struct {
	server.UnimplementedOcpRunnerApiServer
}

// NewRunnerApi constructor
func NewRunnerApi() server.OcpRunnerApiServer {
	return &api{}
}

// CreateRunner creates new runner
func (a *api) CreateRunner(_ context.Context, request *server.Runner) (*server.OperationResult, error) {
	runner := models.Runner{
		Guid: request.Guid,
		Os:   request.Os,
		Arch: request.Arch,
	}
	log.Info().Str("action", "Created runner").Str("runner", runner.String()).Send()
	return &server.OperationResult{Success: true, Guid: runner.Guid}, nil
}

// DescribeRunner updates runner
func (a *api) DescribeRunner(_ context.Context, request *server.Runner) (*server.OperationResult, error) {
	log.Info().Str("action", "Updated runner").Str("guid", request.Guid).Send()
	return &server.OperationResult{Success: true, Guid: request.Guid}, nil
}

// RemoveRunner removes runner
func (a *api) RemoveRunner(_ context.Context, request *server.Runner) (*server.OperationResult, error) {
	log.Info().Str("action", "Removed runner").Str("guid", request.Guid).Send()
	return &server.OperationResult{Success: true, Guid: request.Guid}, nil
}

// ListRunners returns list of all existing runners
func (a *api) ListRunners(_ context.Context, _ *server.ListFilters) (*server.RunnersList, error) {
	log.Info().Str("action", "Listed runners").Int("count", 0).Send()
	return &server.RunnersList{Runners: []*server.Runner{}}, nil
}
