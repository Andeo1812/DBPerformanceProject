package repository

import (
	"context"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
)

type ThreadRepository interface {
	CreatePostsByID(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error)
	GetDetailsThreadByID(ctx context.Context, thread *models.Thread) (*models.Thread, error)
	GetPostsByID(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error)
	UpdateThreadByID(ctx context.Context, thread *models.Thread) (*models.Thread, error)

	CreatePostsBySlug(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error)
	GetDetailsThreadBySlug(ctx context.Context, thread *models.Thread) (*models.Thread, error)
	GetPostsBySlug(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error)
	UpdateThreadBySlug(ctx context.Context, thread *models.Thread) (*models.Thread, error)
}

type threadPostgres struct {
	database *sqltools.Database
}

func NewThreadPostgres(database *sqltools.Database) ThreadRepository {
	return &threadPostgres{
		database,
	}
}

func (t threadPostgres) CreatePostsByID(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error) {
	panic("implement me")
}

func (t threadPostgres) GetDetailsThreadByID(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}

func (t threadPostgres) GetPostsByID(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error) {
	panic("implement me")
}

func (t threadPostgres) UpdateThreadByID(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}

func (t threadPostgres) CreatePostsBySlug(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error) {
	panic("implement me")
}

func (t threadPostgres) GetDetailsThreadBySlug(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}

func (t threadPostgres) GetPostsBySlug(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error) {
	panic("implement me")
}

func (t threadPostgres) UpdateThreadBySlug(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}
