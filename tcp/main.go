package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Message struct {
	Sender  string
	Payload []byte
}

type Server struct {
	Addr     string
	ln       net.Listener
	quit     chan struct{}
	messages chan Message
}

func NewServer(addr string) *Server {
	return &Server{
		Addr:     addr,
		quit:     make(chan struct{}),
		messages: make(chan Message, 10),
	}
}

func (s *Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln

	defer s.ln.Close()
	defer close(s.messages)

	go s.acceptLoop()

	<-s.quit
	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	log.Println(conn.RemoteAddr(), "connected")
	defer log.Println(conn.RemoteAddr(), "disconnected")
	defer conn.Close()

	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println("read error:", err)
			continue
		}

		s.messages <- Message{
			Sender:  conn.RemoteAddr().String(),
			Payload: buf[:n],
		}
		conn.Write(buf[:n])
	}
}

func main() {
	const addr = "127.0.0.1:3000"
	server := NewServer(addr)
	log.Println(addr)

	go func() {
		for m := range server.messages {
			fmt.Printf("%s: %s", m.Sender, m.Payload)
		}
	}()

	log.Fatal(server.ListenAndServe())
}
