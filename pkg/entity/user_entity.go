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
	UpdatedAt time.Time
}

// Merge merge an existing user struct with another one.
func (u *User) Merge(c User) *User {
	if c.Name != "" {
		u.Name = c.Name
	}

	if c.Email != "" {
		u.Email = c.Email
	}

	if c.Password != "" {
		u.Password = c.Password
	}

	if c.Gender != "" {
		u.Gender = c.Gender
	}

	if c.Age != 0 {
		u.Age = c.Age
	}

	return u
}
