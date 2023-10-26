package webracer_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/zoumas/lab/lgwt/select/webracer"
)

func TestRacer(t *testing.T) {
	t.Run("returns the url that completes first", func(t *testing.T) {
		slowServer := newDelayedTestServer(20 * time.Millisecond)
		fastServer := newDelayedTestServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := webracer.Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("\nUnexpected error:\n%q", err)
		}

		if got != want {
			t.Errorf("\ngot:\n%q, want:\n%q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s",
		func(t *testing.T) {
			s := newDelayedTestServer(25 * time.Millisecond)
			defer s.Close()
			url := s.URL

			_, err := webracer.ConfigurableRacer(url, url, 20*time.Millisecond)
			if err == nil {
				t.Error("Expected an error")
			}
		})
}

func newDelayedTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, _ *http.Request) {
				time.Sleep(delay)
				w.WriteHeader(http.StatusOK)
			}))
}
