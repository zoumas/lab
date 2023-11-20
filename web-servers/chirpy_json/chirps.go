package main

import (
	"cmp"
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Chirp struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

func (cfg *ApiConfig) CreateChirp(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Body string `json:"body"`
	}
	reqBody := RequestBody{}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := validateChirp(reqBody.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	chirp, err := cfg.DB.CreateChirp(body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, chirp)
}

func (cfg *ApiConfig) GetAllChirps(w http.ResponseWriter, r *http.Request) {
	chirps := cfg.DB.GetChirps()
	slices.SortStableFunc(chirps, func(a, b Chirp) int {
		return cmp.Compare(a.ID, b.ID)
	})

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *ApiConfig) GetChirpByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to parse url parameter")
		return
	}

	chirp, ok := cfg.DB.Chirps[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, chirp)
}

func validateChirp(body string) (cleanedBody string, err error) {
	const MaxChirpLength = 140
	if len(body) > MaxChirpLength {
		return "", errors.New("chirp is too long")
	}

	profaneWords := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}
	replaceWith := "****"
	cleanedBody = replaceProfaneWords(body, profaneWords, replaceWith)

	return cleanedBody, nil
}

func replaceProfaneWords(body string, profaneWords []string, replaceWith string) string {
	// case insensitive matching
	split := strings.Split(body, " ")

	for i, word := range split {
		for _, profane := range profaneWords {
			if strings.ToLower(word) == profane {
				split[i] = replaceWith
			}
		}
	}

	return strings.Join(split, " ")
}
