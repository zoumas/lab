package main

import (
	"fmt"
	"time"
)

func worker(done chan struct{}) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- struct{}{}
}

func main() {
	done := make(chan struct{}, 1)
	go worker(done)
	<-done
}
