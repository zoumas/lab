package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"golang.org/x/net/websocket"
)

type WebSocketServer struct {
	Conns map[*websocket.Conn]struct{}
	mu    *sync.RWMutex
}

func NewServer() *WebSocketServer {
	return &WebSocketServer{
		Conns: make(map[*websocket.Conn]struct{}),
		mu:    &sync.RWMutex{},
	}
}

func (s *WebSocketServer) HandleWS(ws *websocket.Conn) {
	log.Println("new incoming Web Socket Connection:", ws.RemoteAddr())

	s.mu.Lock()
	s.Conns[ws] = struct{}{}
	s.mu.Unlock()

	s.readLoop(ws)
}

func (s *WebSocketServer) readLoop(ws *websocket.Conn) {
	const defaultBufferLength = 1024
	buf := make([]byte, defaultBufferLength)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println(ws.RemoteAddr(), "connection dropped")
				break
			}
			fmt.Println(ws.RemoteAddr(), "read error", err)
			continue
		}

		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *WebSocketServer) broadcast(msg []byte) {
	for ws := range s.Conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(msg); err != nil {
				log.Println(ws.LocalAddr(), "write error", err)
			}
		}(ws)
	}
}

func main() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	wsServer := NewServer()
	router.Handle("/ws", AllowContentSecurityPolicy(websocket.Handler(wsServer.HandleWS)))

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	server.ListenAndServe()
}

func AllowContentSecurityPolicy(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Security-Policy", "*")
		handler.ServeHTTP(w, r)
	})
}
