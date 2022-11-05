package entity

import "time"

// User represents the internal user structure.
type Swipe struct {
	ID             string
	SwipedBy       string
	ProfileOwnerID string
	Preference     string
	CreatedAt      time.Time
}
