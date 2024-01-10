package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func New() *sql.DB {
	db, err := sql.Open("sqlite3", "./mound.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
