package handler

import (
	"net/http"

	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"github.com/thearyanahmed/mitte_challenge/pkg/serializer"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

type loginHandler struct {
	authService *service.AuthService
}

func NewLoginHandler(authService *service.AuthService) *loginHandler {
	return &loginHandler{
		authService: authService,
	}
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	formRequest := &serializer.LoginRequest{}

	if formErrors := serializer.ValidateJson(r, formRequest); len(formErrors) > 0 {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrorValidationFailed(formErrors))
		return
	}

	user, err := h.authService.FindUserByEmail(r.Context(), formRequest.Email)
	if err != nil {
		// @improvement | maybe we can check if it's a not found error or not.
		// if yes, render a 404 error.
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))

		return
	}

	if !h.authService.ComparePassword(user.Password, []byte(formRequest.Password)) {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrInvalidCredentials())
		return
	}

	token, err := h.authService.GenerateNewToken(r.Context(), user.ID.Hex())

	if err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	presenter.RenderResponse(w, r, http.StatusCreated, presenter.FromToken(token))
}
