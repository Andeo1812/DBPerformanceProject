package service

import (
	"context"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/post/repository"
)

type PostService interface {
	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	GetDetailsPost(ctx context.Context, post *models.Post, params *pkg.PostDetailsParams) (*models.PostDetails, error)
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(r repository.PostRepository) PostService {
	return &postService{
		postRepo: r,
	}
}

func (p postService) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	exist, _ := p.postRepo.CheckExistPost(ctx, post)
	if !exist {
		return nil, errors.Wrap(pkg.ErrSuchPostNotFound, "UpdatePost")
	}

	if post.Message == "" {
		res, err := p.postRepo.GetDetailsPost(ctx, post, &pkg.PostDetailsParams{})
		if err != nil {
			return nil, errors.Wrap(err, "GetDetailsPost")
		}

		return &res.Post, nil
	}

	res, err := p.postRepo.UpdatePost(ctx, post)
	if err != nil {
		return nil, errors.Wrap(err, "UpdatePost")
	}

	return res, nil
}

func (p postService) GetDetailsPost(ctx context.Context, post *models.Post, params *pkg.PostDetailsParams) (*models.PostDetails, error) {
	res, err := p.postRepo.GetDetailsPost(ctx, post, params)
	if err != nil {
		return nil, errors.Wrap(err, "GetDetailsPost")
	}

	return res, nil
}
