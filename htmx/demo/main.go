package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func GetFilms(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}

	tmpl.Execute(w, films)
}

func AddFilm(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)

	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = tmpl.ExecuteTemplate(w, "film-list-element", Film{title, director})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func main() {
	const port = "8000"

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
		respondWithJSON(
			w,
			http.StatusOK,
			statusResponse{http.StatusText(http.StatusOK)},
		)
	})
	router.Get("/err", func(w http.ResponseWriter, _ *http.Request) {
		respondWithError(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	})

	router.Get("/films", GetFilms)
	router.Post("/films", AddFilm)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Println("serving on port:", port)
	log.Fatal(server.ListenAndServe())
}
