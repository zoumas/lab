package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func ValidateChirp(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Body string `json:"body"`
	}
	reqBody := RequestBody{}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const MaxChirpLength = 140
	if len(reqBody.Body) > MaxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	profaneWords := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}
	replaceWith := "****"
	cleanedBody := ReplaceProfaneWords(reqBody.Body, profaneWords, replaceWith)

	type CleanedResponse struct {
		CleanedBody string `json:"cleaned_body"`
	}
	respondWithJSON(w, http.StatusOK, CleanedResponse{cleanedBody})
}

func ReplaceProfaneWords(body string, profaneWords []string, replaceWith string) string {
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
