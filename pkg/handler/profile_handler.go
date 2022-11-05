package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"github.com/thearyanahmed/mitte_challenge/pkg/serializer"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

type profileHandler struct {
	userService *service.UserService
}

func NewProfileHandler(userSvc *service.UserService) *profileHandler {
	return &profileHandler{
		userService: userSvc,
	}
}

func (h *profileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filterRequest := &serializer.ProfileFilterRequest{}
	if err := render.Bind(r, filterRequest); err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	// validate request
	if err := serializer.NewValidator().Struct(filterRequest); err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrorValidationFailed(err))
		return
	}

	collection, err := h.userService.GetProfiles(r.Context(), filterRequest)

	if err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	presenter.RenderResponse(w, r, http.StatusOK, presenter.FromUsers(collection))
}
