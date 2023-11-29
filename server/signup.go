package server

import (
	"github.com/Tesohh/minini/action"
	"github.com/Tesohh/minini/client"
	"github.com/Tesohh/minini/data"
	"github.com/Tesohh/minini/db"
	"github.com/Tesohh/minini/message"
	"github.com/Tesohh/minini/rp"
	"golang.org/x/crypto/bcrypt"
)

type signupMsg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signup(c *client.Client, m message.Msg) error {
	d, err := message.Data[signupMsg](m)
	if err != nil {
		return err
	}

	if d.Username == "" || d.Password == "" {
		return action.ErrFieldEmpty
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(d.Password), 10)
	if err != nil {
		return err
	}

	user := data.User{
		Username: d.Username,
		Password: string(hash),
	}

	err = rp.Global.DB.Users.Put(user)
	if err != nil {
		return err
	}

	dbuser, err := rp.Global.DB.Users.One(db.Query{"username": d.Username})
	if err != nil {
		return err
	}
	c.Authenticated = true
	c.PlayerID = dbuser.ID.String()
	c.State = client.PlayerState{}

	c.Send(message.Msg{Action: "login.ok", Data: map[string]any{"playerid": c.PlayerID}})

	return nil
}
