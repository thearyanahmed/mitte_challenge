package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
)

type TokenSchema struct {
	ID        primitive.ObjectID    `json:"id" bson:"_id"`
	UserId    string    `json:"user_id" bson:"user_id"`
	Token     string    `json:"token" bson:"token"`
	Revoked   bool      `json:"revoked" bson:"revoked"` // use 1 or 0
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
}

func FromToken(e entity.Token) TokenSchema {
	schema := TokenSchema{
		//ID:        e.ID,
		UserId:    e.UserId,
		Token:     e.Token,
		Revoked:   e.Revoked,
		CreatedAt: e.CreatedAt,
	}


	if e.ID != "" {
		id, _ := primitive.ObjectIDFromHex(e.ID)
		schema.ID = id
	}

	return schema
}

func (e TokenSchema) ToEntity() entity.Token {
	return entity.Token{
		ID:        e.ID.Hex(),
		UserId:    e.UserId,
		Token:     e.Token,
		Revoked:   e.Revoked,
		CreatedAt: e.CreatedAt,
	}
}
