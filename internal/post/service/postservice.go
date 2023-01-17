package service

import (
	"context"
	"db-performance-project/internal/pkg"

	"db-performance-project/internal/models"
	"db-performance-project/internal/post/repository"
)

type PostService interface {
	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	GetDetailsPost(ctx context.Context, post *models.Post, params *pkg.PostDetailsParams) (*models.PostDetails, error)
}

type postService struct {
	postRepo repository.PostRepository
}

func NewUserService(r repository.PostRepository) PostService {
	return &postService{
		postRepo: r,
	}
}

func (p postService) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	panic("implement me")
}

func (p postService) GetDetailsPost(ctx context.Context, post *models.Post, params *pkg.PostDetailsParams) (*models.PostDetails, error) {
	panic("implement me")
}
