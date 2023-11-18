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

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(f database.Feed) Feed {
	var t *time.Time
	if f.LastFetchedAt.Valid {
		// copy local var so we don't have to keep the db object in memory
		last_fetched_at := f.LastFetchedAt.Time
		t = &last_fetched_at
	}

	return Feed{
		ID:            f.ID,
		CreatedAt:     f.CreatedAt,
		UpdatedAt:     f.UpdatedAt,
		Name:          f.Name,
		Url:           f.Url,
		LastFetchedAt: t,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	fs := make([]Feed, 0, len(feeds))
	for _, f := range feeds {
		fs = append(fs, databaseFeedToFeed(f))
	}
	return fs
}

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
			fmt.Sprint("unable to follow created feed:", err),
		)
		return
	}

	type response struct {
		Feed       Feed                `json:"feed"`
		FeedFollow database.FeedFollow `json:"feed_follow"`
	}

	respondWithJSON(w, http.StatusCreated, response{databaseFeedToFeed(feed), feedFollow})
}

func (api *Api) DeleteFeed(w http.ResponseWriter, r *http.Request, u database.User) {
	feedID := chi.URLParam(r, "feed_id")

	feedUUID, err := uuid.Parse(feedID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error parsing url parameter")
		return
	}

	err = api.DB.DeleteFeed(r.Context(), database.DeleteFeedParams{
		ID:     feedUUID,
		UserID: u.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to delete feed: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
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

	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
