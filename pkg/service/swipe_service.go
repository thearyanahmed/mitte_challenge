package service

import (
	"context"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"

	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
)

type SwipeService struct {
	swipeRepository swipeRepository
}

type swipeRepository interface {
	Insert(ctx context.Context, schema schema.SwipeSchema) (schema.SwipeSchema, error)
	GetByUserId(ctx context.Context, userId string) ([]schema.SwipeSchema, error)
	CheckIfSwipeExists(ctx context.Context, userId, profileOwnerId string) (schema.SwipeSchema, bool, error)
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
	recordSchema := schema.SwipeSchema{
		SwipedBy:       userId,
		Preference:     record.GetPreference(),
		ProfileOwnerID: record.GetProfileId(),
	}

	swipe, err := s.swipeRepository.Insert(ctx, recordSchema)

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
