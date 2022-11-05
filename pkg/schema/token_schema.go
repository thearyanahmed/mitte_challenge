package schema

import (
	"time"

	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
)

type TokenSchema struct {
	ID        string    `json:"id" dynamodbav:"id"`
	UserId    string    `json:"user_id" dynamodbav:"user_id"`
	Token     string    `json:"token" dynamodbav:"token"`
	Revoked   bool      `json:"revoked" dynamodbav:"revoked"` // use 1 or 0
	CreatedAt time.Time `json:"created_at,omitempty" dynamodbav:"created_at"`
}

func FromToken(e entity.Token) TokenSchema {
	return TokenSchema{
		ID:        e.ID,
		UserId:    e.UserId,
		Token:     e.Token,
		Revoked:   e.Revoked,
		CreatedAt: e.CreatedAt,
	}
}

func (e TokenSchema) ToEntity() entity.Token {
	return entity.Token{
		ID:        e.ID,
		UserId:    e.UserId,
		Token:     e.Token,
		Revoked:   e.Revoked,
		CreatedAt: e.CreatedAt,
	}
}
