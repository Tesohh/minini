package server

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/Tesohh/minini/client"
	"github.com/Tesohh/minini/message"
)

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

		var msg message.Msg
		err = json.Unmarshal(buf[:length], &msg)
		if err != nil {
			slog.Warn("Server.ReadFromClient error while unmarshaling", "error", err)
		}

		msg.From = c.Conn.RemoteAddr()

		// TODO: refuse requests from unauthenticated cleints
		c.Msgch <- msg
	}
}

func (s *Server) HandleMessages(c *client.Client) {
	for msg := range c.Msgch {
		if !c.Authenticated && msg.Action != "login" && msg.Action != "signup" {
			c.Error(fmt.Errorf("unauthenticated"))
			continue
		}

		act, ok := s.Actions[msg.Action]
		if !ok {
			c.Error(fmt.Errorf("action doesn't exist"))
			continue
		}

		err := act(c, msg)
		if err != nil {
			c.Error(err)
			continue
		}
	}
}
