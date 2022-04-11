package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/karmek-k/frontpad/pkg/router/routes"
	"github.com/karmek-k/frontpad/pkg/websocket"
)

// CreateRouter creates a new chi router
func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	m := websocket.CreateWebSocket()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.AllowAll().Handler) // TODO: change when deploying
	
	r.Route("/api", func(r chi.Router) {
		r.Route("/session", func(r chi.Router) {
			r.Post("/", routes.SessionCreate)
			r.Delete("/{id}", routes.SessionDelete)
		})
	})

	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	return r
}