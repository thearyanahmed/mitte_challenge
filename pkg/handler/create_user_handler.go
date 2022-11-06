package handler

import (
	"net/http"

	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

type CreateUserHandler struct {
	service *service.UserService
}

func NewCreateUserHandler(service *service.UserService) *CreateUserHandler {
	return &CreateUserHandler{
		service: service,
	}
}

func (h *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.CreateRandomUser(r.Context())

	if err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	presenter.RenderResponse(w, r, http.StatusCreated, presenter.FromNewUser(user))
}
