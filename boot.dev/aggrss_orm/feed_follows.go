package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) CreateFeedFollow(w http.ResponseWriter, r *http.Request, u User) {
	type requestBody struct {
		FeedID string `json:"feed_id"`
	}

	reqBody := requestBody{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	feedID, err := uuid.Parse(reqBody.FeedID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	feedFollow := FeedFollow{}
	feedFollow.UserID = u.ID
	feedFollow.FeedID = &feedID

	if err := cfg.DB.Create(&feedFollow).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, feedFollow)
}

func (cfg *ApiConfig) DeleteFeedFollow(w http.ResponseWriter, r *http.Request, u User) {
	type requestBody struct {
		FeedID string `json:"feed_id"`
	}

	reqBody := requestBody{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	feedID, err := uuid.Parse(reqBody.FeedID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var feedFollow FeedFollow
	if err := cfg.DB.Where("feed_id = ? AND user_id = ?", feedID.String(), u.ID.String()).First(&feedFollow).Error; err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	if err := cfg.DB.Delete(&feedFollow).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (cfg *ApiConfig) GetFeedFollows(w http.ResponseWriter, r *http.Request, u User) {
	var feedFollows []FeedFollow
	if err := cfg.DB.Where("user_id = ?", u.ID.String()).Find(&feedFollows).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, feedFollows)
}
