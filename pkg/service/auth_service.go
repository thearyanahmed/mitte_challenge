package service

import (
	"context"
	"errors"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"
	"net/http"
	"time"

	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
	"github.com/thearyanahmed/mitte_challenge/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const UserIDKey = "userId"

type AuthService struct {
	userRepository  authRepository
	tokenRepository tokenRepository
}

type authRepository interface {
	FindByEmail(ctx context.Context, email string) (schema.UserSchema, error)
}

type tokenRepository interface {
	Insert(ctx context.Context, token schema.TokenSchema) error
	FindByToken(ctx context.Context, token string) (schema.TokenSchema, error)
}

func NewAuthService(repo authRepository, tokenRepo tokenRepository) *AuthService {
	return &AuthService{
		userRepository:  repo,
		tokenRepository: tokenRepo,
	}
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (string, error) {
	tokenSchema, err := s.tokenRepository.FindByToken(ctx, token)

	if err != nil {
		return "", err
	}

	if tokenSchema.Revoked {
		return "", errors.New("token has expired")
	}

	return tokenSchema.UserId, nil
}

func GetAuthUserId(r *http.Request) string {
	return r.Context().Value(UserIDKey).(string)
}

func (s *AuthService) FindByEmail(ctx context.Context, email string) (schema.UserSchema, error) {
	return s.userRepository.FindByEmail(ctx, email)
}

// ComparePassword @todo hashing can live in it's own service
func (s *AuthService) ComparePassword(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string, so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
func (s *AuthService) GenerateNewToken(ctx context.Context, userId string) (entity.Token, error) {
	// @todo should be configurable
	randomStr, err := utils.CreateRandomString(32)

	if err != nil {
		return entity.Token{}, err
	}

	tok := newToken(userId, randomStr)

	err = s.tokenRepository.Insert(ctx, schema.FromToken(tok))

	if err != nil {
		return entity.Token{}, err
	}

	return tok, nil
}

func newToken(userId, token string) entity.Token {
	return entity.Token{
		//ID:        uuid.New().String(),
		UserId:    userId,
		Token:     token,
		Revoked:   false,
		CreatedAt: time.Now(),
	}
}
