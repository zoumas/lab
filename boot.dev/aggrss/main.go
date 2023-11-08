package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT environment variable is not set")
	}

	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware.Logger)
	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", http.HandlerFunc(readinessHandler))
	v1Router.Get("/err", http.HandlerFunc(errHandler))

	mainRouter.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	log.Println("serving on port:", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
