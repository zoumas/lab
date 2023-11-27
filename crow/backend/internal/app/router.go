package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func ConfiguredRouter(app *App) *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{app.Env.CorsOrigin},
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
	router.Use(middleware.Recoverer)

	router.Get("/healthz", HealthCheck)

	usersRouter := chi.NewRouter()
	usersRouter.Post("/", app.UserCreateHandler)

	router.Mount("/users", usersRouter)

	return router
}
