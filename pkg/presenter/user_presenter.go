package presenter

import "github.com/thearyanahmed/mitte_challenge/pkg/entity"

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int8   `json:"age"`
	Gender   string `json:"gender"`
}

func FromUser(u entity.User) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Age:      u.Age,
		Gender:   u.Gender,
	}
}
