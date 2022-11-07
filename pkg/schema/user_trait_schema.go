package schema

import (
	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
)

type UserTraitSchema struct {
	ID    string `json:"id" bson:"id"`
	Value int8   `json:"value"  bson:"value"`
}

// ToUserTraitSchema todo find a better name
func ToUserTraitSchema(data []entity.UserTrait) []UserTraitSchema {
	var collection []UserTraitSchema

	for _, attr := range data {
		collection = append(collection, UserTraitSchema{
			ID:    attr.ID,
			Value: attr.Value,
		})
	}

	return collection
}
