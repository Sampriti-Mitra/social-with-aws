package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"weekend.side/SocialMedia/config"
)

var DbDriver *sql.DB

func InitialiseDb() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 config.APPLICATION_DB_USERNAME,
		Passwd:               config.APPLICATION_DB_PASSWORD,
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               config.APPLICATION_DB_NAME,
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	DbDriver, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DbDriver.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
