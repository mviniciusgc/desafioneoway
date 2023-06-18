package router

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Handlers(s *HandlerServices) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/api/external/uploadfile", func(r chi.Router) {
		r.Post("/", handle(s))
	})

	return r

}
