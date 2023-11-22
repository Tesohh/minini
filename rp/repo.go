package rp

import (
	"github.com/Tesohh/minini/data"
	"github.com/Tesohh/minini/db"
)

var Global Repo // HOLDUP: do we do this or nah?

type StoreHolder struct {
	Users db.Storer[data.User]
}

type Repo struct {
	DB StoreHolder
	// Server server.Server
}
