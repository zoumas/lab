package app

import (
	"log"

	"github.com/zoumas/lab/crow/backend/internal/database"
	"github.com/zoumas/lab/crow/backend/internal/env"
	"gorm.io/gorm"
)

type App struct {
	Env         *env.Env
	DB          *gorm.DB
	userService *UserService
}

func New(env *env.Env, db *gorm.DB) *App {
	return &App{
		Env:         env,
		DB:          db,
		userService: NewUserService(*database.NewGormUserRepo(db)),
	}
}

// App calls the ListenAndServe() method on the server instance that was passed in.
// In a later step this is where we will implement graceful-shutdown.
func (app *App) Run() {
	router := ConfiguredRouter(app)
	server := ConfiguredServer(app.Env, router)
	log.Fatal(server.ListenAndServe())
}
