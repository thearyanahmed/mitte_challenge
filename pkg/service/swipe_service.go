package service

import (
	"context"

	"github.com/thearyanahmed/mitte_challenge/pkg/repository"
)

type SwipeService struct {
	swipeRepository swipeRepository
}

type swipeRepository interface {
	InsertSwipe(ctx context.Context, schema repository.SwipeSchmea) (repository.SwipeSchmea, error)
	GetSwipesByUserId(ctx context.Context, userId string) ([]repository.SwipeSchmea, error)
}

type recordableSwipe interface {
	GetPreference() string
	GetProfileId() string
}

func NewSwipeService(repo swipeRepository) *SwipeService {
	return &SwipeService{
		swipeRepository: repo,
	}
}

// @todo name variables better
func (s *SwipeService) RecordSwipe(ctx context.Context, userId string, record recordableSwipe) (repository.SwipeSchmea, error) {
	recordSchema := repository.SwipeSchmea{
		SwipedBy:       userId,
		Preference:     record.GetPreference(),
		ProfileOwnerID: record.GetProfileId(),
	}

	return s.swipeRepository.InsertSwipe(ctx, recordSchema)
}
