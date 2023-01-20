package service

import (
	"context"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	threadRepo "db-performance-project/internal/thread/repository"
	voteRepo "db-performance-project/internal/vote/repository"
)

type VoteService interface {
	Vote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (models.Thread, error)
}

type voteService struct {
	voteRepo   voteRepo.VoteRepository
	threadRepo threadRepo.ThreadRepository
}

func NewVoteService(vr voteRepo.VoteRepository, tr threadRepo.ThreadRepository) VoteService {
	return &voteService{
		voteRepo:   vr,
		threadRepo: tr,
	}
}

func (v voteService) Vote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (models.Thread, error) {
	var err error

	threadID := models.Thread{Slug: thread.Slug}

	if thread.Slug != "" {
		threadID, err = v.threadRepo.GetThreadIDByForumAndSlug(ctx, thread)
		if err != nil {
			return models.Thread{}, errors.Wrap(err, "Vote")
		}
	}

	exist, err := v.voteRepo.CheckExistVote(ctx, &threadID, params)
	// if err != nil {
	//	return nil, errors.Wrap(err, "Vote")
	// }

	if exist {
		v.voteRepo.UpdateVote(ctx, &threadID, params)
	} else {
		v.voteRepo.CreateVote(ctx, &threadID, params)
	}
	if err != nil {
		return models.Thread{}, errors.Wrap(err, "Vote")
	}

	threadUPD, _ := v.threadRepo.GetDetailsThreadByID(ctx, &threadID)
	// if err != nil {
	//	return nil, errors.Wrap(err, "Vote")
	// }

	return threadUPD, nil
}
