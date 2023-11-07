package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dat)
}

type errResponse struct {
	ErrorMessage string `json:"error"`
}

func respondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
	if statusCode >= http.StatusInternalServerError {
		log.Println("Responding with 5XX error:", errorMessage)
	}

	respondWithJSON(w, statusCode, errResponse{errorMessage})
}
