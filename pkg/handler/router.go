package handler

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

func SetupRouter(serviceAggregator *service.ServiceAggregator) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/user", func(r chi.Router) {
		r.With(CheckContentTypeJSON).Post("/create", NewCreateUserHandler(serviceAggregator.UserService).ServeHTTP)
	})

	r.Route("/auth", func(r chi.Router) {
		r.With(CheckContentTypeJSON).Post("/login", NewLoginHandler(serviceAggregator.AuthService).ServeHTTP)
	})

	r.Route("/profile", func(r chi.Router) {
		r.With(NewAuthMiddleware(serviceAggregator.AuthService).Handle).Get("/", NewProfileHandler().ServeHTTP)
	})

	return r
}
