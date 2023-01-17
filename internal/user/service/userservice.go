package service

import (
	"context"

	"github.com/pkg/errors"

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
	res, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "CreateUser")
	}

	return res, nil
}

func (u userService) GetProfile(ctx context.Context, user *models.User) (*models.User, error) {
	res, err := u.userRepo.GetProfile(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "GetProfile")
	}

	return res, nil
}

func (u userService) UpdateProfile(ctx context.Context, user *models.User) (*models.User, error) {
	res, err := u.userRepo.UpdateProfile(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "UpdateProfile")
	}

	return res, nil
}
