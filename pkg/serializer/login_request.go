package serializer

import (
	"github.com/thedevsaddam/govalidator"
)

type LoginRequest struct {
	Password string `json:"password" schema:"password"`
	Email    string `json:"email" schema:"email"`
}

func (r *LoginRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"password": []string{"required", "max:100"},
		"email":    []string{"required", "email", "max:100"},
	}
}
