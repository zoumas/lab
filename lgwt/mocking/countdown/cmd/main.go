package main

import (
	"os"
	"time"

	"github.com/zoumas/lab/lgwt/mocking/countdown"
)

func main() {
	sleeper := countdown.NewConfigurableSleeper(time.Second/2, time.Sleep)
	countdown.Countdown(os.Stdout, sleeper)
}
