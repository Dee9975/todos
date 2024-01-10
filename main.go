package main

import (
	db2 "awesomeProject/db"
	"awesomeProject/handlers"
	"awesomeProject/store"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()

	api := e.Group("/api")
	v1 := api.Group("/v1")

	db := db2.New()
	defer db.Close()

	todoStore := store.NewTodoStore(db)

	handler := handlers.NewHandler(todoStore)
	handler.Register(v1)

	log.Fatal(e.Start(":3000"))
}
