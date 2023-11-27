package app

import (
	"net/http"

	"github.com/zoumas/lab/crow/backend/internal/user"
)

func (app *App) UserDeleteHandler(w http.ResponseWriter, r *http.Request, u user.User) {
	if err := app.userService.Delete(u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
