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

func (r *UserRepository) StoreUser(ctx context.Context, user UserSchema) (UserSchema, error) {
	// attribute, err := attributevalue.MarshalMap(user)

	// if err != nil {
	// 	return UserSchema{}, err
	// }

	output, err := r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item: map[string]types.AttributeValue{
			"id":         &types.AttributeValueMemberS{Value: user.ID},
			"name":       &types.AttributeValueMemberS{Value: user.Name},
			"email":      &types.AttributeValueMemberS{Value: user.Email},
			"password":   &types.AttributeValueMemberS{Value: user.Email},
			"gender":     &types.AttributeValueMemberS{Value: user.Gender},
			"age":        &types.AttributeValueMemberS{Value: fmt.Sprint(user.Age)},
			"created_at": &types.AttributeValueMemberS{Value: user.CreatedAt.String()},
			"updated_at": &types.AttributeValueMemberS{Value: user.UpdatedAt.String()},
		},
	})

	if err != nil {
		return UserSchema{}, err
	}

	var createdUser UserSchema

	err = attributevalue.UnmarshalMap(output.Attributes, &createdUser)

	if err != nil {
		fmt.Println("repository json unmarshal map, ERROR is not nil! ", err)
		return UserSchema{}, err
	}
	fmt.Println("output", output.Attributes, createdUser)

	return createdUser, nil
}
