package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal %v to JSON. Error: %s\n", payload, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if statusCode != http.StatusOK {
		if statusCode >= http.StatusInternalServerError {
			log.Printf("Respond with %v status code\n", statusCode)
		}
		w.WriteHeader(statusCode)
	}
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
	type errorResponse struct {
		Message string `json:"error"`
	}

	respondWithJSON(w, statusCode, errorResponse{errorMessage})
}
