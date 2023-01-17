package service

import (
	"context"

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
	panic("implement me")
}

func (t threadService) GetDetailsThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}

func (t threadService) GetPosts(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error) {
	panic("implement me")
}

func (t threadService) UpdateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}

func (t threadService) Vote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error) {
	panic("implement me")
}
