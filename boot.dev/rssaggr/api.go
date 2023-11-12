package main

import (
	"net/http"

	"github.com/zoumas/lab/boot.dev/rssaggr/internal/auth"
	"github.com/zoumas/lab/boot.dev/rssaggr/internal/database"
)

type Api struct {
	DB *database.Queries
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (api *Api) WithAuth(h authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := api.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		h(w, r, user)
	}
}
