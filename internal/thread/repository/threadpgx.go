package repository

import (
	"context"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
)

type ThreadRepository interface {
	CreatePosts(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error)
	GetDetailsThread(ctx context.Context, thread *models.Thread) (*models.Thread, error)
	GetPosts(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error)
	UpdateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error)
	Vote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error)
}

type threadPostgres struct {
	database *sqltools.Database
}

func NewThreadPostgres(database *sqltools.Database) ThreadRepository {
	return &threadPostgres{
		database,
	}
}

func (t threadPostgres) CreatePosts(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error) {
	panic("implement me")
}

func (t threadPostgres) GetDetailsThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}

func (t threadPostgres) GetPosts(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error) {
	panic("implement me")
}

func (t threadPostgres) UpdateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}

func (t threadPostgres) Vote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error) {
	panic("implement me")
}
