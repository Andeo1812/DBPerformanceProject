package models

import "db-performance-project/internal/models"

//go:generate easyjson -disallow_unknown_fields getstatus.go

type ServiceGetStatusResponse struct {
	User   uint32 `json:"user"`
	Forum  uint32 `json:"forum"`
	Thread uint32 `json:"thread"`
	Post   uint32 `json:"post"`
}

func NewServiceGetStatusResponse(service *models.StatusService) *ServiceGetStatusResponse {
	return &ServiceGetStatusResponse{
		User:   service.User,
		Forum:  service.Forum,
		Thread: service.Thread,
		Post:   service.Post,
	}
}
