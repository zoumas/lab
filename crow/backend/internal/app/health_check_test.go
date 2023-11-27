package app_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zoumas/lab/crow/backend/internal/app"
)

func TestHealthCheck(t *testing.T) {
	t.Run("returns 200 status code", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(app.HealthCheck))
		defer server.Close()

		resp, err := http.Get(server.URL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		got := resp.StatusCode
		want := http.StatusOK
		if got != want {
			t.Errorf("\ngot: %d\nwant: %d", got, want)
		}
	})
}
