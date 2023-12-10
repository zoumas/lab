package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/zoumas/lab/tcp"
)

func main() {
	server := tcp.NewServer(":3000")

	go func() {
		for msg := range server.MsgChan {
			payload := strings.TrimSpace(string(msg.Payload))
			fmt.Printf("Received message from connection (%s): %s\n", msg.From, payload)
		}
	}()

	log.Fatal(server.Start())
}
