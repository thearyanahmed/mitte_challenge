package repository

import (
	"time"

	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
)

type UserSchema struct {
	ID        string    `json:"id" dynamodbav:"id"`
	Name      string    `json:"name" dynamodbav:"name"`
	Email     string    `json:"email" dynamodbav:"email"`
	Password  string    `json:"password" dynamodbav:"password"`
	Age       int8      `json:"age" dynamodbav:"age"`
	Gender    string    `json:"gender" dynamodbav:"gender"`
	CreatedAt time.Time `json:"created_at,omitempty" dynamodbav:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" dynamodbav:"updated_at,omitempty"`
}

func FromUser(e entity.User) UserSchema {
	return UserSchema{
		ID:        e.ID,
		Name:      e.Name,
		Email:     e.Email,
		Password:  e.Password,
		Age:       e.Age,
		Gender:    e.Gender,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (u UserSchema) ToEntity() entity.User {
	return entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Age:       u.Age,
		Gender:    u.Gender,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
