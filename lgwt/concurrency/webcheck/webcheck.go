package webcheck

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites checks the status of a list of URLs.
// It returns a mpa of each URL checked to a boolean value: true for a good response; false for a bad response.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultChan <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChan
		results[r.string] = r.bool
	}

	return results
}
