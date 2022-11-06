package handler

import (
	"context"
	"net/http"

	"github.com/thearyanahmed/mitte_challenge/pkg/presenter"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

// CheckContentTypeJSON handles the wrong content-type header requests.
func CheckContentTypeJSON(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			_ = presenter.RenderErrorResponse(w, r, presenter.ErrNotAcceptable())
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	}
	return http.HandlerFunc(fn)
}

type authMiddleware struct {
	service *service.AuthService
}

func NewAuthMiddleware(svc *service.AuthService) *authMiddleware {
	return &authMiddleware{
		service: svc,
	}
}

func (m *authMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			_ = presenter.RenderErrorResponse(w, r, presenter.ErrUnauthorized())
			return
		}

		userId, err := m.service.ValidateToken(r.Context(), authToken)

		if err != nil {
			_ = presenter.RenderErrorResponse(w, r, presenter.ErrUnauthorized())
			return
		}

		ctx := context.WithValue(r.Context(), service.UserIDKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
