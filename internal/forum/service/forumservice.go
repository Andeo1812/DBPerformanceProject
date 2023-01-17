package service

import (
	"context"
	stdErrors "github.com/pkg/errors"

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
	res, err := f.forumRepo.CreateForum(ctx, forum)
	if err != nil {
		return nil, stdErrors.Wrap(err, "CreateForum")
	}

	return res, nil
}

func (f forumService) GetDetailsForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	res, err := f.forumRepo.GetDetailsForum(ctx, forum)
	if err != nil {
		return nil, stdErrors.Wrap(err, "GetDetailsForum")
	}

	return res, nil
}

func (f forumService) GetThreads(ctx context.Context, forum *models.Forum, params *pkg.GetThreadsParams) ([]*models.Thread, error) {
	res, err := f.forumRepo.GetThreads(ctx, forum, params)
	if err != nil {
		return nil, stdErrors.Wrap(err, "GetThreads")
	}

	return res, nil
}

func (f forumService) GetUsers(ctx context.Context, forum *models.Forum, params *pkg.GetUsersParams) ([]*models.User, error) {
	res, err := f.forumRepo.GetUsers(ctx, forum, params)
	if err != nil {
		return nil, stdErrors.Wrap(err, "GetUsers")
	}

	return res, nil
}

func (f forumService) CreateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	res, err := f.forumRepo.CreateThread(ctx, thread)
	if err != nil {
		return nil, stdErrors.Wrap(err, "CreateThread")
	}

	return res, nil
}
