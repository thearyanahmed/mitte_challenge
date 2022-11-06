package service

import (
	"context"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// defaultPasswordLength would be used to create random n Length words
const defaultPasswordLength = 10

type UserService struct {
	userRepository UserRepository
	traitService   traitSvc
}

type RequestFilter interface {
	ToKeyValuePair() map[string]string
}

type traitSvc interface {
	TakeRandom(n int) []entity.Trait
}

// UserRepository the bridge between db and service layer
// context is passed from the handler layer, from the request context
type UserRepository interface {
	Insert(context.Context, *schema.UserSchema) (newlyCreatedId string, err error)
	FindById(context.Context, string) (schema.UserSchema, error)
	FindByEmail(ctx context.Context, email string) (schema.UserSchema, error)
	Find(ctx context.Context, filters map[string]string) ([]schema.UserSchema, error)
	FindMatch()
}

func NewUserService(userRepository UserRepository, traitService traitSvc) *UserService {
	return &UserService{
		userRepository: userRepository,
		traitService:   traitService,
	}
}

func (u *UserService) CreateRandomUser(ctx context.Context) (entity.User, error) {
	randomStr := "secret" // createRandomString(defaultPasswordLength)

	hashed, err := hashAndSalt([]byte(randomStr))

	if err != nil {
		return entity.User{}, err
	}

	randomTraits := toUserTraits(u.traitService.TakeRandom(5))

	usr := entity.User{
		Name:     gofakeit.Name(),
		Password: hashed,
		Email:    gofakeit.Email(),
		Gender:   gofakeit.Gender(),
		Age:      gofakeit.Number(1, 100),
		Traits:   randomTraits,
	}

	newlyCreatedId, err := u.userRepository.Insert(ctx, schema.FromNewUser(usr))

	if err != nil {
		return entity.User{}, err
	}

	usr.ID = newlyCreatedId
	// fetch the user
	createdUser, err := u.userRepository.FindById(ctx, usr.ID)
	if err != nil {
		// user have been created, but can't fetch it, maybe db is unavailable ? or some other issue
		return entity.User{}, err
	}

	// @note: so we can show the user's password, according to the requirement. Ideally, we will not send password with the response
	newUserEntity := createdUser.ToEntity()
	newUserEntity.Password = randomStr

	return newUserEntity, nil
}

func toUserTraits(traits []entity.Trait) []entity.UserTrait {
	var collection []entity.UserTrait

	for _, trait := range traits {
		collection = append(collection, entity.UserTrait{
			ID:    trait.ID,
			Value: int8(gofakeit.Number(1, 100)),
		})
	}

	return collection
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
	// GenerateFromPassword returns a byte slice, so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

func (u *UserService) GetProfiles(ctx context.Context, requestFilter RequestFilter) ([]entity.User, error) {
	filters := requestFilter.ToKeyValuePair()

	// add attractiveness logic
	// get auth user attractiveness
	// where user_id not = auth user id
	// sum
	// sort by sum

	// @note test query
	//u.userRepository.FindMatch()

	users, err := u.userRepository.Find(ctx, filters)
	if err != nil {
		return []entity.User{}, err
	}

	var usersCollection []entity.User
	for _, u := range users {
		usersCollection = append(usersCollection, u.ToEntity())
	}

	return usersCollection, nil
}
