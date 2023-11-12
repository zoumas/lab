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
	"github.com/zoumas/lab/boot.dev/rssaggr/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	port, ok := os.LookupEnv("PORT")
	switch {
	case !ok:
		log.Fatal("PORT environment variable is not set")
	case port == "":
		log.Fatal("PORT environment variable is empty")
	}

	dbURL, ok := os.LookupEnv("DB_URL")
	switch {
	case !ok:
		log.Fatal("DB_URL environment variable is not set")
	case dbURL == "":
		log.Fatal("DB_URL environment variable is empty")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %q", err)
	}

	api := Api{
		DB: database.New(conn),
	}

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
	v1Router.Get("/readiness", readinessHandler)
	v1Router.Get("/err", errHandler)

	v1Router.Post("/users", api.CreateUser)
	v1Router.Get("/users", api.WithAuth(api.GetUserByApiKey))

	v1Router.Post("/feeds", api.WithAuth(api.CreateFeed))
	v1Router.Get("/feeds", api.GetFeeds)

	v1Router.Post("/feed_follows", api.WithAuth(api.CreateFeedFollow))
	v1Router.Delete("/feed_follows/{feedFollowID}", api.WithAuth(api.DeleteFeedFollow))
	v1Router.Get("/feed_follows", api.WithAuth(api.GetFeedFollows))

	mainRouter.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	log.Println("serving on port:", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server error: %q", err)
	}
}
