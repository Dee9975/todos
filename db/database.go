package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func New(conn string) *sql.DB {
	psqlInfo := conn
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
