package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBconn *sql.DB

func SetupDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseURL := os.Getenv("DATABASE_URL")
	DBconn, err = sql.Open("mysql", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
}
