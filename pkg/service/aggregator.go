package service

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/thearyanahmed/mitte_challenge/pkg/repository"
)

type ServiceAggregator struct {
	*UserService
	*AuthService
}

func NewServiceAggregator(db *dynamodb.Client) *ServiceAggregator {
	userRepo := repository.NewUserRepository(db)
	userSvc := NewUserService(userRepo)

	authSvc := NewAuthService(userRepo)

	return &ServiceAggregator{
		UserService: userSvc,
		AuthService: authSvc,
	}
}
