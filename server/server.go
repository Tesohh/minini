package server

import (
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
	Actions    map[string]ActionFunc // TODO: Replace with Action
}

func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr: listenAddr,
		Quitch:     make(chan struct{}),
		Clients:    make(map[net.Addr]*client.Client),
		Actions:    make(map[string]ActionFunc),
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
		go s.HandleMessages(c)
	}
}
