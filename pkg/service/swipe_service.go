package service

import (
	"context"

	"github.com/thearyanahmed/mitte_challenge/pkg/repository"
)

type SwipeService struct {
	swipeRepository swipeRepository
}

type swipeRepository interface {
	InsertSwipe(ctx context.Context, schema repository.SwipeSchmea) error
	GetSwipesByUserId(ctx context.Context, userId string) ([]repository.SwipeSchmea, error)
}

func NewSwipeService(repo swipeRepository) *SwipeService {
	return &SwipeService{
		swipeRepository: repo,
	}
}
