package serializer

import (
	"net/http"
)

type SwipeRequest struct {
	Preference     string `json:"preference,omitempty" validate:"required"`
	ProfileOwnerID string `json:"profile_owner_id,omitempty" validate:"required"`
}

func (s *SwipeRequest) Bind(r *http.Request) error {
	return nil
}
