package handler

import (
	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"net/http"

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

	if formErrors := serializer.ValidateGetQuery(r, filterRequest); len(formErrors) > 0 {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrorValidationFailed(formErrors))
		return
	}

	filterRequest.PopulateUsingQuery(r)

	collection, err := h.userService.GetProfiles(r.Context(), filterRequest)

	if err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	presenter.RenderResponse(w, r, http.StatusOK, presenter.FromUsers(collection))
}
