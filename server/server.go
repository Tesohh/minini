package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/Tesohh/minini/client"
	"github.com/Tesohh/minini/message"
)

type Server struct {
	ListenAddr string
	Ln         net.Listener
	Quitch     chan struct{}
	Clients    map[net.Addr]*client.Client
	Actions    map[string]func() // TODO: Replace with Action
}

func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr: listenAddr,
		Quitch:     make(chan struct{}),
		Clients:    make(map[net.Addr]*client.Client),
		Actions:    make(map[string]func()),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}

	slog.Info("Server started on", "address", s.ListenAddr)

	defer ln.Close()
	s.Ln = ln

	go s.AcceptNewConnections()

	<-s.Quitch

	return nil
}

func (s *Server) AcceptNewConnections() {
	for {
		conn, err := s.Ln.Accept()
		if err != nil {
			slog.Warn("Server.AcceptNewConnections error: ", err)
			continue
		}

		c := &client.Client{
			Conn:          conn,
			Authenticated: false,
			Msgch:         make(chan message.Msg),
			PlayerID:      "",
			State:         client.PlayerState{},
		}

		s.Clients[conn.RemoteAddr()] = c
		slog.Info("Client connected", "address", c.Conn.RemoteAddr())

		go s.ReadFromClient(c)
	}
}

func (s *Server) ReadFromClient(c *client.Client) {
	defer func() {
		addr := c.Conn.RemoteAddr()
		c.Conn.Close()
		delete(s.Clients, addr)
	}()

	buf := make([]byte, 2048)
	for {
		length, err := c.Conn.Read(buf)

		if err != nil {
			if err.Error() == "EOF" {
				slog.Info("Client disconnected", "address", c.Conn.RemoteAddr())
				break
			} else {
				slog.Warn("Server.ReadFromClient error while reading", "error", err)
				continue
			}
		}

		fmt.Print(string(buf[:length]))
	}
}
