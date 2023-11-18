package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ApiConfig struct {
	DB *gorm.DB
}

func main() {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Error loading .env file : %q", err)
	}

	port, ok := env["PORT"]
	if !ok {
		log.Fatal("PORT environment variable is not set")
	}

	dsn, ok := env["DSN"]
	if !ok {
		log.Fatal("DSN environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to establish connection with the database : %q", err)
	}

	err = db.AutoMigrate(&User{}, &Feed{}, &Post{})
	if err != nil {
		log.Fatalf("Auto Migration failed : %q", err)
	}
	err = db.SetupJoinTable(&User{}, "FeedFollows", &FeedFollow{})
	if err != nil {
		log.Fatalf("Failed to setup FeedFollows join table : %q", err)
	}

	cfg := &ApiConfig{DB: db}

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
	mainRouter.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	v1Router := chi.NewRouter()
	adminRouter := chi.NewRouter()

	adminRouter.Get("/readiness", func(w http.ResponseWriter, _ *http.Request) {
		type statusResponse struct {
			Status string `json:"status"`
		}
		respondWithJSON(w, http.StatusOK, statusResponse{http.StatusText(http.StatusOK)})
	})

	adminRouter.Get("/err", func(w http.ResponseWriter, _ *http.Request) {
		const code = http.StatusInternalServerError
		respondWithError(w, code, http.StatusText(code))
	})

	adminRouter.Get("/bad_json", func(w http.ResponseWriter, _ *http.Request) {
		type BadPayload struct {
			C complex128 `json:"complex"`
		}
		respondWithJSON(w, http.StatusOK, BadPayload{2 + 5i})
	})

	adminRouter.Get("/users", cfg.GetUsers)
	v1Router.Post("/users", cfg.CreateUser)
	v1Router.Get("/users", cfg.WithAuth(cfg.GetUserByApiKey))
	v1Router.Delete("/users", cfg.WithAuth(cfg.DeleteUser))

	v1Router.Post("/feeds", cfg.WithAuth(cfg.CreateFeed))
	v1Router.Get("/feeds", cfg.GetAllFeeds)
	v1Router.Delete("/feeds", cfg.WithAuth(cfg.DeleteFeed))

	v1Router.Post("/feed_follows", cfg.WithAuth(cfg.CreateFeedFollow))
	v1Router.Get("/feed_follows", cfg.WithAuth(cfg.GetFeedFollows))
	v1Router.Delete("/feed_follows", cfg.WithAuth(cfg.DeleteFeedFollow))

	v1Router.Get("/posts", cfg.WithAuth(cfg.GetPostsForUser))

	v1Router.Mount("/admin", adminRouter)
	mainRouter.Mount("/v1", v1Router)

	go cfg.StartScraping(3, 30*time.Second)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	log.Println("serving on port:", port)
	log.Fatal(server.ListenAndServe())
}
