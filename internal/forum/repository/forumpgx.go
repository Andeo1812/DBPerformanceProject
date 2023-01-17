package repository

import (
	"context"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
)

type ForumRepository interface {
	CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	GetDetailsForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	GetThreads(ctx context.Context, forum *models.Forum, params *pkg.GetThreadsParams) ([]*models.Thread, error)
	GetUsers(ctx context.Context, forum *models.Forum, params *pkg.GetUsersParams) ([]*models.User, error)
	CreateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error)
}

type forumPostgres struct {
	database *sqltools.Database
}

func NewForumPostgres(database *sqltools.Database) ForumRepository {
	return &forumPostgres{
		database,
	}
}

func (f forumPostgres) CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	panic("implement me")
}

func (f forumPostgres) GetDetailsForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	panic("implement me")
}

func (f forumPostgres) GetThreads(ctx context.Context, forum *models.Forum, params *pkg.GetThreadsParams) ([]*models.Thread, error) {
	panic("implement me")
}

func (f forumPostgres) GetUsers(ctx context.Context, forum *models.Forum, params *pkg.GetUsersParams) ([]*models.User, error) {
	panic("implement me")
}

func (f forumPostgres) CreateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}
