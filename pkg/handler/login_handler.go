package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"github.com/thearyanahmed/mitte_challenge/pkg/serializer"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

type loginHanlder struct {
	authService *service.AuthService
}

func NewLoginHandler(authService *service.AuthService) *loginHanlder {
	return &loginHanlder{
		authService: authService,
	}
}

func (h *loginHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loginReq := &serializer.LoginRequest{}
	if err := render.Bind(r, loginReq); err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	// validate request
	if err := serializer.NewValidator().Struct(loginReq); err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrorValidationFailed(err))
		return
	}

	user, err := h.authService.FindUserByEmail(r.Context(), loginReq.Email)
	if err != nil {
		_ = presenter.ErrBadRequest(err)
		return
	}

	if !h.authService.ComparePassword(user.Password, []byte(loginReq.Password)) {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrInvalidCredentials())
		return
	}

	token, err := h.authService.GenerateNewToken(r.Context(), user.ID)

	if err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	presenter.RenderResponse(w, r, http.StatusCreated, presenter.FromToken(token))
}
