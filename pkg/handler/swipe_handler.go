package handler

import (
	"fmt"
	"net/http"

	"github.com/thearyanahmed/mitte_challenge/pkg/serializer"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

type swipeHandler struct {
	swipeService *service.SwipeService
}

func NewSwipeHanlder(swipeSvc *service.SwipeService) *swipeHandler {
	return &swipeHandler{
		swipeService: swipeSvc,
	}
}

func (h *swipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// @todo validate request
	swipeRequest := &serializer.SwipeRequest{
		ProfileOwnerID: "48e27503-04eb-4a3e-bf09-2af8560679f2",
		Preference:     "yes",
	}

	// Store if it doesn't exist

	h.swipeService.RecordSwipe(r.Context(), service.GetAuthUserId(r), swipeRequest)
	// Check if it has a view where profile_owner_id -> myself, user_id -> profile_owner_id
	// return response based on that.
	fmt.Println("->", swipeRequest)
}
