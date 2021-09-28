package database

import (
	"database/sql"
	"os"
	"strconv"
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

//akan mengembalikan nilai konfigurasi sesuai dari mode produksi
func setConfig(mode string) *config {
	con := config{}
	if mode == "" || mode == "development" || mode == "local" {
		con.URI = ""
		con.MAX_CONNECTION = 5
		con.MAX_CONNECTION_LIFE_TIME = 1
		con.MAX_IDLE_CONNECTION = 5
		con.MAX_IDLE_CONNECTION_LIFE_TIME = 1
	} else {
		URI := os.Getenv("MYSQL_URI")
		MAX_CONNECTION, _ := strconv.Atoi(os.Getenv("MAX_CONNECTION"))
		MAX_CONNECTION_LIFE_TIME, _ := strconv.Atoi(os.Getenv("MAX_CONNECTION_LIFE_TIME"))
		MAX_IDLE_CONNECTION, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTION"))

		con.URI = URI
		con.MAX_CONNECTION = MAX_CONNECTION
		con.MAX_CONNECTION_LIFE_TIME = MAX_CONNECTION_LIFE_TIME
		con.MAX_IDLE_CONNECTION = MAX_IDLE_CONNECTION
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
