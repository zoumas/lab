package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func readinessHandler(w http.ResponseWriter, _ *http.Request) {
	type statusResponse struct {
		StatusText string `json:"status"`
	}

	respondWithJSON(
		w,
		http.StatusOK,
		statusResponse{http.StatusText(http.StatusOK)},
	)
}

func errorHandler(w http.ResponseWriter, _ *http.Request) {
	respondWithError(
		w,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
	)
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response for payload: %#v\nError: %q\n", payload, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
	type errorResponse struct {
		ErrorMessage string `json:"error"`
	}

	respondWithJSON(w, statusCode, errorResponse{errorMessage})
}
