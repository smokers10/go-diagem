package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/smokers10/go-diagem.git/infrastructure/config"
)

/*
	MYSQLConn digunakan untuk membuat koneksi dengan server MYSQL
	parameter mode adalah penentu apakah dalam mode deploy(online) atau development(pengembangan)
*/
func MYSQLConn() (database *sql.DB, err error) {
	config := config.ReadConfig()
	mysql := config.Database
	return viaTCP(mysql)
}

func viaTCP(config *config.MYSQL) (database *sql.DB, err error) {
	db, err := sql.Open("mysql", config.MYSQL_URI)
	if err != nil {
		return &sql.DB{}, err
	}

	db.SetMaxOpenConns(config.MYSQL_Max_Connection)
	db.SetMaxIdleConns(config.MYSQL_Max_Idle_Connection)
	db.SetConnMaxLifetime(time.Duration(config.MYSQL_Max_Connection_Life_Time) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(config.MYSQL_Max_Idle_Connection_Life_Time) * time.Minute)

	return db, nil
}
