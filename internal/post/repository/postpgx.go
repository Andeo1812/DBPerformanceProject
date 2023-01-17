package repository

import (
	"context"
	"db-performance-project/internal/pkg"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg/sqltools"
)

type PostRepository interface {
	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	GetDetailsPost(ctx context.Context, post *models.Post, params *pkg.PostDetailsParams) (*models.PostDetails, error)
}

type postPostgres struct {
	database *sqltools.Database
}

func NewPostPostgres(database *sqltools.Database) PostRepository {
	return &postPostgres{
		database,
	}
}

func (p postPostgres) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	panic("implement me")
}

func (p postPostgres) GetDetailsPost(ctx context.Context, post *models.Post, params *pkg.PostDetailsParams) (*models.PostDetails, error) {
	panic("implement me")
}
