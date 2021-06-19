package repo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-runner-api/internal/models"
	server "github.com/ozoncp/ocp-runner-api/pkg/ocp-runner-api"
)

type Repo interface {
	AddRunner(ctx context.Context, runner *models.Runner) error
	AddRunners(ctx context.Context, runner []*models.Runner) error
	UpdateRunner(ctx context.Context, runner *models.Runner) error
	RemoveRunner(ctx context.Context, guid string) error
	DescribeRunner(ctx context.Context, guid string) (*models.Runner, error)
	ListRunners(ctx context.Context, filters *server.ListFiltersRequest) ([]*models.Runner, error)
	Close()
}

type repo struct {
	db *sqlx.DB
}

// New creates new instance
func New(db *sqlx.DB) Repo {
	return &repo{db: db}
}

// AddRunner adds new runner
func (r *repo) AddRunner(ctx context.Context, runner *models.Runner) error {
	query := squirrel.
		Insert("runners").
		Columns("guid", "os", "arch").
		Values(runner.Guid, runner.Os, runner.Arch).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	_, err := query.ExecContext(ctx)

	return err
}

// AddRunners adds bunch of new runners
func (r *repo) AddRunners(ctx context.Context, runners []*models.Runner) error {
	query := squirrel.
		Insert("runners").
		Columns("guid", "os", "arch").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for _, runner := range runners {
		query = query.Values(runner.Guid, runner.Os, runner.Arch)
	}

	_, err := query.ExecContext(ctx)

	return err
}

// UpdateRunner updates runner data
func (r *repo) UpdateRunner(ctx context.Context, runner *models.Runner) error {
	query := squirrel.
		Update("runners").
		Where(squirrel.Eq{"guid": runner.Guid}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	if len(runner.Os) > 0 {
		query = query.Set("os", runner.Os)
	}

	if len(runner.Arch) > 0 {
		query = query.Set("arch", runner.Arch)
	}

	_, err := query.ExecContext(ctx)

	return err
}

// RemoveRunner removes runner by unique guid
func (r *repo) RemoveRunner(ctx context.Context, guid string) error {
	query := squirrel.
		Delete("runners").
		Where(squirrel.Eq{"guid": guid}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	_, err := query.ExecContext(ctx)

	return err
}

// DescribeRunner returns single runner by guid
func (r *repo) DescribeRunner(_ context.Context, guid string) (*models.Runner, error) {
	query := squirrel.
		Select("*").
		From("runners").
		Where(squirrel.Eq{"guid": guid}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	row := query.QueryRow()

	var runner models.Runner
	if err := row.Scan(&runner.Guid, &runner.Os, &runner.Arch); err != nil {
		return nil, err
	}

	return &runner, nil
}

// ListRunners returns list of runners
func (r *repo) ListRunners(ctx context.Context, filters *server.ListFiltersRequest) ([]*models.Runner, error) {
	query := squirrel.
		Select("*").
		From("runners").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	query = appendFilter(query, filters.Guids, "guid")
	query = appendFilter(query, filters.Os, "os")
	query = appendFilter(query, filters.Arch, "arch")

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var runners []*models.Runner
	for rows.Next() {
		var runner models.Runner
		if err := rows.Scan(&runner.Guid, &runner.Os, &runner.Arch); err != nil {
			continue
		}
		runners = append(runners, &runner)
	}

	return runners, nil
}

func appendFilter(query squirrel.SelectBuilder, filter []string, field string) squirrel.SelectBuilder {
	if len(filter) > 0 {
		or := squirrel.Or{}
		for _, value := range filter {
			eq := squirrel.Eq{field: value}
			or = append(or, eq)
		}
		query = query.Where(or)
	}

	return query
}

// Close closes db connection
func (r *repo) Close() {
	if err := r.db.Close(); err != nil {
		log.Error().Err(err).Send()
	}
}
