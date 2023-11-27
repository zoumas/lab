package app

import (
	"net/http"

	"github.com/zoumas/lab/crow/backend/internal/user"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, u user.User)

func (app *App) WithPassword(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			respondWithError(w, http.StatusUnauthorized, "Basic Auth required")
			return
		}

		u, err := app.userService.GetByName(username)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if !passwordsMatch(u.Password, password) {
			respondWithError(w, http.StatusUnauthorized, "wrong password")
			return
		}

		handler(w, r, u)
	}
}
