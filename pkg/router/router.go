package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/karmek-k/frontpad/pkg/router/routes"
)

// CreateRouter creates a new chi router
func CreateRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	r.Route("/api", func(r chi.Router) {
		r.Route("/session", func(r chi.Router) {
			r.Post("/", routes.SessionCreate)
			r.Delete("/{id}", routes.SessionDelete)
		})
	})

	return r
}