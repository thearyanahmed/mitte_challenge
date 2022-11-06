package repository

import (
	"errors"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"
)

var traits []schema.TraitSchema

func init() {
	traits = []schema.TraitSchema{
		{ID: "1", Name: "self confidence"},
		{ID: "2", Name: "curiosity"},
		{ID: "3", Name: "creativity"},
		{ID: "4", Name: "honesty"},
		{ID: "5", Name: "friendliness"},
		{ID: "6", Name: "self control"},
		{ID: "7", Name: "trustworthiness"},
		{ID: "8", Name: "reliability"},
		{ID: "9", Name: "patience"},
		{ID: "10", Name: "humility"},
	}
}

type TraitRepository struct{}

func NewTraitRepository() *TraitRepository {
	return &TraitRepository{}
}

func (r *TraitRepository) All() ([]schema.TraitSchema, error) {
	return traits, nil
}

func (r *TraitRepository) FindById(id string) (schema.TraitSchema, error) {
	for _, trait := range traits {
		if trait.ID == id {
			return trait, nil
		}
	}

	return schema.TraitSchema{}, errors.New("trait not found")
}
