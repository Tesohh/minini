package background

import (
	"log/slog"
	"time"

	"github.com/Tesohh/minini/db"
	"github.com/Tesohh/minini/rp"
	"github.com/Tesohh/minini/server"
)

func SaveUsers(s *server.Server) {
	for {
		time.Sleep(10 * time.Second)
		for _, c := range s.Clients {
			user, err := rp.Global.DB.Users.One(db.Query{"_id": c.PlayerID})
			if err != nil {
				slog.Warn(err.Error(), "id", c.PlayerID)
				continue
			}

			// avoid updating in case no changes are done
			if user.State == c.State {
				continue
			}

			user.State = c.State

			err = rp.Global.DB.Users.Update(db.Query{"_id": c.PlayerID}, *user)
			if err != nil {
				slog.Warn(err.Error(), "id", c.PlayerID)
				continue
			}
		}

	}
}
