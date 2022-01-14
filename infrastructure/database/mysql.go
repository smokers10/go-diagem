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
	configMYSQL := config.ReadConfig()

	db, err := sql.Open("mysql", configMYSQL.MYSQL_URI)
	if err != nil {
		return &sql.DB{}, err
	}

	db.SetMaxOpenConns(configMYSQL.MYSQL_Max_Connection)
	db.SetMaxIdleConns(configMYSQL.MYSQL_Max_Idle_Connection)
	db.SetConnMaxLifetime(time.Duration(configMYSQL.MYSQL_Max_Connection_Life_Time) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(configMYSQL.MYSQL_Max_Idle_Connection_Life_Time) * time.Minute)

	return db, nil
}
