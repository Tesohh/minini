package message

import (
	"encoding/json"
	"net"
)

type Msg struct {
	From   net.Addr       `json:"from,omitempty"`
	To     net.Addr       `json:"to,omitempty"` // HOLDUP: net.Addr or player id?
	Action string         `json:"action,omitempty"`
	Data   map[string]any `json:"data,omitempty"`
}

func Data[T any](m Msg) (*T, error) {
	remarsh, err := json.Marshal(m.Data)
	if err != nil {
		return nil, err
	}

	var doc T
	err = json.Unmarshal(remarsh, &doc)
	return &doc, err
}
