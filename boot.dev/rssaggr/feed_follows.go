package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/rssaggr/internal/database"
)

func (api *Api) CreateFeedFollow(w http.ResponseWriter, r *http.Request, u database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}
	params := parameters{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("unable to process request:", err))
		return
	}

	feedUUID, err := uuid.Parse(params.FeedID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("unable to parse feed id:", err))
		return
	}

	now := time.Now().UTC()
	feedFollow, err := api.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feedUUID,
		UserID:    u.ID,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprint("unable to create feed follow:", err),
		)
		return
	}

	respondWithJSON(w, http.StatusCreated, feedFollow)
}

func (api *Api) DeleteFeedFollow(w http.ResponseWriter, r *http.Request, u database.User) {
	feedFollowID := chi.URLParam(r, "feedFollowID")
	feedFollowUUID, err := uuid.Parse(feedFollowID)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprint("unable to parse feed follow id:", err.Error()),
		)
		return
	}

	err = api.DB.DeleteFeedFollow(r.Context(), feedFollowUUID)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprint("unable to delete feed follow", err),
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *Api) GetFeedFollows(w http.ResponseWriter, r *http.Request, u database.User) {
	feedFollows, err := api.DB.GetFeedFollows(r.Context(), u.ID)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprint("unable to get followed feeds:", err),
		)
		return
	}
	respondWithJSON(w, http.StatusOK, feedFollows)
}
