package rp

import "github.com/Tesohh/minini/server"

var Global Repo // HOLDUP: do we do this or nah?

type Repo struct {
	DB     struct{} // TODO: Replace with stores
	Server server.Server
}
