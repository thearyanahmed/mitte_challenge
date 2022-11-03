package presenter

import "github.com/thearyanahmed/mitte_challenge/pkg/entity"

type NewUserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int8   `json:"age"`
	Gender   string `json:"gender"`
}

type UserResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int8   `json:"age"`
	Gender string `json:"gender"`
	// add interestes ?
}

func FromNewUser(u entity.User) NewUserResponse {
	return NewUserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Age:      u.Age,
		Gender:   u.Gender,
	}
}

func FromUser(u entity.User) UserResponse {
	return UserResponse{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		Age:    u.Age,
		Gender: u.Gender,
	}
}

func FromUsers(users []entity.User) []UserResponse {
	result := []UserResponse{}
	for _, u := range users {
		result = append(result, FromUser(u))
	}
	return result
}
