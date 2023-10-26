package webcheck_test

import (
	"maps"
	"testing"
	"time"

	"github.com/zoumas/lab/lgwt/concurrency/webcheck"
)

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"http://google.com",
		"http://blog.gyspydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gyspydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := webcheck.CheckWebsites(mockWebsiteChecker, urls)

	if !maps.Equal(got, want) {
		t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
	}
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

// sequentially: 2034070379 ns/op => 2.034 sec
// channels: 20937743 ns/op => 0.0209 sec. About 100 times faster
// WaitGroup & Mutext: 2032130619 ns/op => 2.03213062 sec. Just like the sequential

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func slowStubWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
