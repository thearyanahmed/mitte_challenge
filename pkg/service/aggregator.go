package service

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/thearyanahmed/mitte_challenge/pkg/repository"
)

type ServiceAggregator struct {
	*UserService
	*AuthService
	*SwipeService
}

func NewServiceAggregator(db *dynamodb.Client) *ServiceAggregator {
	userRepo := repository.NewUserRepository(db)
	userSvc := NewUserService(userRepo)

	tokenRepository := repository.NewTokenRepository(db)
	authSvc := NewAuthService(userRepo, tokenRepository)

	swipeRepo := repository.NewSwipeRepository(db)
	swipeSvc := NewSwipeService(swipeRepo)

	return &ServiceAggregator{
		UserService:  userSvc,
		AuthService:  authSvc,
		SwipeService: swipeSvc,
	}
}
