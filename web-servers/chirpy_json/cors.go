package main

import "net/http"

func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// w.Header().Set("Access-Control-Allow-Credentials", "false")
		// w.Header().Set("Access-Control-Max-Age", "300")

		// Options is used as a PREFLIGHT SANITY CHECK for methods such as PUT, PATCH, DELETE
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
