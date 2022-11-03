package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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
	// email = "clarissajewess@goldner.org"
	filt := expression.Name("email").Equal(expression.Value(email))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	result, err := r.db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(table),
		// Limit:            aws.Int32(5),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	fmt.Println("all records", result.Items, "COUNT ->", result.Count)

	if err != nil {
		fmt.Println("NOT FOUND ERR != NIL", err)
		return UserSchema{}, err
	}

	// if result.Count < 1 {
	// 	return UserSchema{}, errors.New("no records found")
	// }

	user := UserSchema{}

	var marshalErr error
	for _, v := range result.Items {
		marshalErr = attributevalue.UnmarshalMap(v, &user)
		break
	}

	fmt.Println("USER", user, marshalErr)

	return user, marshalErr
}

func (r *UserRepository) findUserBy(ctx context.Context, key, value string) (UserSchema, error) {
	result, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			key: &types.AttributeValueMemberS{Value: value},
		},
	})

	if err != nil {
		return UserSchema{}, err
	}

	user := UserSchema{}

	err = attributevalue.UnmarshalMap(result.Item, &user)

	if err != nil {
		return UserSchema{}, err
	}

	return user, nil
}
