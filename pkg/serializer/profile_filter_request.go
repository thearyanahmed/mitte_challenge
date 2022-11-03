package serializer

import (
	"fmt"
	"net/http"
)

// UserUpdateRequest represents the request for create user.
type ProfileFilterRequest struct {
	Age    string `json:"age" validate:"numeric,omitempty"`
	Gender string `json:"gender" validate:"omitempty,oneof=male female"`
}

// Bind only implements interface contract.
func (u *ProfileFilterRequest) Bind(r *http.Request) error {
	fmt.Println("IS bind")

	return nil
}

func (u *ProfileFilterRequest) ToKeyValuePair() map[string]string {
	mapped := map[string]string{}

	if u.Age != "" {
		// @todo parse int
		mapped["age"] = u.Age
	}

	if u.Gender != "" {
		mapped["gender"] = u.Gender
	}

	return mapped
}
