package tcp

import (
	"io"
	"log"
	"net"
)

type Message struct {
	From    string
	Payload []byte
}

// Server is a custom TCP Server.
type Server struct {
	Addr     string
	Ln       net.Listener
	QuitChan chan struct{}
	MsgChan  chan Message
}

func NewServer(addr string) *Server {
	return &Server{
		Addr:     addr,
		QuitChan: make(chan struct{}),
		MsgChan:  make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	defer ln.Close()
	s.Ln = ln

	go s.acceptLoop()

	<-s.QuitChan
	close(s.MsgChan)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.Ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		log.Println("Accepted a new connection from", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("Connection", conn.RemoteAddr(), "droppped")
				return
			}
			log.Println("Read error:", err)
			continue
		}

		s.MsgChan <- Message{
			From:    conn.RemoteAddr().String(),
			Payload: buf[:n],
		}

		conn.Write([]byte("thank you for your message!\n"))
	}
}
