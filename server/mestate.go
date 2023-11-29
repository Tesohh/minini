package server

import (
	"github.com/Tesohh/minini/client"
	"github.com/Tesohh/minini/message"
)

func mestate(c *client.Client, m message.Msg) error {
	json, err := message.DataToJson(c.State)
	if err != nil {
		return err
	}
	c.Send(message.Msg{Action: "me.state", Data: json})
	return nil
}
