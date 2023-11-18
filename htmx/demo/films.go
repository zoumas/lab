package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/zoumas/lab/htmx/demo/internal/database"
)

func (s *Shared) MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbFilms, err := s.DB.GetFilms(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string][]database.Film{"Films": dbFilms})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Shared) CreateFilm(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type requestBody struct {
		Title    string `json:"title"`
		Director string `json:"director"`
	}
	reqBody := requestBody{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, ErrorResponse{
			Error:   "failed to decode request body",
			Details: err.Error(),
		})
		return
	}

	film, err := s.DB.CreateFilm(r.Context(), database.CreateFilmParams{
		Title:    reqBody.Title,
		Director: reqBody.Director,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "failed to create film",
			Details: err.Error(),
		})
		return
	}

	respondWithJSON(w, http.StatusCreated, film)
}

func (s *Shared) GetFilms(w http.ResponseWriter, r *http.Request) {
	films, err := s.DB.GetFilms(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "failed to retrieve films",
			Details: err.Error(),
		})
		return
	}

	respondWithJSON(w, http.StatusOK, films)
}

func (s *Shared) CreateFilmFromHTMX(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	film, err := s.DB.CreateFilm(r.Context(), database.CreateFilmParams{
		Title:    title,
		Director: director,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "film-list-element", film)
}
