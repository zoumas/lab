package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/rssaggr/internal/database"
)

func (api *Api) GetUserByApiKey(w http.ResponseWriter, r *http.Request, u database.User) {
	respondWithJSON(w, http.StatusOK, u)
}

func (api *Api) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprint("unable to decode request body:", err),
		)
		return
	}

	now := time.Now().UTC()
	user, err := api.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprint("unable to create user:", err),
		)
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}
