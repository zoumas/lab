package webcheck

import "sync"

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites checks the status of a list of URLs.
// It returns a map of each URL checked to a boolean value:
// true for a good response; false for a bad response.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultChannel <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

// CheckWebsites checks the status of a list of URLs.
// It returns a map of each URL checked to a boolean value:
// true for a good response; false for a bad response.
func checkWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	l := len(urls)
	results := make(map[string]bool, l)

	var wg sync.WaitGroup
	wg.Add(l)
	var mu sync.Mutex
	for _, url := range urls {
		go func(url string) {
			mu.Lock()
			results[url] = wc(url)
			mu.Unlock()

			wg.Done()
		}(url)
	}
	wg.Wait()

	return results
}
