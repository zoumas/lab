package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"
)

func (c *apiConfig) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "name shouldn't be empty")
		return
	}
	if params.Url == "" {
		respondWithError(w, http.StatusBadRequest, "url shouldn't be empty")
		return
	}

	feed, err := c.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to create feed: %q", err),
		)
	}

	respondWithJSON(w, http.StatusCreated, feed)
}

func (c *apiConfig) getUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}

func (c *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// should be checked by the database?
	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Field 'name' must not be empty")
		return
	}

	user, err := c.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

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
