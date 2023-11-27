package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *App) UserRetrieveHandler(w http.ResponseWriter, r *http.Request) {
	users, err := app.userService.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func (app *App) UserGetByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		respondWithError(w, http.StatusBadRequest, "missing URL parameter")
		return
	}

	user, err := app.userService.GetByName(name)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}
