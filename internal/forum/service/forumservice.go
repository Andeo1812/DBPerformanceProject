package service

import (
	"context"

	"db-performance-project/internal/forum/repository"
	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
)

type ForumService interface {
	CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	GetDetailsForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	GetThreads(ctx context.Context, forum *models.Forum, params *pkg.GetThreadsParams) ([]*models.Thread, error)
	GetUsers(ctx context.Context, forum *models.Forum, params *pkg.GetUsersParams) ([]*models.User, error)
	CreateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error)
}

type forumService struct {
	forumRepo repository.ForumRepository
}

func NewForumService(r repository.ForumRepository) ForumService {
	return &forumService{
		forumRepo: r,
	}
}

func (f forumService) CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	panic("implement me")
}

func (f forumService) GetDetailsForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	panic("implement me")
}

func (f forumService) GetThreads(ctx context.Context, forum *models.Forum, params *pkg.GetThreadsParams) ([]*models.Thread, error) {
	panic("implement me")
}

func (f forumService) GetUsers(ctx context.Context, forum *models.Forum, params *pkg.GetUsersParams) ([]*models.User, error) {
	panic("implement me")
}

func (f forumService) CreateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}
