package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func New(conn string) *sql.DB {
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
