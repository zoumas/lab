package main

import (
	"html/template"
	"net/http"
)

func (cfg *ApiConfig) IncrementMetrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits++

		h.ServeHTTP(w, r)
	})
}

func (cfg *ApiConfig) ReportMetrics(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("metrics.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, cfg.fileServerHits)
}

func (cfg *ApiConfig) ResetMetrics(w http.ResponseWriter, _ *http.Request) {
	cfg.fileServerHits = 0
	w.WriteHeader(http.StatusOK)
}
