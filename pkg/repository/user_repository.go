package repository

import (
	"context"
	"errors"
	"time"

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

func (r *UserRepository) Insert(ctx context.Context, user *schema.UserSchema) (string, error) {
	user.ID = newObjectId()
	user.CreatedAt = time.Now()

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

func (r *UserRepository) FindById(ctx context.Context, hex string) (schema.UserSchema, error) {
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

// FindByEmail @todo change returns, manage errors better
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (schema.UserSchema, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var user schema.UserSchema

	if err := r.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return schema.UserSchema{}, nil
	}

	return user, nil
}

func (r *UserRepository) Find(ctx context.Context, pipeline mongo.Pipeline) ([]schema.UserSchema, error) {
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return []schema.UserSchema{}, err
	}

	var results []schema.UserSchema
	if err = cursor.All(ctx, &results); err != nil {
		return []schema.UserSchema{}, err
	}

	return results, nil
}
