package serializer

import "net/http"

// UserUpdateRequest represents the request for create user.
type LoginRequest struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// Bind only implements interface contract.
func (u *LoginRequest) Bind(r *http.Request) error {
	return nil
}
