package service

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
	"github.com/thearyanahmed/mitte_challenge/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

const defaultPasswordLength = 10

type UserService struct {
	repository UserRepository
}

type RequestFilter interface {
	ToKeyValuePair() map[string]string
}

// UserRepository the bridge between db and service layer
// context is passed from the handler layer, from the request context
type UserRepository interface {
	StoreUser(context.Context, repository.UserSchema) error
	FindUserById(context.Context, string) (repository.UserSchema, error)
	FindUserByEmail(ctx context.Context, email string) (repository.UserSchema, error)
	FindUsers(ctx context.Context, filters map[string]string) ([]repository.UserSchema, error)
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

// CreateRandomUser
func (u *UserService) CreateRandomUser(ctx context.Context) (entity.User, error) {
	randomStr := "secret" // createRandomString(defaultPasswordLenght)

	hashed, err := hashAndSalt([]byte(randomStr))

	if err != nil {
		return entity.User{}, err
	}

	usr := entity.User{
		ID:       uuid.New().String(),
		Name:     gofakeit.Name(),
		Password: hashed,
		Email:    gofakeit.Email(),
		Gender:   gofakeit.Gender(),
		Age:      int8(gofakeit.Number(1, 100)),
	}

	err = u.repository.StoreUser(ctx, repository.FromNewUser(usr))

	if err != nil {
		return entity.User{}, err
	}

	// fetch the user
	createdUser, err := u.repository.FindUserById(ctx, usr.ID)
	if err != nil {
		// user have been created, but can't fetch it, maybe db is unavailable ? or some other issue
		return entity.User{}, err
	}

	// @note: so we can show the user's password, according to the requirement. Ideally, we will not send password with the response
	newUserEntity := createdUser.ToEntity()
	newUserEntity.Password = randomStr

	return newUserEntity, nil
}

func hashAndSalt(pwd []byte) (string, error) {
	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

func (s *UserService) GetProfiles(ctx context.Context, requestFilter RequestFilter) ([]entity.User, error) {

	filters := requestFilter.ToKeyValuePair()

	users, err := s.repository.FindUsers(ctx, filters)

	if err != nil {
		return []entity.User{}, err
	}

	var usersCollection []entity.User

	for _, u := range users {
		usersCollection = append(usersCollection, u.ToEntity())
	}

	return usersCollection, nil
}
