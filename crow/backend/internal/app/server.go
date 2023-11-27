package app

import (
	"net/http"

	"github.com/zoumas/lab/crow/backend/internal/env"
)

func ConfiguredServer(env *env.Env, router http.Handler) *http.Server {
	return &http.Server{
		Addr:    env.Addr,
		Handler: router,
	}
}
