package server

import (
	"github.com/Tesohh/minini/client"
	"github.com/Tesohh/minini/message"
)

type ActionFunc func(c *client.Client, m message.Msg) error

var Actions = map[string]ActionFunc{
	"login":    login,
	"signup":   signup,
	"me.state": mestate,
}
