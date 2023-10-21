package countdown_test

import (
	"bytes"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/zoumas/lab/lgwt/mocking/countdown"
)

func ExampleCountdown() {
	duration := time.Duration(0)
	sleeper := countdown.NewConfigurableSleeper(duration, time.Sleep)
	countdown.Countdown(os.Stdout, sleeper)

	// Output:
	// 3
	// 2
	// 1
	// Go!
}

func TestCountdown(t *testing.T) {
	t.Run("printing", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpyCountdownOperations{}

		countdown.Countdown(buffer, sleeper)

		want := `3
2
1
Go!
`
		got := buffer.String()
		assertStrings(t, got, want)
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &SpyCountdownOperations{}

		countdown.Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(spySleeperPrinter.Calls, want) {
			t.Errorf("\nmismatched calls\ngot: %v\nwant: %v", spySleeperPrinter.Calls, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}

	sleeper := countdown.NewConfigurableSleeper(sleepTime, spyTime.Sleep)
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("\nshould have slept for %v\nbut slept for %v", sleepTime, spyTime.durationSlept)
	}
}

// type SpySleeper struct {
// 	Calls int
// }
//
// func (s *SpySleeper) Sleep() {
// 	s.Calls++
// }

type SpyCountdownOperations struct {
	Calls []string
}

const (
	sleep = "sleep"
	write = "write"
)

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
