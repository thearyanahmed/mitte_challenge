package service

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
	"github.com/thearyanahmed/mitte_challenge/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

const defaultPasswordLenght = 10

type UserService struct {
	repository userRepository
}

type userRepository interface {
	StoreUser(context.Context, repository.UserSchema) error
}

func NewUserService(repo userRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (u *UserService) CreateRandomUser(ctx context.Context) (entity.User, error) {
	// create random user
	randomStr, err := createRandomString(defaultPasswordLenght)
	if err != nil {
		return entity.User{}, err
	}

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

	fmt.Println("Service user: ", usr)
	createdUser, err := u.repository.StoreUser(ctx, repository.FromUser(usr))

	if err != nil {
		return entity.User{}, err
	}

	x := createdUser.ToEntity()
	x.Password = randomStr // @note: so we can show the user's password, according to the requirement. Ideally, we will not send password with the response

	return x, nil
}

func createRandomString(n int8) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", b), nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
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
