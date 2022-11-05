package repository

import (
	"context"
	"fmt"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const TokensCollection = "tokens"

type TokenRepository struct {
	collection *mongo.Collection
}

func NewTokenRepository(db *mongo.Collection) *TokenRepository {
	return &TokenRepository{
		collection: db,
	}
}

func (r *TokenRepository) StoreToken(ctx context.Context, token schema.TokenSchema) error {
	id , err := r.collection.InsertOne(ctx, token)
	fmt.Println(id.InsertedID)
	return err
}

func (r *TokenRepository) FindToken(ctx context.Context, token string) (schema.TokenSchema, error) {
	filter := bson.D{{"token", token}}

	var tokenSchema schema.TokenSchema

	if err := r.collection.FindOne(ctx, filter).Decode(&tokenSchema); err != nil {
		return schema.TokenSchema{}, nil
	}

	return tokenSchema, nil
}
