package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ApiConfig struct {
	fileServerHits int
	DB             *DB
}

func main() {
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if *debug == true {
		log.Println("DEBUG MODE")
	}

	port := "8080"

	db, err := NewDB("database.json")
	if err != nil {
		log.Fatalf("Failed to connect to database : %q", err)
	}

	cfg := &ApiConfig{DB: db}
	mainRouter := chi.NewRouter()

	appRouter := chi.NewRouter()

	fileServer := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	appRouter.Handle("/", cfg.IncrementMetrics(fileServer))
	appRouter.Handle("/*", cfg.IncrementMetrics(fileServer))

	apiRouter := chi.NewRouter()
	apiRouter.Get("/reset", http.HandlerFunc(cfg.ResetMetrics))
	apiRouter.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// what a simple GET routing function would look like had we not used Chi
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	}))

	chirpRouter := chi.NewRouter()
	chirpRouter.Post("/", cfg.CreateChirp)
	chirpRouter.Get("/", cfg.GetAllChirps)
	chirpRouter.Get("/{id}", cfg.GetChirpByID)
	apiRouter.Mount("/chirps", chirpRouter)

	userRouter := chi.NewRouter()
	userRouter.Post("/", cfg.CreateUser)
	apiRouter.Mount("/users", userRouter)

	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", http.HandlerFunc(cfg.ReportMetrics))

	mainRouter.Mount("/app", appRouter)
	mainRouter.Mount("/api", apiRouter)
	mainRouter.Mount("/admin", adminRouter)

	handler := Logger(CORS(mainRouter))
	server := &http.Server{
		Handler: handler,
		Addr:    ":" + port,
	}

	log.Println("serving on port:", port)
	log.Fatal(server.ListenAndServe())
}
