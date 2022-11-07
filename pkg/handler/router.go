package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

// BootstrapRouter sets up a new router, proper middlewares and sets up route-to-handler mapping
func BootstrapRouter(serviceAggregator *service.Aggregator) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.With(CheckContentTypeJSON)

	r.Route("/user", func(r chi.Router) {
		r.Post("/create", NewCreateUserHandler(serviceAggregator.UserService).ServeHTTP)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", NewLoginHandler(serviceAggregator.AuthService).ServeHTTP)
	})

	r.Route("/profile", func(r chi.Router) {
		r.With(NewAuthMiddleware(serviceAggregator.AuthService).Handle).
			Get("/", NewProfileHandler(serviceAggregator.UserService).ServeHTTP)
	})

	r.Route("/swipe", func(r chi.Router) {
		r.With(NewAuthMiddleware(serviceAggregator.AuthService).Handle).
			Post("/", NewSwipeHandler(serviceAggregator.SwipeService).ServeHTTP)
	})

	return r
}
