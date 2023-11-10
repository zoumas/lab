package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"
)

func (c *config) followFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Failed to marshal request: %q", err),
		)
		return
	}

	feedFollow, err := c.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("Failed to follow feed: %q", err),
		)
		return
	}

	respondWithJSON(w, http.StatusCreated, feedFollow)
}

func (c *config) listFollowedFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	followedFeeds, err := c.DB.ListFeedFollowsByUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to retrieve followed fields: %q", err),
		)
		return
	}

	respondWithJSON(w, http.StatusOK, followedFeeds)
}

func (c *config) unfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDParam := chi.URLParam(r, "feed_follow_id")
	feedFollowID, err := uuid.Parse(feedFollowIDParam)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't parse feed_follow_id")
		return
	}

	err = c.DB.DeleteFeedFollowByID(r.Context(), database.DeleteFeedFollowByIDParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to unfollow feed: %q", err),
		)
	}

	w.WriteHeader(http.StatusAccepted)
}
