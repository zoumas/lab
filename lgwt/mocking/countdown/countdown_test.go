package countdown_test

import (
	"bytes"
	"slices"
	"testing"
	"time"

	"github.com/zoumas/lab/lgwt/mocking/countdown"
)

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
	write = "write"
	sleep = "sleep"
)

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("printing", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		sleeper := &SpyCountdownOperations{}

		countdown.Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
		}
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

		if got := spySleeperPrinter.Calls; !slices.Equal(got, want) {
			t.Errorf("\nwanted calls:\n%v\ngot:\n%v", want, got)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}

	sleeper := countdown.NewConfigurableSleeper(sleepTime, spyTime.Sleep)
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("\nshould have slept for %v\nslept for %v", sleepTime, spyTime.durationSlept)
	}
}
