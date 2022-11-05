package entity

import "time"

// User represents the internal user structure.
type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Gender    string
	Age       int
	CreatedAt time.Time

	Traits []UserTrait
}


type Trait struct {
	ID string
	Name string
}

type UserTrait struct {
	ID string
	Value int8
}