package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

const tokens_table = "tokens"

type TokenRepository struct {
	db *dynamodb.Client
}

func NewTokenRepository(db *dynamodb.Client) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) StoreToken(ctx context.Context, token TokenSchema) error {
	attribute, err := attributevalue.MarshalMap(token)

	if err != nil {
		return err
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tokens_table),
		Item:      attribute,
	})

	return err
}

func (r *TokenRepository) FindToken(ctx context.Context, tokenStr string) (TokenSchema, error) {
	filt := expression.Name("token").Equal(expression.Value(tokenStr))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		// @todo these should to send raw error
		fmt.Println("CHECKPOINT 5")

		return TokenSchema{}, err
	}

	result, err := r.db.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(users_table),
		Limit:                     aws.Int32(1),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	if err != nil {
		fmt.Println("CHECKPOINT 6")

		return TokenSchema{}, err
	}

	if result.Count < 1 {
		fmt.Println("CHECKPOINT 7")

		return TokenSchema{}, errors.New("no records found")
	}

	token := TokenSchema{}

	var marshalErr error
	for _, v := range result.Items {
		marshalErr = attributevalue.UnmarshalMap(v, &token)
		break
	}
	fmt.Println("CHECKPOINT 8", marshalErr)

	return token, marshalErr
}
