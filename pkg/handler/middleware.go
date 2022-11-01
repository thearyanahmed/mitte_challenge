package handler

import (
	"errors"
	"net/http"

	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
)

// CheckContentTypeJSON handles the wrong content-type header requests.
func CheckContentTypeJSON(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			presenter.ErrResponse(w, errors.New("content-type must be application/json"))
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	}
	return http.HandlerFunc(fn)
}
