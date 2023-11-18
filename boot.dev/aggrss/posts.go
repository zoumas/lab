package main

import (
	"net/http"
	"strconv"

	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"
)

func (cfg *ApiConfig) GetPostsForUser(w http.ResponseWriter, r *http.Request, u database.User) {
	limitString := r.URL.Query().Get("limit")
	if limitString == "" {
		cfg.GetAllPostsForUser(w, r, u)
		return
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	posts, err := cfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: u.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}

func (cfg *ApiConfig) GetAllPostsForUser(w http.ResponseWriter, r *http.Request, u database.User) {
	posts, err := cfg.DB.GetAllPostsForUser(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
