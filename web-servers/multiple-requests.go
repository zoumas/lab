package main

import (
	"fmt"
	"log"
	"time"
)

type Request struct {
	Path string
}

func handleRequests(requests <-chan Request) {
	for r := range requests {
		go handleRequest(r)
	}
}

func handleRequest(r Request) {
	log.Println("Handling request for", r.Path)
	time.Sleep(500 * time.Millisecond)
	log.Println("Done with request for", r.Path)
}

func main() {
	requests := make(chan Request, 100)
	go handleRequests(requests)

	for i := 0; i < 4; i++ {
		requests <- Request{Path: fmt.Sprintf("/path/%d", i)}
		time.Sleep(25 * time.Millisecond)
	}

	time.Sleep(time.Second)
	log.Println("1 second passed, killing server")
}
