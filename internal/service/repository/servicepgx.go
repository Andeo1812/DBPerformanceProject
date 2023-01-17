package repository

import (
	"context"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg/sqltools"
)

type ServiceRepository interface {
	Clear(ctx context.Context)
	GetStatus(ctx context.Context) *models.StatusService
}

type servicePostgres struct {
	database *sqltools.Database
}

func NewServicePostgres(database *sqltools.Database) ServiceRepository {
	return &servicePostgres{
		database,
	}
}

func (s servicePostgres) Clear(ctx context.Context) {
	panic("implement me")
}

func (s servicePostgres) GetStatus(ctx context.Context) *models.StatusService {
	panic("implement me")
}
