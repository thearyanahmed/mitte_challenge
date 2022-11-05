package service

import (
	"context"

	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
	"github.com/thearyanahmed/mitte_challenge/pkg/repository"
)

type SwipeService struct {
	swipeRepository swipeRepository
}

type swipeRepository interface {
	InsertSwipe(ctx context.Context, schema repository.SwipeSchmea) (repository.SwipeSchmea, error)
	GetSwipesByUserId(ctx context.Context, userId string) ([]repository.SwipeSchmea, error)
	CheckIfSwipeExists(ctx context.Context, userId, profileOwnerId string) (repository.SwipeSchmea, bool, error)
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

// RecordSwipe @todo name variables better
func (s *SwipeService) RecordSwipe(ctx context.Context, userId string, record recordableSwipe) (entity.Swipe, error) {
	recordSchema := repository.SwipeSchmea{
		SwipedBy:       userId,
		Preference:     record.GetPreference(),
		ProfileOwnerID: record.GetProfileId(),
	}

	swipe, err := s.swipeRepository.InsertSwipe(ctx, recordSchema)

	if err != nil {
		return entity.Swipe{}, err
	}

	return swipe.ToEntity(), nil
}

func (s *SwipeService) CheckIfSwipeExists(ctx context.Context, userId, profileOwnerId string) (entity.Swipe, bool, error) {
	swipe, exists, err := s.swipeRepository.CheckIfSwipeExists(ctx, userId, profileOwnerId)

	if err != nil {
		return entity.Swipe{}, false, err
	}

	return swipe.ToEntity(), exists, nil
}