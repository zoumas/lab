package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, httpStatusCode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		type errMarshalResponse struct {
			Error   string `json:"error"`
			Details string `json:"details"`
		}
		respondWithJSON(w, http.StatusInternalServerError, errMarshalResponse{
			Error:   "Failed to marshal into JSON",
			Details: err.Error(),
		})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, httpStatusCode int, errorMessage string) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, httpStatusCode, errorResponse{errorMessage})
}
