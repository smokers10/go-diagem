package database

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	URI                           string // URI dari alamat server
	MAX_CONNECTION                int    //koneksi maksimal yang aktif
	MAX_IDLE_CONNECTION           int    //koneksi maskimal idle
	MAX_CONNECTION_LIFE_TIME      int    //rentan waktu maksimal tiap koneksi aktif (menit)
	MAX_IDLE_CONNECTION_LIFE_TIME int    //rentan waktu masimal tiap koneksi idle (menit)
}

//akan mengembalikan nilai konfigurasi sesuai dari mode produksi PRIVATE
func setConfig(mode string) *config {
	con := config{}
	if mode == "" || mode == "development" || mode == "local" {
		con.URI = "remote_admin:kingmesix@/diagem"
		con.MAX_CONNECTION = 5
		con.MAX_CONNECTION_LIFE_TIME = 1
		con.MAX_IDLE_CONNECTION = 5
		con.MAX_IDLE_CONNECTION_LIFE_TIME = 1
	} else {
		URI := os.Getenv("MYSQL_URI")

		con.URI = URI
		con.MAX_CONNECTION = 10
		con.MAX_CONNECTION_LIFE_TIME = 1
		con.MAX_IDLE_CONNECTION = 10
		con.MAX_IDLE_CONNECTION_LIFE_TIME = 1
	}
	return &con
}

/*
	MYSQLConn digunakan untuk membuat koneksi dengan server MYSQL
	parameter mode adalah penentu apakah dalam mode deploy(online) atau development(pengembangan)
*/
func MYSQLConn(mode string) (database *sql.DB, err error) {
	con := setConfig(mode)
	db, err := sql.Open("mysql", con.URI)
	if err != nil {
		return &sql.DB{}, err
	}

	db.SetMaxOpenConns(con.MAX_CONNECTION)
	db.SetMaxIdleConns(con.MAX_IDLE_CONNECTION)
	db.SetConnMaxLifetime(time.Duration(con.MAX_CONNECTION_LIFE_TIME) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(con.MAX_IDLE_CONNECTION_LIFE_TIME) * time.Minute)

	return db, nil
}
