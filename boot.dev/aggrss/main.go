package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"

	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Error loading .env file: %q", err)
	}

	port, set := env["PORT"]
	if !set {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL, set := env["DB"]
	if !set {
		log.Fatal("DB environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)

	cfg := &ApiConfig{DB: database.New(db)}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(middleware.Logger)

	v1Router := chi.NewRouter()
	v1Router.Get("/users", cfg.WithAuth(cfg.GetUserByApiKey))
	v1Router.Post("/users", cfg.CreateUser)

	v1Router.Get("/feeds", cfg.GetFeeds)
	v1Router.Post("/feeds", cfg.WithAuth(cfg.CreateFeed))

	v1Router.Get("/feed_follows", cfg.WithAuth(cfg.GetFeedFollows))
	v1Router.Post("/feed_follows", cfg.WithAuth(cfg.CreateFeedFollow))
	v1Router.Delete("/feed_follows/{feed_follow_id}", cfg.WithAuth(cfg.DeleteFeedFollow))

	v1Router.Get("/posts", cfg.WithAuth(cfg.GetPostsForUser))

	adminRouter := chi.NewRouter()
	adminRouter.HandleFunc("/readiness", func(w http.ResponseWriter, _ *http.Request) {
		type statusResponse struct {
			Status string `json:"status"`
		}

		respondWithJSON(
			w,
			http.StatusOK,
			statusResponse{http.StatusText(http.StatusOK)},
		)
	})
	adminRouter.HandleFunc("/err", func(w http.ResponseWriter, _ *http.Request) {
		respondWithError(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	})
	adminRouter.Get("/users", cfg.GetUsers)

	v1Router.Mount("/admin", adminRouter)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go startScraping(cfg.DB, 5, time.Minute/2)

	log.Println("serving on port:", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
