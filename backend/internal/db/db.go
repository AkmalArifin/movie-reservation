package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	dataSrc := fmt.Sprintf("user=%s password=%s, dbname=%s host=%s sslmode=disable", user, password, dbName, host)

	var err error
	DB, err = sql.Open("postgres", dataSrc)
	if err != nil {
		panic("could not connect database")
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
