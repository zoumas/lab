package webracer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

// Race takes two URLs and "races" them by hitting them
// with an HTTP GET and returning the URL which returned first.
// If none of them return within 10 seconds it returns an error.
func Racer(x, y string) (winner string, err error) {
	return ConfigurableRacer(x, y, tenSecondTimeout)
}

func ConfigurableRacer(x, y string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(x):
		return x, nil
	case <-ping(y):
		return y, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for %s and %s", x, y)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// Race takes two URLs and "races" them by hitting them
// with an HTTP GET and returning the URL which returned first.
// If none of them return within 10 seconds it returns an error.
// func Race(x, y string) (winner string) {
// 	if responseTime(x) < responseTime(y) {
// 		return x
// 	}
// 	return y
// }
//
// func responseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
