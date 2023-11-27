package app

import (
	"encoding/json"
	"net/http"

	"github.com/zoumas/lab/crow/backend/internal/user"
)

func (app *App) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "failed to decode request body")
		return
	}

	if err := validateName(body.Name); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validatePassword(body.Name); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := hashPassword(body.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unacceptable password")
		return
	}

	user := user.User{
		Name:     body.Name,
		Password: hashedPassword,
	}
	if err := app.userService.Save(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type ErrUser string

func (e ErrUser) Error() string {
	return string(e)
}

const (
	ErrUserNameEmpty     = ErrUser("User's name is empty")
	ErrUserPasswordEmpty = ErrUser("User's password is empty")
)

func validateName(name string) error {
	if name == "" {
		return ErrUserNameEmpty
	}
	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return ErrUserPasswordEmpty
	}
	return nil
}
