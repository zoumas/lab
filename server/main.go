package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	// Basic Rate Limiting
	// Handle a Connection every 200 milliseconds
	t := time.Tick(200 * time.Millisecond)
	for ; ; <-t {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) error {
	defer conn.Close()

	log.Println("Connection established!", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	requestLine, err := strings.TrimSpace(scanner.Text()), scanner.Err()
	if err != nil && err != io.EOF {
		return err
	}

	log.Println(requestLine)

	var statusLine, filename string
	request := strings.Fields(requestLine)
	path := request[1]

	switch path {
	case "/":
		statusLine = "HTTP/1.1 200 OK"
		filename = "hello.html"
	case "/sleep":
		time.Sleep(5 * time.Second)
		statusLine = "HTTP/1.1 200 OK"
		filename = "hello.html"
	default:
		statusLine = "HTTP/1.1 404 NOT FOUND"
		filename = "404.html"
	}

	contents, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	length := len(contents)

	response := fmt.Sprintf(
		"%s\r\nContent-Length: %d\r\n\r\n%s",
		statusLine, length, contents,
	)
	_, err = conn.Write([]byte(response))
	if err != nil {
		return err
	}

	return nil
}
