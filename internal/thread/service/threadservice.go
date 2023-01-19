package service

import (
	"context"

	"github.com/pkg/errors"

	repoForum "db-performance-project/internal/forum/repository"
	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	repoThread "db-performance-project/internal/thread/repository"
	repoUser "db-performance-project/internal/user/repository"
)

type ThreadService interface {
	CreateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error)
	CreatePosts(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error)
	GetDetailsThread(ctx context.Context, thread *models.Thread) (*models.Thread, error)
	GetPosts(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error)
	UpdateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error)
}

type threadService struct {
	threadRepo repoThread.ThreadRepository
	forumRepo  repoForum.ForumRepository
	userRepo   repoUser.UserRepository
}

func NewThreadService(rt repoThread.ThreadRepository, rf repoForum.ForumRepository, ru repoUser.UserRepository) ThreadService {
	return &threadService{
		threadRepo: rt,
		forumRepo:  rf,
		userRepo:   ru,
	}
}

func (t threadService) CreateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	var err error

	_, err = t.threadRepo.GetThreadIDBySlug(ctx, thread)
	if err == nil {
		return nil, errors.Wrap(pkg.ErrSuchThreadExist, "CreateThread")
	}

	res, err := t.threadRepo.CreateThread(ctx, thread)
	if err != nil {
		_, err = t.userRepo.GetUserByNickname(ctx, &models.User{Nickname: thread.Author})
		if err != nil {
			return nil, errors.Wrap(err, "CreateForum")
		}

		_, err = t.forumRepo.GetDetailsForum(ctx, &models.Forum{User: thread.Author})
		if err != nil {
			return nil, errors.Wrap(err, "CreateForum")
		}

		return nil, errors.Wrap(err, "CreateThread")
	}

	return res, err
}

func (t threadService) CreatePosts(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error) {
	var err error

	threadID := thread

	if thread.Slug != "" {
		threadID, err = t.threadRepo.GetThreadIDBySlug(ctx, thread)
		if err != nil {
			return nil, errors.Wrap(err, "CreatePosts")
		}
	}

	res, err := t.threadRepo.CreatePostsByID(ctx, threadID, posts)
	if err != nil {
		return nil, errors.Wrap(err, "CreatePosts")
	}

	return res, nil
}

func (t threadService) GetDetailsThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	var err error

	threadID := thread

	if thread.Slug != "" {
		threadID, err = t.threadRepo.GetThreadIDBySlug(ctx, thread)
		if err != nil {
			return nil, errors.Wrap(err, "GetDetailsThread")
		}
	}

	res, err := t.threadRepo.GetDetailsThreadByID(ctx, threadID)
	if err != nil {
		return nil, errors.Wrap(err, "GetDetailsThread")
	}

	return res, nil
}

func (t threadService) UpdateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	var err error

	threadID := thread

	if thread.Slug != "" {
		threadID, err = t.threadRepo.GetThreadIDBySlug(ctx, thread)
		if err != nil {
			return nil, errors.Wrap(err, "UpdateThread")
		}
	}

	res, err := t.threadRepo.UpdateThreadByID(ctx, threadID)
	if err != nil {
		return nil, errors.Wrap(err, "UpdateThread")
	}

	return res, nil
}

func (t threadService) GetPosts(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error) {
	var res []*models.Post
	var err error

	threadID := thread

	if thread.Slug != "" {
		threadID, err = t.threadRepo.GetThreadIDBySlug(ctx, thread)
		if err != nil {
			return nil, errors.Wrap(err, "GetPosts")
		}
	}

	switch params.Sort {
	case pkg.TypeSortFlat:
		res, err = t.threadRepo.GetPostsByIDFlat(ctx, threadID, params)
	case pkg.TypeSortTree:
		res, err = t.threadRepo.GetPostsByIDTree(ctx, threadID, params)
	case pkg.TypeSortParentTree:
		res, err = t.threadRepo.GetPostsByIDParentTree(ctx, threadID, params)
	default:
		return nil, errors.Wrap(pkg.ErrNoSuchRuleSortPosts, "GetPosts")
	}
	if err != nil {
		return nil, errors.Wrap(err, "GetPosts")
	}

	return res, nil
}
