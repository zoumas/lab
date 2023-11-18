package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (cfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, u User) {
	type requestBody struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	reqBody := requestBody{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	feed := Feed{}
	feed.Name = reqBody.Name
	feed.Url = reqBody.Url
	feed.UserID = u.ID
	if err := cfg.DB.Create(&feed).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	feedFollow := FeedFollow{}
	feedFollow.UserID = u.ID
	feedFollow.FeedID = feed.ID
	if err := cfg.DB.Create(&feedFollow).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	type responseBody struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}
	respondWithJSON(w, http.StatusCreated, responseBody{feed, feedFollow})
}

func (cfg *ApiConfig) DeleteFeed(w http.ResponseWriter, r *http.Request, u User) {
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

	if err := cfg.DB.Where("user_id = ?", u.ID).Delete(&Feed{}, feedID).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (cfg *ApiConfig) GetAllFeeds(w http.ResponseWriter, _ *http.Request) {
	var feeds []Feed
	if err := cfg.DB.Find(&feeds).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}

func (cfg *ApiConfig) GetNextFeedsToFetch(limit int) ([]Feed, error) {
	var feeds []Feed
	query := `SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT ?`
	if err := cfg.DB.Raw(query, limit).Scan(&feeds).Error; err != nil {
		return nil, err
	}

	return feeds, nil
}

func (cfg *ApiConfig) MarkFeedAsFetched(feed Feed) error {
	return cfg.DB.Model(&feed).Update("last_fetched_at", gorm.Expr("NOW()")).Error
}
