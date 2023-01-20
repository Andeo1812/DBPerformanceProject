package service

import (
	"context"

	"github.com/pkg/errors"

	repoForum "db-performance-project/internal/forum/repository"
	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	repoPost "db-performance-project/internal/post/repository"
	repoThread "db-performance-project/internal/thread/repository"
	repoUser "db-performance-project/internal/user/repository"
)

type ThreadService interface {
	CreateThread(ctx context.Context, thread *models.Thread) (models.Thread, error)
	CreatePosts(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]models.Post, error)
	GetDetailsThread(ctx context.Context, thread *models.Thread) (models.Thread, error)
	GetPosts(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]models.Post, error)
	UpdateThread(ctx context.Context, thread *models.Thread) (models.Thread, error)
}

type threadService struct {
	threadRepo repoThread.ThreadRepository
	forumRepo  repoForum.ForumRepository
	userRepo   repoUser.UserRepository
	postRepo   repoPost.PostRepository
}

func NewThreadService(rt repoThread.ThreadRepository, rf repoForum.ForumRepository, ru repoUser.UserRepository, rp repoPost.PostRepository) ThreadService {
	return &threadService{
		threadRepo: rt,
		forumRepo:  rf,
		userRepo:   ru,
		postRepo:   rp,
	}
}

func (t threadService) CreateThread(ctx context.Context, thread *models.Thread) (models.Thread, error) {
	threadExist, err := t.threadRepo.GetThreadIDBySlug(ctx, thread)
	if err == nil {
		var res models.Thread

		res, err = t.threadRepo.GetDetailsThreadByID(ctx, &threadExist)
		if err != nil {
			return res, errors.Wrap(pkg.ErrSuchThreadExist, "CreateForum")
		}

		return models.Thread{}, errors.Wrap(err, "CreateThread")
	}

	res, err := t.threadRepo.CreateThread(ctx, thread)
	if err != nil {
		return models.Thread{}, errors.Wrap(err, "CreateThread")
	}

	return res, err
}

func (t threadService) CreatePosts(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]models.Post, error) {
	var err error

	threadID := models.Thread{Slug: thread.Slug}

	if thread.Slug != "" {
		threadID, err = t.threadRepo.GetThreadIDBySlug(ctx, thread)
		if err != nil {
			return []models.Post{}, errors.Wrap(err, "CreatePosts")
		}
	} else {
		exist, _ := t.threadRepo.CheckExistThread(ctx, &threadID)
		if !exist {
			return []models.Post{}, errors.Wrap(pkg.ErrSuchThreadNotFound, "CreatePosts")
		}
	}

	if posts[0].Parent != 0 {
		var postWithParent *models.Post

		postWithParent, err = t.postRepo.GetParentPost(ctx, posts[0])
		if err != nil {
			return []models.Post{}, errors.Wrap(err, "CreatePosts")
		}

		if postWithParent.Parent != thread.ID {
			return nil, errors.Wrap(pkg.ErrInvalidParent, "CreatePosts")
		}
	}

	res, err := t.threadRepo.CreatePostsByID(ctx, &threadID, posts)
	if err != nil {
		return []models.Post{}, errors.Wrap(err, "CreatePosts")
	}

	return res, nil
}

func (t threadService) GetDetailsThread(ctx context.Context, thread *models.Thread) (models.Thread, error) {
	var err error

	threadID := models.Thread{Slug: thread.Slug}

	if thread.Slug != "" {
		threadID, err = t.threadRepo.GetThreadIDBySlug(ctx, thread)
		if err != nil {
			return models.Thread{}, errors.Wrap(err, "GetDetailsThread")
		}
	}

	res, err := t.threadRepo.GetDetailsThreadByID(ctx, &threadID)
	if err != nil {
		return models.Thread{}, errors.Wrap(err, "GetDetailsThread")
	}

	return res, nil
}

func (t threadService) UpdateThread(ctx context.Context, thread *models.Thread) (models.Thread, error) {
	var err error

	threadID := models.Thread{Slug: thread.Slug}

	if thread.Slug != "" {
		threadID, err = t.threadRepo.GetThreadIDBySlug(ctx, thread)
		if err != nil {
			return models.Thread{}, errors.Wrap(err, "UpdateThread")
		}
	} else {
		exist, _ := t.threadRepo.CheckExistThread(ctx, &threadID)
		if !exist {
			return models.Thread{}, errors.Wrap(pkg.ErrSuchThreadNotFound, "UpdateThread")
		}
	}

	res, err := t.threadRepo.UpdateThreadByID(ctx, &threadID)
	if err != nil {
		return models.Thread{}, errors.Wrap(err, "UpdateThread")
	}

	return res, nil
}

func (t threadService) GetPosts(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]models.Post, error) {
	var res []models.Post
	var err error

	threadID := models.Thread{Slug: thread.Slug}

	if thread.Slug != "" {
		threadID, err = t.threadRepo.GetThreadIDBySlug(ctx, thread)
		if err != nil {
			return nil, errors.Wrap(err, "GetPosts")
		}
	} else {
		exist, _ := t.threadRepo.CheckExistThread(ctx, &threadID)
		if !exist {
			return nil, errors.Wrap(pkg.ErrSuchThreadNotFound, "GetPosts")
		}
	}

	switch params.Sort {
	case pkg.TypeSortFlat:
		res, err = t.threadRepo.GetPostsByIDFlat(ctx, &threadID, params)
	case pkg.TypeSortTree:
		res, err = t.threadRepo.GetPostsByIDTree(ctx, &threadID, params)
	case pkg.TypeSortParentTree:
		res, err = t.threadRepo.GetPostsByIDParentTree(ctx, &threadID, params)
	default:
		return nil, errors.Wrap(pkg.ErrNoSuchRuleSortPosts, "GetPosts")
	}
	if err != nil {
		return nil, errors.Wrap(err, "GetPosts")
	}

	return res, nil
}
