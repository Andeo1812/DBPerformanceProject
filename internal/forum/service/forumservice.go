package service

import (
	"context"
	"db-performanc-eproject/internal/forum/repository"

	"db-performanc-eproject/internal/models"
)

type ForumService interface {
	CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	GetDetails(ctx context.Context) error
	GetSlugThreads(ctx context.Context) error
	GetSlugUsers(ctx context.Context) error
	CreateSlug(ctx context.Context) error
	GetCollectionNotAuthorized(ctx context.Context) error
	GetSimilarFilms(ctx context.Context) error
}

type forumService struct {
	collectionRepo repository.ForumRepository
}

func NewForumService(r repository.ForumRepository) ForumService {
	return &forumService{
		collectionRepo: r,
	}
}

func (f forumService) CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	panic("implement me")
}

func (f forumService) GetDetails(ctx context.Context) error {
	panic("implement me")
}

func (f forumService) GetSlugThreads(ctx context.Context) error {
	panic("implement me")
}

func (f forumService) GetSlugUsers(ctx context.Context) error {
	panic("implement me")
}

func (f forumService) CreateSlug(ctx context.Context) error {
	panic("implement me")
}

func (f forumService) GetCollectionNotAuthorized(ctx context.Context) error {
	panic("implement me")
}

func (f forumService) GetSimilarFilms(ctx context.Context) error {
	panic("implement me")
}
