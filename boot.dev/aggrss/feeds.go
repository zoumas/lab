package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"
)

func (cfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, u database.User) {
	type requestBody struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	body := requestBody{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	switch "" {
	case body.Name:
		respondWithError(w, http.StatusBadRequest, "name field is empty")
		return
	case body.Url:
		respondWithError(w, http.StatusBadRequest, "url field is empty")
		return
	}

	now := time.Now().UTC()

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      body.Name,
		Url:       body.Url,
		UserID:    u.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		FeedID:    feed.ID,
		UserID:    u.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	type responseBody struct {
		Feed       database.Feed       `json:"feed"`
		FeedFollow database.FeedFollow `json:"feed_follow"`
	}
	respondWithJSON(w, http.StatusCreated, responseBody{feed, feedFollow})
}

func (cfg *ApiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}
