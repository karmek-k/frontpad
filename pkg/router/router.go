package router

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/karmek-k/frontpad/pkg/db"
)

var ctx = context.Background()

// CreateRouter creates a new chi router
func CreateRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		result, _ := db.RDB.Echo(ctx, "Hello redis").Result()

		w.Write([]byte(result))
	})

	return r
}