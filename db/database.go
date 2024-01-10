package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func New() *sql.DB {
	psqlInfo := "postgresql://Dee9975:XVN9pOyMUF3a@ep-lucky-lab-98347479.eu-central-1.aws.neon.tech/syrupstore?sslmode=require"
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
