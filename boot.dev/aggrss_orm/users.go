package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Name string `json:"name"`
	}

	reqBody := requestBody{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := User{
		Name: reqBody.Name,
	}
	if err := cfg.DB.Create(&user).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (cfg *ApiConfig) GetUserByApiKey(w http.ResponseWriter, r *http.Request, u User) {
	respondWithJSON(w, http.StatusOK, u)
}

func (cfg *ApiConfig) GetUsers(w http.ResponseWriter, _ *http.Request) {
	var users []User
	if err := cfg.DB.Find(&users).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (cfg *ApiConfig) DeleteUser(w http.ResponseWriter, r *http.Request, u User) {
	if err := cfg.DB.Delete(&u).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
