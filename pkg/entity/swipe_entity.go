package entity

import "time"

type Swipe struct {
	ID             string
	SwipedBy       string
	ProfileOwnerID string
	Preference     string
	CreatedAt      time.Time
}

func (s Swipe) MatchedWith(comparable Swipe) bool {
	if s.Preference == comparable.Preference && s.Preference == "yes" {
		return true
	}

	return false
}
