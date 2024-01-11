package main

import (
	db2 "awesomeProject/db"
	"awesomeProject/handlers"
	"awesomeProject/store"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func main() {
	e := echo.New()

	api := e.Group("/api")
	v1 := api.Group("/v1")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env variables %s", err)
	}

	db := db2.New(os.Getenv("DB_CONNECTION_STRING"))
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	todoStore := store.NewTodoStore(db)

	handler := handlers.NewHandler(todoStore)
	handler.Register(v1)

	log.Fatal(e.Start(":8080"))
}
