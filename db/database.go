package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func New(conn string) *sql.DB {
	log.Println("Connecting to database..." + conn)
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
