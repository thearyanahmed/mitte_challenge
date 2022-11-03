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

	// get 

	fmt.Println("->", swipeRequest)
}
