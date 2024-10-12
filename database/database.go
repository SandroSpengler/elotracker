package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect(connectionString string) {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
