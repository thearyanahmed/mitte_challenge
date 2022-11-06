package schema

import (
	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
)

type TraitSchema struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name"  bson:"name"`
}

func (t TraitSchema) ToEntity() entity.Trait {
	return entity.Trait{
		ID:   t.ID,
		Name: t.Name,
	}
}
