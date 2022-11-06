package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/thearyanahmed/mitte_challenge/pkg/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollection = "users"

// UserRepository represents the user repository that communicates with the database.
type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: db,
	}
}

func (r *UserRepository) StoreUser(ctx context.Context, user *schema.UserSchema) (string, error) {
	user.ID = newObjectId()

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return "", errors.New("could not convert inserted id to primitive id")
	}

	return oid.Hex(), nil
}

func (r *UserRepository) FindUserById(ctx context.Context, hex string) (schema.UserSchema, error) {
	objectId, err := primitive.ObjectIDFromHex(hex)

	if err != nil {
		return schema.UserSchema{}, err
	}

	var user schema.UserSchema

	if err := r.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user); err != nil {
		return schema.UserSchema{}, nil
	}

	return user, nil
}

// FindUserByEmail @todo change returns, manage errors better
func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (schema.UserSchema, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var user schema.UserSchema

	if err := r.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return schema.UserSchema{}, nil
	}

	return user, nil
}

func (r *UserRepository) FindUsers(ctx context.Context, requestFilters map[string]string) ([]schema.UserSchema, error) {
	filters := mapPropertyFilter(requestFilters)

	fmt.Println("received filters", filters)

	cursor, err := r.collection.Find(ctx, filters)

	if err != nil {
		return []schema.UserSchema{}, err
	}

	var results []schema.UserSchema

	if err = cursor.All(ctx, &results); err != nil {
		return []schema.UserSchema{}, err
	}

	return results, nil
}

func mapPropertyFilter(requestFilters map[string]string) bson.M {
	x := bson.M{}

	for k, v := range requestFilters {
		if numeric, err := strconv.Atoi(v); err == nil {
			x[k] = numeric
		} else {
			x[k] = v
		}
	}

	return x
}
