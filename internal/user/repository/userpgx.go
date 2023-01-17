package repository

import (
	"context"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg/sqltools"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetProfile(ctx context.Context, user *models.User) (*models.User, error)
	UpdateProfile(ctx context.Context, user *models.User) (*models.User, error)
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

func (u userPostgres) GetProfile(ctx context.Context, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (u userPostgres) UpdateProfile(ctx context.Context, user *models.User) (*models.User, error) {
	panic("implement me")
}
