package api

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-runner-api/internal/models"
	"github.com/ozoncp/ocp-runner-api/internal/repo"
	"github.com/ozoncp/ocp-runner-api/internal/utils"
	server "github.com/ozoncp/ocp-runner-api/pkg/ocp-runner-api"
)

// CRUD runner API
type api struct {
	server.UnimplementedOcpRunnerServiceServer
	repo repo.Repo
}

// NewRunnerApi constructor
func NewRunnerApi(repo repo.Repo) server.OcpRunnerServiceServer {
	return &api{repo: repo}
}

// CreateRunner creates new runner
func (a *api) CreateRunner(ctx context.Context, request *server.CreateRunnerRequest) (*server.CreateRunnerResponse, error) {
	log.Info().Str("action", "create runner").Send()

	guid, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create new guid")
	}

	runner := models.Runner{
		Guid: guid.String(),
		Os:   request.Os,
		Arch: request.Arch,
	}

	if err := a.repo.AddRunner(ctx, &runner); err != nil {
		return nil, status.Error(codes.Internal, "failed to add new runner")
	}

	return &server.CreateRunnerResponse{}, nil
}

// MultiCreateRunner create new runners
func (a *api) MultiCreateRunner(ctx context.Context, request *server.MultiCreateRunnerRequest) (*server.MultiCreateRunnerResponse, error) {
	var runners []*models.Runner
	for _, r := range request.Runners {
		guid, _ := uuid.NewUUID()
		runners = append(runners, &models.Runner{
			Guid: guid.String(),
			Os:   r.Os,
			Arch: r.Arch,
		})
	}

	runnersBulk := utils.SplitToBulks(runners, int(request.BatchSize))
	for _, bulk := range runnersBulk {
		if err := a.repo.AddRunners(ctx, bulk); err != nil {
			return nil, status.Error(codes.Internal, "failed to add new runners")
		}
	}

	return &server.MultiCreateRunnerResponse{}, nil
}

// UpdateRunner updates runner
func (a *api) UpdateRunner(ctx context.Context, request *server.UpdateRunnerRequest) (*server.UpdateRunnerResponse, error) {
	log.Info().Str("action", "update runner").Send()

	if len(request.Guid) == 0 {
		return nil, status.Error(codes.InvalidArgument, "received empty guid")
	}

	runner := models.Runner{
		Guid: request.Guid,
		Os:   request.Os,
		Arch: request.Arch,
	}

	if err := a.repo.UpdateRunner(ctx, &runner); err != nil {
		return nil, status.Error(codes.Internal, "failed to update runner")
	}

	return &server.UpdateRunnerResponse{}, nil
}

// RemoveRunner removes runner
func (a *api) RemoveRunner(ctx context.Context, request *server.RemoveRunnerRequest) (*server.RemoveRunnerResponse, error) {
	log.Info().Str("action", "remove runner").Send()

	if len(request.Guid) == 0 {
		return nil, status.Error(codes.InvalidArgument, "received empty guid")
	}

	if err := a.repo.RemoveRunner(ctx, request.Guid); err != nil {
		return nil, status.Error(codes.Internal, "failed to remove runner")
	}

	return &server.RemoveRunnerResponse{}, nil
}

// ListRunners returns list of all existing runners
func (a *api) ListRunners(ctx context.Context, request *server.ListFiltersRequest) (*server.RunnersListResponse, error) {
	log.Info().Str("action", "list runners").Send()

	runners, err := a.repo.ListRunners(ctx, request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to list runners")
	}

	grpcRunners := make([]*server.Runner, len(runners))
	for i, runner := range runners {
		grpcRunners[i] = &server.Runner{
			Guid: runner.Guid,
			Os:   runner.Os,
			Arch: runner.Arch,
		}
	}

	return &server.RunnersListResponse{Runners: grpcRunners}, nil
}
