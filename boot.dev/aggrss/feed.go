package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"
)

func (c *config) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
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
		return
	}

	feedFollow, err := c.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to follow feed: %q", err),
		)
		return
	}

	type response struct {
		Feed       database.Feed       `json:"feed"`
		FeedFollow database.FeedFollow `json:"feed_follow"`
	}

	respondWithJSON(w, http.StatusCreated, response{feed, feedFollow})
}

func (c *config) listFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := c.DB.ListFeeds(r.Context())
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to retrieve feeds: %q", err),
		)
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}
