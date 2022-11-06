package handler

import (
	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"github.com/thearyanahmed/mitte_challenge/pkg/serializer"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
	"net/http"
)

type swipeHandler struct {
	swipeService *service.SwipeService
}

type swipeResponse struct {
	Message         string `json:"message"`
	Matched         bool   `json:"preferenceMatched"`
	RecordedSwipeId string `json:"recorded_swipe_id"`
	MatchedSwipeId  string `json:"matched_swipe_id"`
}

func NewSwipeHandler(swipeSvc *service.SwipeService) *swipeHandler {
	return &swipeHandler{
		swipeService: swipeSvc,
	}
}

// ServeHttp
// After validating the request, we should check if
func (h *swipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	swipeRequest := &serializer.SwipeRequest{}

	if formErrors := serializer.ValidatePostForm(r, swipeRequest); len(formErrors) > 0 {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrorValidationFailed(formErrors))
		return
	}

	// @todo check if new user id and auth id is same or not
	authUserId := service.GetAuthUserId(r)

	swipe, swiped, err := h.swipeService.CheckIfSwipeExists(r.Context(), authUserId, swipeRequest.ProfileOwnerID)

	if err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	status := http.StatusOK

	if !swiped {
		swipe, err = h.swipeService.RecordSwipe(r.Context(), service.GetAuthUserId(r), swipeRequest)
		// could not create any record, display error
		if err != nil {
			_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
			return
		}

		status = http.StatusCreated
	}

	// Check if it has a view where profile_owner_id -> myself, user_id -> profile_owner_id
	authUserSwipe, authUserSwiped, err := h.swipeService.CheckIfSwipeExists(r.Context(), swipeRequest.ProfileOwnerID, authUserId)
	if err != nil {
		_ = presenter.RenderErrorResponse(w, r, presenter.ErrBadRequest(err))
		return
	}

	matched := swipe.MatchedWith(authUserSwipe)

	response := swipeResponse{
		Message:         "swipe recorded.",
		Matched:         matched,
		RecordedSwipeId: swipe.ID,
	}

	if swiped && authUserSwiped {
		// check preference
		response.Matched = matched

		if matched {
			response.MatchedSwipeId = authUserSwipe.ID
		}
	}

	presenter.RenderResponse(w, r, status, response)
}
