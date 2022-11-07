package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/thearyanahmed/mitte_challenge/pkg/entity"
)

type SwipeSchema struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	SwipedBy       string             `json:"swiped_by" bson:"swiped_by"`               // authenticated user's id, the one performing the action
	ProfileOwnerID string             `json:"profile_owner_id" bson:"profile_owner_id"` // the other user's id
	Preference     string             `json:"preference" bson:"preference"`
	CreatedAt      time.Time          `json:"created_at,omitempty" bson:"created_at"`
}

func (s SwipeSchema) ToEntity() entity.Swipe {
	return entity.Swipe{
		ID:             s.ID.Hex(),
		SwipedBy:       s.SwipedBy,
		ProfileOwnerID: s.ProfileOwnerID,
		Preference:     s.Preference,
		CreatedAt:      s.CreatedAt,
	}
}
