package service

import "github.com/thearyanahmed/mitte_challenge/pkg/repository"

type UserService struct {
	repository userRepository
}

type userRepository interface {
	CreateUser() (repository.UserSchema, error)
}

func NewUserService(repo userRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}