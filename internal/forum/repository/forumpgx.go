package repository

import (
	"context"

	"db-performanc-eproject/internal/models"
	"db-performanc-eproject/internal/pkg/sqltools"
)

type ForumRepository interface {
	CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	GetDetails(ctx context.Context) error
	GetSlugThreads(ctx context.Context) error
	GetSlugUsers(ctx context.Context) error
	CreateSlug(ctx context.Context) error
	GetCollectionNotAuthorized(ctx context.Context) error
	GetSimilarFilms(ctx context.Context) error
}

type forumPostgres struct {
	database *sqltools.Database
}

func NewForumPostgres(database *sqltools.Database) ForumRepository {
	return &forumPostgres{
		database,
	}
}

func (r forumPostgres) CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	panic("implement me")
}

func (r forumPostgres) GetDetails(ctx context.Context) error {
	panic("implement me")
}

func (r forumPostgres) GetSlugThreads(ctx context.Context) error {
	panic("implement me")
}

func (r forumPostgres) GetSlugUsers(ctx context.Context) error {
	panic("implement me")
}

func (r forumPostgres) CreateSlug(ctx context.Context) error {
	panic("implement me")
}

func (r forumPostgres) GetCollectionNotAuthorized(ctx context.Context) error {
	panic("implement me")
}

func (r forumPostgres) GetSimilarFilms(ctx context.Context) error {
	panic("implement me")
}
