package main

import (
	"log"

	"github.com/zoumas/lab/crow/backend/internal/app"
	"github.com/zoumas/lab/crow/backend/internal/database"
	"github.com/zoumas/lab/crow/backend/internal/env"
)

func main() {
	env, err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(env.DSN)
	if err != nil {
		log.Fatal(err)
	}

	var dbVersion string
	db.Raw("SELECT version()").Scan(&dbVersion)
	log.Println("Database Version :", dbVersion)

	app := app.New(env, db)
	log.Printf("backend serving from %s", env.Addr)
	app.Run()
}
