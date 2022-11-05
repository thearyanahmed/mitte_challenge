package entity

import "time"

// User represents the internal user structure.
type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Gender    string
	Age       int8
	CreatedAt time.Time
}

