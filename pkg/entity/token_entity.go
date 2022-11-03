package entity

import "time"

type Token struct {
	ID        string    `json:"id" dynamodbav:"id"`
	UserId    string    `json:"user_id" dynamodbav:"user_id"`
	Token     string    `json:"email" dynamodbav:"email"`
	Revoked   bool      `json:"revoked" dynamodbav:"revoked"`
	CreatedAt time.Time `json:"created_at,omitempty" dynamodbav:"created_at"`
}
