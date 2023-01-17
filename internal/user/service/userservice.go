package service

import (
	"context"

	"db-performance-project/internal/models"
	"db-performance-project/internal/user/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetProfile(ctx context.Context, user *models.User) (*models.User, error)
	UpdateProfile(ctx context.Context, user *models.User) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		userRepo: r,
	}
}

func (u userService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (u userService) GetProfile(ctx context.Context, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (u userService) UpdateProfile(ctx context.Context, user *models.User) (*models.User, error) {
	panic("implement me")
}
