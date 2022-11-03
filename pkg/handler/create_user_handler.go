package handler

import (
	"net/http"

	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

type createUserHandler struct {
	service *service.UserService
}

func NewCreateUserHandler(service *service.UserService) *createUserHandler {
	return &createUserHandler{
		service: service,
	}
}

func (h *createUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.CreateRandomUser(r.Context())

	if err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	presenter.RenderResponse(w, r, http.StatusCreated, presenter.FromUser(user))
}
