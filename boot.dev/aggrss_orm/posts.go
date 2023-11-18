package main

import (
	"net/http"
	"strconv"
)

func (cfg *ApiConfig) GetPostsForUser(w http.ResponseWriter, r *http.Request, u User) {
	limit := 10

	limitQueryParam := r.URL.Query().Get("limit")
	if limitQueryParam != "" {
		var err error
		limit, err = strconv.Atoi(limitQueryParam)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "malformed limit query parameter")
		}
	}

	var posts []Post
	query := `
  SELECT * 
  FROM posts p JOIN feed_follows ff ON p.feed_id = ff.feed_id 
  WHERE ff.user_id = ?
  LIMIT ?`
	if err := cfg.DB.Raw(query, u.ID.String(), limit).Find(&posts).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
