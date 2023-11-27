package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/zoumas/lab/crow/backend/internal/env"
)

func ConfiguredRouter(env *env.Env) *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{env.CorsOrigin},
		AllowedHeaders: []string{"*"},
		// Support the basic CRUD operations
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodPut,
			http.MethodDelete,
		},
		// Enable credentials for Cookies
		AllowCredentials: true,
	}))
	router.Use(middleware.Logger)

	router.Get("/healthz", HealthCheck)

	return router
}
