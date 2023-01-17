package service

import (
	"context"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/thread/repository"
)

type ThreadService interface {
	CreatePosts(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error)
	GetDetailsThread(ctx context.Context, thread *models.Thread) (*models.Thread, error)
	GetPosts(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error)
	UpdateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error)
	Vote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error)
}

type threadService struct {
	threadRepo repository.ThreadRepository
}

func NewThreadService(r repository.ThreadRepository) ThreadService {
	return &threadService{
		threadRepo: r,
	}
}

func (t threadService) CreatePosts(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error) {
	var err error
	var res []*models.Post

	if thread.Slug != "" {
		res, err = t.threadRepo.CreatePostsBySlug(ctx, thread, posts)
	} else {
		res, err = t.threadRepo.CreatePostsByID(ctx, thread, posts)
	}

	if err != nil {
		return nil, errors.Wrap(err, "CreatePosts")
	}

	return res, nil
}

func (t threadService) GetDetailsThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	var err error
	var res *models.Thread

	if thread.Slug != "" {
		res, err = t.threadRepo.GetDetailsThreadBySlug(ctx, thread)
	} else {
		res, err = t.threadRepo.GetDetailsThreadByID(ctx, thread)
	}

	if err != nil {
		return nil, errors.Wrap(err, "GetDetailsThread")
	}

	return res, nil
}

func (t threadService) GetPosts(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error) {
	var err error
	var res []*models.Post

	if thread.Slug != "" {
		res, err = t.threadRepo.GetPostsBySlug(ctx, thread, params)
	} else {
		res, err = t.threadRepo.GetPostsByID(ctx, thread, params)
	}

	if err != nil {
		return nil, errors.Wrap(err, "GetPosts")
	}

	return res, nil
}

func (t threadService) UpdateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	var err error
	var res *models.Thread

	if thread.Slug != "" {
		res, err = t.threadRepo.UpdateThreadBySlug(ctx, thread)
	} else {
		res, err = t.threadRepo.UpdateThreadByID(ctx, thread)
	}

	if err != nil {
		return nil, errors.Wrap(err, "UpdateThread")
	}

	return res, nil
}

func (t threadService) Vote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error) {
	var err error
	var res *models.Thread

	if thread.Slug != "" {
		res, err = t.threadRepo.VoteBySlug(ctx, thread, params)
	} else {
		res, err = t.threadRepo.VoteByID(ctx, thread, params)
	}

	if err != nil {
		return nil, errors.Wrap(err, "Vote")
	}

	return res, nil
}
