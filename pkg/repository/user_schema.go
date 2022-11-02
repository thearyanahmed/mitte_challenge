package repository

import (
	"time"

	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
)

type UserSchema struct {
	ID        string    `db:"id" dynamodbav:"id"`
	Name      string    `db:"name" dynamodbav:"name"`
	Email     string    `db:"email" dynamodbav:"email"`
	Password  string    `db:"password" dynamodbav:"password"`
	Age       int8      `db:"age" dynamodbav:"age"`
	Gender    string    `db:"gender" dynamodbav:"gender"`
	CreatedAt time.Time `db:"created_at" dynamodbav:"created_at"`
	UpdatedAt time.Time `db:"updated_at" dynamodbav:"updated_at"`
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
