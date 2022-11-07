package service

import (
	"github.com/thearyanahmed/mitte_challenge/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Aggregator struct {
	*UserService
	*AuthService
	*SwipeService

	db *mongo.Database
}

func NewServiceAggregator(db *mongo.Database) *Aggregator {
	userRepo := repository.NewUserRepository(db.Collection(repository.UsersCollection))
	traitRepo := repository.NewTraitRepository() // don't need db, faking it

	traitSvc := NewTraitService(traitRepo)

	userSvc := NewUserService(userRepo, traitSvc)

	tokenRepository := repository.NewTokenRepository(db.Collection(repository.TokensCollection))
	authSvc := NewAuthService(userRepo, tokenRepository)

	swipeRepo := repository.NewSwipeRepository(db.Collection(repository.SwipesCollection))
	swipeSvc := NewSwipeService(swipeRepo)

	return &Aggregator{
		UserService:  userSvc,
		AuthService:  authSvc,
		SwipeService: swipeSvc,
		db:           db,
	}
}

func (a *Aggregator) GetDB() *mongo.Database {
	return a.db
}
