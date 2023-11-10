package main

import (
	"fmt"
	"net/http"

	"github.com/zoumas/lab/boot.dev/aggrss/internal/auth"
	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (c *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(
				w,
				http.StatusForbidden,
				fmt.Sprintf("Authentication error: %q", err),
			)
			return
		}

		user, err := c.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(
				w,
				http.StatusInternalServerError,
				fmt.Sprintf("Failed to retrieve user: %q", err),
			)
			return
		}

		handler(w, r, user)
	}
}
