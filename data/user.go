package data

import (
	"github.com/Tesohh/minini/client"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"-" bson:"password,omitempty"`
	State    client.PlayerState `json:"state" bson:"state,omitempty"`
}

func (u User) IsEmpty() bool {
	return u == User{}
}
