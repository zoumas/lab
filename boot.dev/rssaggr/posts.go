package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/rssaggr/internal/database"
)

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(p database.Post) Post {
	var description *string
	if p.Description.Valid {
		d := p.Description.String
		description = &d
	}

	return Post{
		ID:          p.ID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Url:         p.Url,
		Title:       p.Title,
		Description: description,
		PublishedAt: p.PublishedAt,
		FeedID:      p.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	ps := make([]Post, 0, len(posts))
	for _, p := range posts {
		ps = append(ps, databasePostToPost(p))
	}
	return ps
}

func (api *Api) GetPostsForUser(w http.ResponseWriter, r *http.Request, u database.User) {
	limitStr := r.URL.Query().Get("limit")

	if limitStr == "" {
		posts, err := api.DB.GetAllPostsForUser(r.Context(), u.ID)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprint("failed to retrieve posts:", err))
			return
		}

		respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to parse query parameters:"+err.Error())
		return
	}

	posts, err := api.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: u.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("failed to retrieve posts:", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
