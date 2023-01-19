package service

import (
	"context"

	"github.com/pkg/errors"

	repoForum "db-performance-project/internal/forum/repository"
	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	repoUser "db-performance-project/internal/user/repository"
)

type ForumService interface {
	CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	GetDetailsForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	GetThreads(ctx context.Context, forum *models.Forum, params *pkg.GetThreadsParams) ([]*models.Thread, error)
	GetUsers(ctx context.Context, forum *models.Forum, params *pkg.GetUsersParams) ([]*models.User, error)
}

type forumService struct {
	forumRepo repoForum.ForumRepository
	userRepo  repoUser.UserRepository
}

func NewForumService(rf repoForum.ForumRepository, ru repoUser.UserRepository) ForumService {
	return &forumService{
		forumRepo: rf,
		userRepo:  ru,
	}
}

func (f forumService) CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	res, err := f.forumRepo.GetDetailsForum(ctx, forum)
	if err == nil {
		return res, errors.Wrap(pkg.ErrSuchForumExist, "CreateForum")
	}

	res, err = f.forumRepo.CreateForum(ctx, forum)
	if err != nil {
		_, err = f.userRepo.GetUserByNickname(ctx, &models.User{Nickname: forum.User})
		if err != nil {
			return nil, errors.Wrap(err, "CreateForum")
		}
	}

	return res, nil
}

func (f forumService) GetDetailsForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	res, err := f.forumRepo.GetDetailsForum(ctx, forum)
	if err != nil {
		return nil, errors.Wrap(err, "GetDetailsForum")
	}

	return res, nil
}

func (f forumService) GetThreads(ctx context.Context, forum *models.Forum, params *pkg.GetThreadsParams) ([]*models.Thread, error) {
	_, err := f.forumRepo.GetDetailsForum(ctx, forum)
	if err != nil {
		return nil, errors.Wrap(err, "CreateForum")
	}

	res, err := f.forumRepo.GetThreads(ctx, forum, params)
	if err != nil {
		return nil, errors.Wrap(err, "GetThreads")
	}

	return res, nil
}

func (f forumService) GetUsers(ctx context.Context, forum *models.Forum, params *pkg.GetUsersParams) ([]*models.User, error) {
	_, err := f.forumRepo.GetDetailsForum(ctx, forum)
	if err != nil {
		return nil, errors.Wrap(err, "CreateForum")
	}

	res, err := f.forumRepo.GetUsers(ctx, forum, params)
	if err != nil {
		return nil, errors.Wrap(err, "GetUsers")
	}

	return res, nil
}
