package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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
