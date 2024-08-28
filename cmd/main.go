package main

import (
	"log"

	todoapp "github.com/1tututu1"
	"github.com/1tututu1/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todoapp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running")
	}
}
