package handler

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(db *dynamodb.Client) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/user", func(r chi.Router) {
		r.With(CheckContentTypeJSON).Post("/create", NewCreateUserHandler(db).ServeHTTP)
	})

	return r
}
