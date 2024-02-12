package database

import (
	// "gorm.io/driver/postgres"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
	// "os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// postgres://adqvftfx:Jw_jPFtv8-rSu0iJKmJSyKJLhJWaJHFg@lucky.db.elephantsql.com/adqvftfx

	// dsn := "host=lucky.db.elephantsql.com user=adqvftfx password=Jw_jPFtv8-rSu0iJKmJSyKJLhJWaJHFg dbname=adqvftfx port=5432 sslmode=disable"
	// dsn := os.Getenv("DB_URL")
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// github.com/mattn/go-sqlite3
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Fail to connect to database")
	}
}

/*

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

*/
