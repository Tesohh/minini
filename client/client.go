package client

import (
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
	PlayerID      string // TODO: Change with mongodb ids?
	State         PlayerState
}
