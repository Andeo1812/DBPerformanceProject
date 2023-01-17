package service

import (
	"context"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/vote/repository"
)

type VoteService interface {
	Vote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error)
}

type voteService struct {
	voteRepo repository.VoteRepository
}

func NewVoteService(r repository.VoteRepository) VoteService {
	return &voteService{
		voteRepo: r,
	}
}

func (t voteService) Vote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (*models.Thread, error) {
	var err error
	var res *models.Thread

	if thread.Slug != "" {
		res, err = t.voteRepo.VoteBySlug(ctx, thread, params)
	} else {
		res, err = t.voteRepo.VoteByID(ctx, thread, params)
	}

	if err != nil {
		return nil, errors.Wrap(err, "Vote")
	}

	return res, nil
}
