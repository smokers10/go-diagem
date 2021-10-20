package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql"
)

func MYSQLSessionStore(productionStore string) *session.Store {
	if productionStore == "" || productionStore == "local" || productionStore == "development" {
		mysql := mysql.New(mysql.Config{
			Host:       "127.0.0.1",
			Port:       3306,
			Username:   "remote_admin",
			Password:   "kingmesix",
			Database:   "diagem",
			Table:      "session",
			Reset:      false,
			GCInterval: 10 * time.Second,
		})

		return session.New(session.Config{
			Storage:    mysql,
			Expiration: (24 * time.Hour) * 7,
		})
	} else {
		mysql := mysql.New(mysql.ConfigDefault)

		return session.New(session.Config{
			Storage:    mysql,
			Expiration: (24 * time.Hour) * 7,
		})
	}
}
