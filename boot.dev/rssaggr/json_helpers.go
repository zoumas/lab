package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, statusCode, errorResponse{errorMessage})
}

func readinessHandler(w http.ResponseWriter, _ *http.Request) {
	type statusResponse struct {
		Status string `json:"status"`
	}

	respondWithJSON(
		w,
		http.StatusOK,
		statusResponse{http.StatusText(http.StatusOK)},
	)
}

func errHandler(w http.ResponseWriter, _ *http.Request) {
	respondWithError(
		w,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
	)
}
