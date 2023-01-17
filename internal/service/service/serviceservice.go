package service

import (
	"context"
	"db-performance-project/internal/models"
	"db-performance-project/internal/service/repository"
)

type Service interface {
	Clear(ctx context.Context)
	GetStatus(ctx context.Context) *models.StatusService
}

type service struct {
	serviceRepo repository.ServiceRepository
}

func NewService(r repository.ServiceRepository) Service {
	return &service{
		serviceRepo: r,
	}
}

func (s service) Clear(ctx context.Context) {
	s.serviceRepo.Clear(ctx)
}

func (s service) GetStatus(ctx context.Context) *models.StatusService {
	return s.serviceRepo.GetStatus(ctx)
}
