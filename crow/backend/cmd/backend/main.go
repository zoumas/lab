package main

import (
	"log"

	"github.com/zoumas/lab/crow/backend/internal/app"
	"github.com/zoumas/lab/crow/backend/internal/env"
)

func main() {
	env, err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	server := app.ConfiguredServer(env, app.ConfiguredRouter())
	log.Printf("backend serving from %s", env.Addr)
	app.New(env).Run(server)
}
