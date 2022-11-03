package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func (r *UserRepository) FindUserById(ctx context.Context, id string) (UserSchema, error) {
	return r.findUserBy(ctx, "id", id)
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (UserSchema, error) {
	return r.findUserBy(ctx, "email", email)
}

func (r *UserRepository) findUserBy(ctx context.Context, key, value string) (UserSchema, error) {
	result, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			key: &types.AttributeValueMemberS{Value: value},
		},
	})

	if err != nil {
		fmt.Println("TRYING TO FIND RECORD BY", key, value)
		return UserSchema{}, err
	}

	user := UserSchema{}

	err = attributevalue.UnmarshalMap(result.Item, &user)

	if err != nil {
		fmt.Println("ERROR MARSHAL UNMAP")
		return UserSchema{}, err
	}

	return user, nil
}
