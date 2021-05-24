package main

import (
	"log"
	"rest_todo"
)

func main() {
	srv := new(rest_todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
