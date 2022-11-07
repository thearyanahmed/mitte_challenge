package service

import (
	"context"
	"fmt"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"

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
	ToKeyValuePair() map[string]interface{}
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
	Find(ctx context.Context, pipeline mongo.Pipeline) ([]schema.UserSchema, error)
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

func (u *UserService) GetProfilesFor(ctx context.Context, requestFilter RequestFilter, userId string) ([]entity.User, error) {
	user, err := u.userRepository.FindById(ctx, userId)

	if err != nil {
		return []entity.User{}, nil
	}

	var userTraitIds []string

	for _, v := range user.ToEntity().Traits {
		userTraitIds = append(userTraitIds, v.ID)
	}

	pipeline := getPipeline(requestFilter.ToKeyValuePair(), userTraitIds, userId)
	users, err := u.userRepository.Find(ctx, pipeline)

	if err != nil {
		return []entity.User{}, err
	}

	var usersCollection []entity.User
	for _, u := range users {
		usersCollection = append(usersCollection, u.ToEntity())
	}

	return usersCollection, nil
}

func getPipeline(requestFilters map[string]interface{}, traitIds []string, userId string) mongo.Pipeline {
	if len(traitIds) > 0 {
		requestFilters["$or"] = mapTraitIds(traitIds)
	}

	mappedFilters := mapPropertyFilter(requestFilters)
	mappedFilters = append(mappedFilters, bson.D{{"_id", bson.D{{"$ne", userId}}}})

	match := bson.D{{"$and", mappedFilters}}
	fmt.Print(match)

	return mongo.Pipeline{
		{{"$match", match}},
		{{"$project", getProjection()}},
		{{"$sort", getSort()}},
	}
}

func getProjection() bson.D {
	return bson.D{
		{"name", "$name"},
		{"email", "$email"},
		{"age", "$age"},
		{"gender", "$gender"},
		{"traits", "$traits"},
		{"attractiveness_score", bson.D{
			{"$sum", "$traits.value"},
		}},
	}
}

func getSort() bson.D {
	return bson.D{
		{"attractiveness_score", -1},
	}
}

func mapTraitIds(traitIds []string) bson.A {
	var mapped bson.A

	for _, v := range traitIds {
		mapped = append(mapped, bson.D{{"traits.id", v}})
	}

	return mapped
}

func mapPropertyFilter(requestFilters map[string]interface{}) bson.A {
	var mapped bson.A

	for k, v := range requestFilters {

		var val interface{}
		// if the value is string
		if s, ok := v.(string); ok {
			// check if the value is numeric or not
			// if yes, cast before appending
			if numeric, err := strconv.Atoi(s); err == nil {
				val = numeric
			} else {
				val = v
			}
		} else {
			val = v
		}

		mapped = append(mapped, bson.D{{k, val}})
	}

	return mapped
}
