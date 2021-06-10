package api

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-runner-api/internal/models"
	server "github.com/ozoncp/ocp-runner-api/pkg/ocp-runner-api"
)

// CRUD runner API
type api struct {
	server.UnimplementedOcpRunnerServiceServer
}

// NewRunnerApi constructor
func NewRunnerApi() server.OcpRunnerServiceServer {
	return &api{}
}

// CreateRunner creates new runner
func (a *api) CreateRunner(_ context.Context, request *server.CreateRunnerRequest) (*server.CreateRunnerResponse, error) {
	guid, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create new GUID")
	}

	runner := models.Runner{
		Guid: guid.String(),
		Os:   request.Os,
		Arch: request.Arch,
	}
	log.Info().Str("action", "Created runner").Str("runner", runner.String()).Send()
	return &server.CreateRunnerResponse{}, nil
}

// DescribeRunner updates runner
func (a *api) DescribeRunner(_ context.Context, request *server.DescribeRunnerRequest) (*server.DescribeRunnerResponse, error) {
	if len(request.Guid) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Received empty GUID")
	}

	log.Info().Str("action", "Updated runner").Str("guid", request.Guid).Send()
	return &server.DescribeRunnerResponse{}, nil
}

// RemoveRunner removes runner
func (a *api) RemoveRunner(_ context.Context, request *server.RemoveRunnerRequest) (*server.RemoveRunnerResponse, error) {
	if len(request.Guid) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Received empty GUID")
	}

	log.Info().Str("action", "Removed runner").Str("guid", request.Guid).Send()
	return &server.RemoveRunnerResponse{}, nil
}

// ListRunners returns list of all existing runners
func (a *api) ListRunners(_ context.Context, request *server.ListFiltersRequest) (*server.RunnersListResponse, error) {
	log.Info().Str("action", "Listed runners").Int("count", 0).Send()
	return &server.RunnersListResponse{Runners: []*server.Runner{}}, nil
}
