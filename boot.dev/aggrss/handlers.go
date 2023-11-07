package main

import "net/http"

func readinessHandler(w http.ResponseWriter, _ *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}

func errorHandler(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, http.StatusBadRequest, "Something went wrong")
}
