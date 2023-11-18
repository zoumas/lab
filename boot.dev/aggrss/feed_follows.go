package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"
)

func (cfg *ApiConfig) CreateFeedFollow(w http.ResponseWriter, r *http.Request, u database.User) {
	type requestBody struct {
		FeedID string `json:"feed_id"`
	}
	body := requestBody{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if body.FeedID == "" {
		respondWithError(w, http.StatusBadRequest, "feed_id field is empty")
		return
	}

	feedID, err := uuid.Parse(body.FeedID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now().UTC()
	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		FeedID:    feedID,
		UserID:    u.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, feedFollow)
}

func (cfg *ApiConfig) DeleteFeedFollow(w http.ResponseWriter, r *http.Request, u database.User) {
	feedFollowIDString := chi.URLParam(r, "feed_follow_id")

	feedFollowID, err := uuid.Parse(feedFollowIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: u.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (cfg *ApiConfig) GetFeedFollows(w http.ResponseWriter, r *http.Request, u database.User) {
	feedFollows, err := cfg.DB.GetFeedFollows(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, feedFollows)
}
