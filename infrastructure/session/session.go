package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql"
	"github.com/smokers10/go-diagem.git/infrastructure/config"
)

func MYSQLSessionStore(productionStore string) *session.Store {
	mysqlConfig := config.ReadConfig().Database

	mysql := mysql.New(mysql.Config{
		Host:       mysqlConfig.MYSQL_Host,
		Port:       mysqlConfig.MYSQL_Port,
		Username:   mysqlConfig.MYSQL_Username,
		Password:   mysqlConfig.MYSQL_Password,
		Database:   mysqlConfig.MYSQL_Database,
		Table:      mysqlConfig.MYSQL_SessionTable,
		Reset:      false,
		GCInterval: 10 * time.Second,
	})

	return session.New(session.Config{
		Storage:    mysql,
		Expiration: (24 * time.Hour) * 7,
	})
}
