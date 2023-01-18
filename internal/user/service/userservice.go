package service

import (
	"context"
	"db-performance-project/internal/pkg"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/user/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) ([]*models.User, error)
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

func (u userService) CreateUser(ctx context.Context, user *models.User) ([]*models.User, error) {
	var err error
	res := make([]*models.User, 1)

	res[0], err = u.userRepo.CreateUser(ctx, user)
	if errors.Is(errors.Cause(err), pkg.ErrSuchUserExist) {
		res, _ = u.userRepo.GetUserByEmailOrNickname(ctx, user)
		// if err != nil {
		//	return nil, err
		// }

		return res, nil
	}

	// else {
	//	return nil, errors.Wrap(err, "CreateUser")
	// }

	return res, nil
}

func (u userService) GetProfile(ctx context.Context, user *models.User) (*models.User, error) {
	res, err := u.userRepo.GetUserByNickname(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "GetUserByEmailOrNickname")
	}

	return res, nil
}

func (u userService) UpdateProfile(ctx context.Context, user *models.User) (*models.User, error) {
	res, err := u.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "UpdateUser")
	}

	return res, nil
}
