package counter_test

import (
	"sync"
	"testing"

	"github.com/zoumas/lab/lgwt/sync/counter"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		c := counter.NewCounter()

		want := 3
		for i := 0; i < want; i++ {
			c.Inc()
		}
		assertCounter(t, c, want)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		want := 1_000
		c := counter.NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)
		for i := 0; i < want; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, c, want)
	})
}

func assertCounter(t testing.TB, c *counter.Counter, want int) {
	t.Helper()
	if got := c.Value(); got != want {
		t.Errorf("\ngot:\n%d\nwant:\n%d", got, want)
	}
}
