package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env: %q", err)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("Failed to load PORT environment variable from .env")
	}
	if port == "" {
		log.Fatal("PORT environment variable shouldn't be empty")
	}

	dbURL, ok := os.LookupEnv("DB_URL")
	if !ok {
		log.Fatal("Failed to load DB_URL environment variable from .env")
	}
	if dbURL == "" {
		log.Fatal("DB_URL environment variable shouldn't be empty")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %q", err)
	}

	apiCfg := &apiConfig{DB: database.New(conn)}

	mainRouter := chi.NewRouter()
	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	mainRouter.Use(middleware.Logger)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("/err", errorHandler)
	v1Router.Post("/users", apiCfg.createUserHandler)
	v1Router.Get("/users", apiCfg.getUserByApiKey)

	mainRouter.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    "localhost:" + port,
		Handler: mainRouter,
	}

	log.Println("[aggrss]: serving on port:", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
