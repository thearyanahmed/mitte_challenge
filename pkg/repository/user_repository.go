package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

const table = "users"

// UserRepository represents the user repository that communicates with the database.
type UserRepository struct {
	db *dynamodb.Client
}

func NewUserRepository(db *dynamodb.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) StoreUser(ctx context.Context, user UserSchema) error {
	attribute, err := attributevalue.MarshalMap(user)

	if err != nil {
		return err
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      attribute,
	})

	return err
}
