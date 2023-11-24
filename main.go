package main

import (
	"log"
	"os"

	"github.com/Tesohh/minini/action"
	"github.com/Tesohh/minini/background"
	"github.com/Tesohh/minini/data"
	"github.com/Tesohh/minini/db"
	"github.com/Tesohh/minini/rp"
	"github.com/Tesohh/minini/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	s := server.NewServer(":8080")
	s.Actions = action.Actions

	dbc, err := db.NewMongoClient(os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatal(err)
	}
	rp.Global.DB = rp.StoreHolder{
		Users: db.MongoStore[data.User]{Client: dbc, Coll: dbc.Database("main").Collection("users")},
	}

	go background.SaveUsers(s)
	s.Start()
}
