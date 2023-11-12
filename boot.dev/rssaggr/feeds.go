package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/rssaggr/internal/database"
)

func (api *Api) CreateFeed(w http.ResponseWriter, r *http.Request, u database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprint("unable to process request body:", err),
		)
		return
	}

	now := time.Now().UTC()
	feed, err := api.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      params.Name,
		Url:       params.Url,
		UserID:    u.ID,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprint("unable to create feed:", err),
		)
		return
	}

	feedFollow, err := api.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    u.ID,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprint("unable to create follow feed:", err),
		)
		return
	}

	type response struct {
		Feed       database.Feed       `json:"feed"`
		FeedFollow database.FeedFollow `json:"feed_follow"`
	}

	respondWithJSON(w, http.StatusCreated, response{feed, feedFollow})
}

func (api *Api) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := api.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprint("unable to retrieve feeds:", err),
		)
	}

	respondWithJSON(w, http.StatusOK, feeds)
}
