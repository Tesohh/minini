package action

import (
	"github.com/Tesohh/minini/client"
	"github.com/Tesohh/minini/db"
	"github.com/Tesohh/minini/message"
	"github.com/Tesohh/minini/rp"
	"golang.org/x/crypto/bcrypt"
)

type loginMsg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c *client.Client, m message.Msg) error {
	d, err := message.Data[loginMsg](m)
	if err != nil {
		return err
	}

	dbuser, err := rp.Global.DB.Users.One(db.Query{"username": d.Username})
	if err != nil {
		return ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbuser.Password), []byte(d.Password))
	if err != nil {
		return ErrWrongPassword
	}

	c.Authenticated = true
	c.PlayerID = dbuser.ID.String()
	c.State = client.PlayerState{}

	return nil
}
