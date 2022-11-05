package service

import (
	"github.com/thearyanahmed/mitte_challenge/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Aggregator struct {
	*UserService
	*AuthService
	*SwipeService
}

func NewServiceAggregator(db *mongo.Database) *Aggregator {
	userRepo := repository.NewUserRepository(db.Collection(repository.UsersCollection))
	userSvc := NewUserService(userRepo)

	tokenRepository := repository.NewTokenRepository(db.Collection(repository.TokensCollection))
	authSvc := NewAuthService(userRepo, tokenRepository)

	swipeRepo := repository.NewSwipeRepository(db.Collection(repository.SwipesCollection))
	swipeSvc := NewSwipeService(swipeRepo)

	return &Aggregator{
		UserService:  userSvc,
		AuthService:  authSvc,
		SwipeService: swipeSvc,
	}
}
