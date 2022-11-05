package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
)
type UserSchema struct {
	ID        primitive.ObjectID    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	Age       int      `json:"age" bson:"age"`
	Gender    string    `json:"gender" bson:"gender"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`

	Traits []UserTraitSchema `json:"traits" bson:"traits"`
}

func FromNewUser(e entity.User) *UserSchema {
	schema := UserSchema{
		Name:      e.Name,
		Email:     e.Email,
		Password:  e.Password,
		Age:       e.Age,
		Gender:    e.Gender,
		CreatedAt: e.CreatedAt,
		Traits:    attributesEntityToSchemaCollection(e.Traits),
	}

	if e.ID != "" {
		id, _ := primitive.ObjectIDFromHex(e.ID)
		schema.ID = id
	}

	return &schema
}

// todo find a better name
func attributesEntityToSchemaCollection(data []entity.UserTrait) []UserTraitSchema {
	var collection []UserTraitSchema

	for _, attr := range data {
		collection = append(collection, UserTraitSchema{
			ID:    attr.ID,
			Value: attr.Value,
		})
	}

	return collection
}

func (u UserSchema) ToEntity() entity.User {
	return entity.User{
		ID:        u.ID.Hex(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Age:       u.Age,
		Gender:    u.Gender,
		CreatedAt: u.CreatedAt,
	}
}