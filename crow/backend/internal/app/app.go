package app

import (
	"log"
	"net/http"

	"github.com/zoumas/lab/crow/backend/internal/env"
)

type App struct {
	Env *env.Env
}

func New(env *env.Env) *App {
	return &App{Env: env}
}

// App calls the ListenAndServe() method on the server instance that was passed in.
// In a later step this is where we will implement graceful-shutdown.
func (app *App) Run(server *http.Server) {
	log.Fatal(server.ListenAndServe())
}
