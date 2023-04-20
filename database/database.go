package database

import (
	"database/sql"
	"log"
)

var DBconn *sql.DB

func SetupDB() {
	var err error
	DBconn, err = sql.Open("mysql", "${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(127.0.0.1:3306)/${MYSQL_DATABASE}")
	if err != nil {
		log.Fatal(err)
	}
}
