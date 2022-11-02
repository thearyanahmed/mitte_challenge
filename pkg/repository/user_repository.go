package repository

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const collection = "users"

// UserRepository represents the user repository that communicates with the database.
type UserRepository struct {
	db *dynamodb.Client
}

func NewUserRepository(db *dynamodb.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser() (UserSchema, error) {
	return UserSchema{}, nil
}
