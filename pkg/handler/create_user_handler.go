package handler

import (
	"net/http"

	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

type createUserHandler struct {
	// @todo use interface
	service *service.UserService
}

func NewCreateUserHandler(service *service.UserService) *createUserHandler {
	return &createUserHandler{
		service: service,
	}
}

func (h *createUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter.Response(w, http.StatusOK, map[string]string{"message": "/user/create"})
}
