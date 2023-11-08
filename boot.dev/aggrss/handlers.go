package main

import "net/http"

func readinessHandler(w http.ResponseWriter, _ *http.Request) {
	type statusResponse struct {
		Status string `json:"status"`
	}

	respondWithJSON(w,
		http.StatusOK,
		statusResponse{Status: http.StatusText(http.StatusOK)})
}

func errHandler(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError))
}
