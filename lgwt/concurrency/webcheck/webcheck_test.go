package webcheck_test

import (
	"maps"
	"testing"
	"time"

	"github.com/zoumas/lab/lgwt/concurrency/webcheck"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := webcheck.CheckWebsites(mockWebsiteChecker, websites)

	assertMaps(t, got, want)
}

func assertMaps[K, V comparable](t testing.TB, got, want map[K]V) {
	t.Helper()

	if !maps.Equal(got, want) {
		t.Errorf("\ngot:\n%#v\nwant:\n%#v", got, want)
	}
}

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		webcheck.CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

// sequentiatly =>  2.04 sec
// concurrently =>  0.02 sec

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
