package repository

import (
	"context"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg/sqltools"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmailOrNickname(ctx context.Context, user *models.User) ([]*models.User, error)
	GetUserByNickname(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
}

type userPostgres struct {
	database *sqltools.Database
}

func NewUserPostgres(database *sqltools.Database) UserRepository {
	return &userPostgres{
		database,
	}
}

func (u userPostgres) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (u userPostgres) GetUserByEmailOrNickname(ctx context.Context, user *models.User) ([]*models.User, error) {
	panic("implement me")
}

func (u userPostgres) GetUserByNickname(ctx context.Context, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (u userPostgres) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	panic("implement me")
}
