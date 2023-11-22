package client

import (
	"encoding/json"
	"net"

	"github.com/Tesohh/minini/message"
)

type PlayerState struct {
	X, Y int
}

type Client struct {
	Conn          net.Conn
	Authenticated bool
	Msgch         chan message.Msg
	// Inputch and Outputch?
	PlayerID string // TODO: Change with mongodb ids?
	State    PlayerState
}

func (c *Client) Send(m message.Msg) error {
	marsh, err := json.Marshal(m)
	if err != nil {
		return err
	}

	_, err = c.Conn.Write(marsh)
	return err
}

func (c *Client) Error(err error) {
	c.Send(message.Msg{Data: map[string]any{"error": err.Error()}})
}
