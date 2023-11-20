package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ApiConfig struct {
	fileServerHits int
}

func main() {
	port := "8080"

	cfg := &ApiConfig{}
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

	apiRouter.Post("/validate_chirp", ValidateChirp)

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
