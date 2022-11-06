package service

import (
	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
	"github.com/thearyanahmed/mitte_challenge/pkg/schema"
	"math/rand"
	"time"
)

type TraitService struct {
	repository TraitRepository
}

type TraitRepository interface {
	All() ([]schema.TraitSchema, error)
	FindById(id string) (schema.TraitSchema, error)
}

func NewTraitService(repo TraitRepository) *TraitService {
	return &TraitService{repository: repo}
}

func (s *TraitService) All() ([]schema.TraitSchema, error) {
	return s.repository.All()
}

func (s *TraitService) FindById(id string) (schema.TraitSchema, error) {
	return s.repository.FindById(id)
}

func (s *TraitService) TakeRandom(n int) []entity.Trait {
	traits, err := s.repository.All()

	if err != nil {
		return []entity.Trait{}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(traits), func(i, j int) {
		traits[i], traits[j] = traits[j], traits[i]
	})

	var collection []entity.Trait

	i := 0
	for _, trait := range traits {
		if i >= n {
			break
		}
		collection = append(collection, trait.ToEntity())
		i++
	}

	return collection
}
