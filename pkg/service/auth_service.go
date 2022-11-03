package service

import (
	"context"
	"fmt"

	"github.com/thearyanahmed/mitte_challenge/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository authRepository
}

type authRepository interface {
	FindUserByEmail(context.Context, string) (repository.UserSchema, error)
}

func NewAuthService(repo authRepository) *AuthService {
	return &AuthService{
		repository: repo,
	}
}

func (s *AuthService) Attempt(ctx context.Context, email, password string) bool {
	user, err := s.repository.FindUserByEmail(ctx, email)

	if err != nil {
		fmt.Println("USER EMAIL NOT FOUND", email, err)
		return false
	}

	return comparePasswords(user.Password, []byte(password)) == true
}

// @todo hashing can live in it's own service
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	fmt.Println("hashed", hashedPwd, "plain", string(plainPwd))
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println("FAILED", err)
		return false
	}

	fmt.Println("should return true")

	return true
}
