package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/zoumas/lab/htmx/demo/internal/database"

	_ "github.com/lib/pq"
)

type Shared struct {
	DB *database.Queries
}

func main() {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalf("error loading .env file : %q", err)
	}

	port, ok := env["PORT"]
	if !ok {
		log.Fatal("PORT is not set in .env")
	}

	dsn, ok := env["DSN"]
	if !ok {
		log.Fatal("DSN is not set in .env")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database : %q", err)
	}
	s := Shared{DB: database.New(db)}

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

	router.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		type statusResponse struct {
			Status string `json:"status"`
		}
		respondWithJSON(w, http.StatusOK, statusResponse{http.StatusText(http.StatusOK)})
	})
	router.Get("/err", func(w http.ResponseWriter, _ *http.Request) {
		respondWithError(w, http.StatusInternalServerError, ErrorResponse{
			Error:   http.StatusText(http.StatusInternalServerError),
			Details: "This endpoint always returns an error",
		})
	})

	router.Get("/", s.MainPage)

	router.Post("/films", s.CreateFilm)
	router.Get("/films", s.GetFilms)

	router.Post("/add-film", s.CreateFilmFromHTMX)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("serving on port:", port)
	log.Fatal(server.ListenAndServe())
}
