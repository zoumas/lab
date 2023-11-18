package main

import (
	"net/http"

	"github.com/zoumas/lab/boot.dev/aggrss/internal/auth"
	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) WithAuth(h authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKeyString, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKeyString)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		h(w, r, user)
	}
}
