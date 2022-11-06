package repository

import (
	"context"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
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

func (r *TokenRepository) Insert(ctx context.Context, token schema.TokenSchema) error {
	token.ID = newObjectId()
	token.CreatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, token)

	return err
}

func (r *TokenRepository) FindByToken(ctx context.Context, token string) (schema.TokenSchema, error) {
	filter := bson.D{{"token", token}}

	var tokenSchema schema.TokenSchema

	if err := r.collection.FindOne(ctx, filter).Decode(&tokenSchema); err != nil {
		return schema.TokenSchema{}, err
	}

	return tokenSchema, nil
}
