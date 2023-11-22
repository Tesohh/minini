package main

import (
	"github.com/Tesohh/minini/action"
	"github.com/Tesohh/minini/server"
)

func main() {
	s := server.NewServer(":8080")
	s.Actions = action.Actions
	s.Start()
}
