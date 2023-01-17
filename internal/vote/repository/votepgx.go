package repository

import (
	"context"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
)

type VoteRepository interface {
	VoteByID(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error)
	VoteBySlug(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error)
}

type votePostgres struct {
	database *sqltools.Database
}

func NewVotePostgres(database *sqltools.Database) VoteRepository {
	return &votePostgres{
		database,
	}
}

func (t votePostgres) VoteByID(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error) {
	panic("implement me")
}

func (t votePostgres) VoteBySlug(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error) {
	panic("implement me")
}
