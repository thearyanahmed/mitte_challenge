package schema

import (
	"time"

	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
)

type SwipeSchema struct {
	ID       string `json:"id" dynamodbav:"id"`
	SwipedBy string `json:"swiped_by" dynamodbav:"swiped_by"`

	// the other user's id
	ProfileOwnerID string    `json:"profile_owner_id" dynamodbav:"profile_owner_id"`
	Preference     string    `json:"preference" dynamodbav:"preference"`
	CreatedAt      time.Time `json:"created_at,omitempty" dynamodbav:"created_at"`
}

func (s SwipeSchema) ToEntity() entity.Swipe {
	return entity.Swipe{
		ID:             s.ID,
		SwipedBy:       s.SwipedBy,
		ProfileOwnerID: s.ProfileOwnerID,
		Preference:     s.Preference,
		CreatedAt:      s.CreatedAt,
	}
}
