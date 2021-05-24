package main

import (
	"log"
	"rest_todo"
	"rest_todo/pkg/handler"
)

func main() {
	srv := new(rest_todo.Server)
	handlers := new(handler.Handler)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
