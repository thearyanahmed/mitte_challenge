package handler

import (
	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"github.com/thearyanahmed/mitte_challenge/pkg/serializer"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
	"net/http"
)

type ProfileHandler struct {
	userService *service.UserService
}

func NewProfileHandler(userSvc *service.UserService) *ProfileHandler {
	return &ProfileHandler{
		userService: userSvc,
	}
}

func (h *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filterRequest := &serializer.ProfileFilterRequest{}

	if formErrors := serializer.ValidateGetQuery(r, filterRequest); len(formErrors) > 0 {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrorValidationFailed(formErrors))
		return
	}

	filterRequest.PopulateUsingQuery(r)

	collection, err := h.userService.GetProfilesFor(r.Context(), filterRequest, service.GetAuthUserId(r))

	if err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	presenter.RenderResponse(w, r, http.StatusOK, presenter.FromUsers(collection))
}
