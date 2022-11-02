package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/thearyanahmed/mitte_challenge/pkg/service"
)

func SetupRouter(serviceAggregator *service.ServiceAggregator) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/user", func(r chi.Router) {
		r.With(CheckContentTypeJSON).With(NewAuthMiddleware().Handle).Post("/create", NewCreateUserHandler(serviceAggregator.UserService).ServeHTTP)
	})

	return r
}

type middlewareStruct struct {
}

func NewAuthMiddleware() *middlewareStruct {
	return &middlewareStruct{}
}

func (m *middlewareStruct) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logged!")
		next.ServeHTTP(w, r)
		// id, err := strconv.Atoi(chi.URLParam(r, "articleID"))

		// if err != nil {
		// 	render.Status(r, http.StatusBadRequest)
		// 	render.JSON(w, r, http.StatusText(http.StatusBadRequest)) // TODO that does not return json :(
		// 	return
		// }

		// article, err := db.GetArticle(id)

		// if err != nil {
		// 	render.Status(r, http.StatusNotFound)
		// 	render.JSON(w, r, http.StatusText(http.StatusNotFound)) // TODO that does not return json :(
		// 	return
		// }

		// ctx := context.WithValue(r.Context(), "article", article)
		// next.ServeHTTP(w, r.WithContext(ctx))
	})
}
