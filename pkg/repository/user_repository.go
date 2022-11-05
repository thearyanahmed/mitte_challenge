package repository

import (
	"context"
	"errors"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

const users_table = "users"

// UserRepository represents the user repository that communicates with the database.
type UserRepository struct {
	db *dynamodb.Client
}

func NewUserRepository(db *dynamodb.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) StoreUser(ctx context.Context, user schema.UserSchema) error {
	attribute, err := attributevalue.MarshalMap(user)

	if err != nil {
		return err
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(users_table),
		Item:      attribute,
	})

	return err
}

func (r *UserRepository) FindUserById(ctx context.Context, id string) (schema.UserSchema, error) {
	return r.findUserBy(ctx, "id", id)
}

// FindUserByEmail @todo change returns, manage errors better
func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (schema.UserSchema, error) {
	filt := expression.Name("email").Equal(expression.Value(email))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		// @todo these should to send raw error
		return schema.UserSchema{}, err
	}

	result, err := r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(users_table),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	if err != nil {
		return schema.UserSchema{}, err
	}

	if result.Count < 1 {
		return schema.UserSchema{}, errors.New("no records found")
	}

	user := schema.UserSchema{}

	var marshalErr error
	for _, v := range result.Items {
		marshalErr = attributevalue.UnmarshalMap(v, &user)
		break
	}

	return user, marshalErr
}

func (r *UserRepository) findUserBy(ctx context.Context, key, value string) (schema.UserSchema, error) {
	result, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(users_table),
		Key: map[string]types.AttributeValue{
			key: &types.AttributeValueMemberS{Value: value},
		},
	})

	if err != nil {
		return schema.UserSchema{}, err
	}

	user := schema.UserSchema{}

	err = attributevalue.UnmarshalMap(result.Item, &user)

	if err != nil {
		return schema.UserSchema{}, err
	}

	return user, nil
}

func (r *UserRepository) FindUsers(ctx context.Context, filters map[string]string) ([]schema.UserSchema, error) {
	builder := expression.NewBuilder()

	for k, v := range filters {
		builder = builder.WithFilter(expression.Name(k).Equal(expression.Value(v)))
	}

	expr, err := builder.Build()

	if err != nil {
		// @todo these should to send raw error
		return []schema.UserSchema{}, err
	}

	result, err := r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(users_table),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	if err != nil {
		return []schema.UserSchema{}, err
	}

	var collection []schema.UserSchema

	err = attributevalue.UnmarshalListOfMaps(result.Items, &collection)

	return collection, err
}
